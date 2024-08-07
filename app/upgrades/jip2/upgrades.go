package jip2

import (
	"context"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/airchains-network/junction/app"
	wasmtypes "github.com/airchains-network/junction/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/airchains-network/junction/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name
const UpgradeName = "jip2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{
			wasmtypes.ModuleName,
		},
		Deleted: []string{},
	},
}

func CreateUpgradeHandler(
	mm upgrades.ModuleManager,
	configurator module.Configurator,
	ak *upgrades.AppKeepers,
	app *app.App,
) upgradetypes.UpgradeHandler {
	// sdk 47 to sdk 50
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{wasmtypes.StoreKey},
		}
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(plan.Height, &storeUpgrades))

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
