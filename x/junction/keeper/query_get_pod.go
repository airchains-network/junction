package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetPod(goCtx context.Context, req *types.QueryGetPodRequest) (*types.QueryGetPodResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var stationID = req.StationId
	var podNumber = req.PodNumber

	// check if station id exists
	_, err := k.getStationById(ctx, stationID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Station ID do not exists")
	}

	podDetails, sdkError := k.GetPodHelper(ctx, stationID, podNumber)
	if sdkError != nil {
		return nil, status.Error(codes.InvalidArgument, "Pod Not found")
	}

	return &types.QueryGetPodResponse{
		Pod: &podDetails,
	}, nil
}
