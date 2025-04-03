package jip2

import (
	"context"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	rolluptypes "github.com/airchains-network/junction/x/rollup/types"

	"github.com/airchains-network/junction/app/upgrades"
	"github.com/cosmos/cosmos-sdk/types/module"
)

const UpgradeName = "jip-2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{
			rolluptypes.ModuleName,
		},
		Deleted: []string{},
	},
}

// CreateUpgradeHandler returns an upgrade handler that migrates the chain from
func CreateUpgradeHandler(
	mm upgrades.ModuleManager,
	configurator module.Configurator,
	ak *upgrades.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
