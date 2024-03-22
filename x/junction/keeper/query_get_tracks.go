package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetTracks(goCtx context.Context, req *types.QueryGetTracksRequest) (*types.QueryGetTracksResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	stationData, stationErr := k.getStationById(ctx, req.StationId)
	if stationErr != nil {
		return nil, status.Error(codes.InvalidArgument, "station not found")
	}

	trackMembers := stationData.Tracks

	return &types.QueryGetTracksResponse{
		Tracks: trackMembers,
	}, nil
}
