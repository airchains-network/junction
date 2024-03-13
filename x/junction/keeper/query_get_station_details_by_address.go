package keeper

import (
	"context"

	"github.com/ComputerKeeda/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetStationDetailsByAddress(goCtx context.Context, req *types.QueryGetStationDetailsByAddressRequest) (*types.QueryGetStationDetailsByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	address := req.Address
	stationIds, found := k.GetStationsIdByAddressHelper(ctx, address)

	if len(stationIds) == 0 || !found {
		// No station found on this address
		return nil, status.Error(codes.NotFound, "station not found")
	}

	var stations []types.Stations

	for _, stationId := range stationIds {
		station, err := k.getStationById(ctx, stationId)
		if err != nil {
			return nil, err
		}
		stations = append(stations, station)
	}

	return &types.QueryGetStationDetailsByAddressResponse{
		Stations: stations,
	}, nil
}
