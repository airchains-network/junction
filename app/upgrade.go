package app

import (
	"context"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

// for regular normal upgrades
func CreateDefaultUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		versionMap := module.VersionMap{
			"junction": 101, // version 1.0.1
		}

		return versionMap, nil
	}
}
