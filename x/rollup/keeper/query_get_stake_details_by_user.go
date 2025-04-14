package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/rollup/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetStakeDetailsByUser(goCtx context.Context, req *types.QueryGetStakeDetailsByUserRequest) (*types.QueryGetStakeDetailsByUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	innerResponse, err := k.Inner_GetStakeDetailsByUser(ctx, req.Address)
	if err != nil {
		return nil, err
	}

	return &innerResponse, nil
}
