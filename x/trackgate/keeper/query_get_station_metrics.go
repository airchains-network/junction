package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetStationMetrics(goCtx context.Context, req *types.QueryGetStationMetricsRequest) (*types.QueryGetStationMetricsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	extTrackStationId := req.ExtTrackStationId
	extTrackStationsDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExtTrackStationsDataStoreKey))
	extTrackStationIdByte := []byte(extTrackStationId)

	stationDataByte := extTrackStationsDataDB.Get(extTrackStationIdByte)
	if stationDataByte == nil {
		return nil, status.Errorf(codes.NotFound, "ext track station %s not found", extTrackStationId)
	}

	// Station Metrics Check
	stationMetricsDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TrackGateFigureStoreKey))
	stationMetricsDataBytes := stationMetricsDB.Get(extTrackStationIdByte)
	if stationMetricsDataBytes == nil {
		return nil, status.Errorf(codes.NotFound, "ext track station metrics %s not found", extTrackStationId)
	}

	var stationMetrics types.StationMetrics
	k.cdc.MustUnmarshal(stationMetricsDataBytes, &stationMetrics)

	return &types.QueryGetStationMetricsResponse{
		Metrics: &stationMetrics,
	}, nil
}
