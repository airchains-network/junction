package keeper

import (
	"github.com/airchains-network/junction/x/rollup/types"
)

var _ types.QueryServer = Keeper{}
