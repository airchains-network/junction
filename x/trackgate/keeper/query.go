package keeper

import (
	"github.com/airchains-network/junction/x/trackgate/types"
)

var _ types.QueryServer = Keeper{}
