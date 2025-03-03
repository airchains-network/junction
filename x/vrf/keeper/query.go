package keeper

import (
	"github.com/airchains-network/junction/x/vrf/types"
)

var _ types.QueryServer = Keeper{}
