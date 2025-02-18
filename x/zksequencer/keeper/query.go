package keeper

import (
	"github.com/airchains-network/junction/x/zksequencer/types"
)

var _ types.QueryServer = Keeper{}
