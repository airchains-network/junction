package app

import (
	"context"
	storetypes "cosmossdk.io/store/types"
	trackgatemoduletypes "github.com/airchains-network/junction/x/trackgate/types"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

// for regular normal upgrades
func CreateDefaultUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	app *App,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{trackgatemoduletypes.StoreKey},
		}

		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(plan.Height, &storeUpgrades))

		//return versionMap, nil
		return app.ModuleManager.RunMigrations(ctx, configurator, fromVM)
	}
}
