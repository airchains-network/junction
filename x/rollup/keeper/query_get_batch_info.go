package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/rollup/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetBatchInfo(goCtx context.Context, req *types.QueryGetBatchInfoRequest) (*types.QueryGetBatchInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	batchInfo, err := k.getRollupBatchInfoByBatchNo(ctx, req.RollupId, req.BatchNo)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetBatchInfoResponse{BatchInfo: &batchInfo}, nil
}
