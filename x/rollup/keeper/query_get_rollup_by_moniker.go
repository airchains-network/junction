package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/rollup/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRollupByMoniker(goCtx context.Context, req *types.QueryGetRollupByMonikerRequest) (*types.QueryGetRollupByMonikerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	rollupInfo, err := k.getRollupByMoniker(ctx, req.Moniker)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetRollupByMonikerResponse{
		RollupInfo: &rollupInfo,
	}, nil
}
