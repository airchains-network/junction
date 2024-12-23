package keeper

import (
	"github.com/airchains-network/junction/x/cipherpodledger/types"
)

var _ types.QueryServer = Keeper{}
