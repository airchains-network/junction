package app

import (
	"context"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"fmt"
	"github.com/airchains-network/junction/app/upgrades"
	"github.com/airchains-network/junction/app/upgrades/jip2"
	"github.com/airchains-network/junction/app/upgrades/noop"
	v2 "github.com/airchains-network/junction/x/wasm/migrations/v2"
	wasmtypes "github.com/airchains-network/junction/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var Upgrades = []upgrades.Upgrade{jip2.Upgrade}

// for regular normal upgrades
func CreateDefaultUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	app *App,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		//versionMap := module.VersionMap{
		//	"junction": 101, // version 1.0.1
		//}

		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{wasmtypes.StoreKey},
		}

		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(plan.Height, &storeUpgrades))

		//return versionMap, nil
		return app.ModuleManager.RunMigrations(ctx, configurator, fromVM)
	}
}

func (app *App) RegisterUpgradeHandlers() {
	setupLegacyKeyTables(&app.ParamsKeeper)
	if len(Upgrades) == 0 {
		// always have a unique upgrade registered for the current version to test in system tests
		Upgrades = append(Upgrades, noop.NewUpgrade(app.Version()))
	}

	keepers := upgrades.AppKeepers{
		AccountKeeper:         &app.AccountKeeper,
		ParamsKeeper:          &app.ParamsKeeper,
		ConsensusParamsKeeper: &app.ConsensusParamsKeeper,
		CapabilityKeeper:      app.CapabilityKeeper,
		IBCKeeper:             app.IBCKeeper,
		Codec:                 app.appCodec,
		GetStoreKey:           app.GetKey,
	}
	app.GetStoreKeys()
	// register all upgrade handlers
	for _, upgrade := range Upgrades {
		app.UpgradeKeeper.SetUpgradeHandler(
			upgrade.UpgradeName,
			upgrade.CreateUpgradeHandler(
				app.ModuleManager,
				app.SSConfigurator,
				&keepers,
			),
		)
	}

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	// register store loader for current upgrade
	for _, upgrade := range Upgrades {
		if upgradeInfo.Name == upgrade.UpgradeName {
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &upgrade.StoreUpgrades)) // nolint:gosec
			break
		}
	}
}

func setupLegacyKeyTables(k *paramskeeper.Keeper) {
	for _, subspace := range k.GetSubspaces() {
		subspace := subspace

		var keyTable paramstypes.KeyTable
		switch subspace.Name() {
		case authtypes.ModuleName:
			keyTable = authtypes.ParamKeyTable() //nolint:staticcheck
		case banktypes.ModuleName:
			keyTable = banktypes.ParamKeyTable() //nolint:staticcheck
		case stakingtypes.ModuleName:
			keyTable = stakingtypes.ParamKeyTable() //nolint:staticcheck
		case minttypes.ModuleName:
			keyTable = minttypes.ParamKeyTable() //nolint:staticcheck
		case distrtypes.ModuleName:
			keyTable = distrtypes.ParamKeyTable() //nolint:staticcheck
		case slashingtypes.ModuleName:
			keyTable = slashingtypes.ParamKeyTable() //nolint:staticcheck
		case govtypes.ModuleName:
			keyTable = govv1.ParamKeyTable() //nolint:staticcheck
		case crisistypes.ModuleName:
			keyTable = crisistypes.ParamKeyTable() //nolint:staticcheck
			// wasm
		case wasmtypes.ModuleName:
			keyTable = v2.ParamKeyTable() //nolint:staticcheck
		default:
			continue
		}

		if !subspace.HasKeyTable() {
			subspace.WithKeyTable(keyTable)
		}
	}
	// sdk 47
	k.Subspace(baseapp.Paramspace).
		WithKeyTable(paramstypes.ConsensusParamsKeyTable())
}
