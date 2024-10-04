package app

import (
	"context"
	"strconv"

	"cosmossdk.io/log"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	trackgateKeeper "github.com/airchains-network/junction/x/trackgate/keeper"
	trackgate "github.com/airchains-network/junction/x/trackgate/module"
	trackgateTypes "github.com/airchains-network/junction/x/trackgate/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	app *App,
	logger log.Logger,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		//storeUpgrades := storetypes.StoreUpgrades{
		//	Added: []string{trackgateTypes.StoreKey},
		//}
		//
		//app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(plan.Height, &storeUpgrades))

		// Check if the capability index is already set before attempting to initialize it
		latestIndex := app.CapabilityKeeper.GetLatestIndex(sdkCtx)
		indexString := strconv.FormatUint(latestIndex, 10)
		logger.Debug(indexString)

		// Run migrations for all modules
		versionMap, err := mm.RunMigrations(sdkCtx, configurator, fromVM)
		if err != nil {
			return nil, err
		}

		if latestIndex == 0 {
			// The index is not set, so we can safely initialize it
			err := app.CapabilityKeeper.InitializeIndex(sdkCtx, 1) // Initialize with index 1 or a value > 0
			if err != nil {
				return nil, err
			}
		} else {
			logger.Info("Capability index already initialized, skipping re-initialization")
		}

		authority := authtypes.NewModuleAddress(govtypes.ModuleName)

		// Create the Trackgate Keeper
		app.TrackgateKeeper = trackgateKeeper.NewKeeper(
			app.AppCodec(),
			runtime.NewKVStoreService(app.GetKey(trackgateTypes.StoreKey)),
			logger,
			authority.String(),
			app.BankKeeper,
		)

		// Create the Trackgate AppModule
		trackgateModule := trackgate.NewAppModule(
			app.AppCodec(),
			app.TrackgateKeeper,
			app.AccountKeeper,
			app.BankKeeper,
		)

		// Register the Trackgate module using app.RegisterModules
		err = app.RegisterModules(trackgateModule)
		if err != nil {
			return nil, err
		}

		return versionMap, nil
	}
}
