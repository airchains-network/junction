package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/ComputerKeeda/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListStations(goCtx context.Context, req *types.QueryListStationsRequest) (*types.QueryListStationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var stations []types.Stations

	stationDataDB := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StationDataKey))

	pageRes, err := query.Paginate(stationDataDB, req.Pagination, func(key []byte, value []byte) error {
		var singleStationData types.Stations
		if err := k.cdc.Unmarshal(value, &singleStationData); err != nil {
			return err
		}

		stations = append(stations, singleStationData)
		return nil
	})

	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListStationsResponse{StationsList: stations, Pagination: pageRes}, nil
}
