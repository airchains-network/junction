package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/rollup/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRollupStakedAmount(goCtx context.Context, req *types.QueryGetRollupStakedAmountRequest) (*types.QueryGetRollupStakedAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	innerResponse, err := k.Inner_GetRollupStakedAmount(ctx, req.RollupId)
	if err != nil {
		return nil, err
	}

	return &innerResponse, nil
}
