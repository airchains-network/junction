package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IsTrackMember(goCtx context.Context, req *types.QueryIsTrackMemberRequest) (*types.QueryIsTrackMemberResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	trackAddress := req.TrackAddress
	stationId := req.StationId

	station, err := k.getStationById(ctx, stationId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "station not found")
	}

	tracks := station.Tracks
	var found bool
	for _, track := range tracks {
		if track == trackAddress {
			found = true
			break
		}
	}
	if !found {
		return &types.QueryIsTrackMemberResponse{
			Result: false,
		}, status.Error(codes.NotFound, "address is not a member of the valid tracks")
	} else {
		return &types.QueryIsTrackMemberResponse{
			Result: true,
		}, nil
	}
}
