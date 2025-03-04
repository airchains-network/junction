package jip3

import (
	"context"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/airchains-network/junction/app/upgrades"
	vrftypes "github.com/airchains-network/junction/x/vrf/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

// Upgrade defines the upgrade handler for the jip-3 upgrade
const UpgradeName = "jip-3"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{
			vrftypes.ModuleName,
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
