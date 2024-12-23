package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type JunctionKeeper interface {
	// TODO Add methods imported from junction should be defined here
}

type WasmKeeper interface {
	// TODO Add methods imported from wasm should be defined here
}

type TrackgateKeeper interface {
	// TODO Add methods imported from trackgate should be defined here
}

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}
