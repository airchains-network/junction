package app

import (
	"context"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	wasmtypes "github.com/airchains-network/junction/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

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
