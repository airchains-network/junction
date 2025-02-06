package app

import (
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	wasmkeeper "github.com/airchains-network/junction/x/wasm/keeper"
)

func (app *JunctionApp) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

func (app *JunctionApp) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

func (app *JunctionApp) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

func (app *JunctionApp) GetBankKeeper() bankkeeper.Keeper {
	return app.BankKeeper
}

func (app *JunctionApp) GetStakingKeeper() *stakingkeeper.Keeper {
	return app.StakingKeeper
}

func (app *JunctionApp) GetAccountKeeper() authkeeper.AccountKeeper {
	return app.AccountKeeper
}

func (app *JunctionApp) GetWasmKeeper() wasmkeeper.Keeper {
	return app.WasmKeeper
}
