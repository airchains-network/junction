package keeper

import (
	"context"

	"github.com/ComputerKeeda/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetStation(goCtx context.Context, req *types.QueryGetStationRequest) (*types.QueryGetStationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// get station by Id
	stations, err := k.getStationById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetStationResponse{
		Stations: &stations,
	}, nil
}
