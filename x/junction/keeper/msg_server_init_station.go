package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) InitStation(goCtx context.Context, msg *types.MsgInitStation) (*types.MsgInitStationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgInitStationResponse{}, nil
}
