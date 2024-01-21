package keeper

import (
	"github.com/airchains-network/junction/x/junction/types"
)

var _ types.QueryServer = Keeper{}
