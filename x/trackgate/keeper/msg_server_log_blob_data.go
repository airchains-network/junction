package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) LogBlobData(goCtx context.Context, msg *types.MsgLogBlobData) (*types.MsgLogBlobDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgLogBlobDataResponse{}, nil
}
