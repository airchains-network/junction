package app

import (
	"context"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	trackgatemoduletypes "github.com/airchains-network/junction/x/trackgate/types"
	wasmtypes "github.com/airchains-network/junction/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

func (app *App) RegisterUpgradeHandlers() {
	// Centralized registry for upgrade handlers
	upgradeHandlers := map[string]upgradetypes.UpgradeHandler{
		"jip-2": app.handleJIP2Upgrade,
		"jip-3": app.handleJIP3Upgrade,
		// Add more handlers as needed
	}

	// Centralized registry for store upgrades
	storeUpgradeConfigs := map[string]storetypes.StoreUpgrades{
		"jip-2": {
			Added: []string{trackgatemoduletypes.StoreKey},
		},
		"jip-3": {
			Added: []string{wasmtypes.StoreKey},
		},
	}

	// Iterate over registered handlers
	for name, handler := range upgradeHandlers {
		if storeUpgrade, exists := storeUpgradeConfigs[name]; exists {
			// Read upgrade info from disk
			upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
			if err != nil {
				panic(err)
			}
			// Set store loader for upgrades with store changes
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrade))
		}
		app.UpgradeKeeper.SetUpgradeHandler(name, handler)
	}
}

func (app *App) handleJIP2Upgrade(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Check if the capability index is already set
	latestIndex := app.CapabilityKeeper.GetLatestIndex(sdkCtx)
	if latestIndex == 0 {
		// Initialize the capability index
		if err := app.CapabilityKeeper.InitializeIndex(sdkCtx, 1); err != nil {
			return nil, err
		}
	} else {
		app.Logger().Info("Capability index already initialized, skipping re-initialization")
	}

	// Define the version map for the upgrade
	versionMap := module.VersionMap{
		"trackgate": 1,
	}

	return versionMap, nil
}

func (app *App) handleJIP3Upgrade(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
	// Custom logic for JIP-3 upgrade
	//sdkCtx := sdk.UnwrapSDKContext(ctx)
	app.Logger().Info("Executing JIP-3 upgrade")

	// Example: Adding a new module or feature
	versionMap := module.VersionMap{
		"wasm": 1,
	}

	return versionMap, nil
}
