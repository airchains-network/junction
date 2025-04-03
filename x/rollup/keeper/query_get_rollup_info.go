package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/rollup/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRollupInfo(goCtx context.Context, req *types.QueryGetRollupInfoRequest) (*types.QueryGetRollupInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	rollupInfo, err := k.getRollupById(ctx, req.RollupId)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetRollupInfoResponse{
		RollupInfo: &rollupInfo,
	}, nil
}
