package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/ComputerKeeda/junction/x/junction/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListStations(goCtx context.Context, req *types.QueryListStationsRequest) (*types.QueryListStationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var stations []types.Stations

	stationDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StationDataKey))

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
