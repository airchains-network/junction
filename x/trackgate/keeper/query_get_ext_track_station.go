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

func (k Keeper) GetExtTrackStation(goCtx context.Context, req *types.QueryGetExtTrackStationRequest) (*types.QueryGetExtTrackStationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	stationIdByte := []byte(req.Id)
	extTrackStationsDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExtTrackStationsDataStoreKey))
	extTrackStationsByte := extTrackStationsDataDB.Get(stationIdByte)

	var extTrackStation types.ExtTrackStations
	if extTrackStationsByte == nil {
		return nil, status.Error(codes.NotFound, "station not found")
	}

	k.cdc.MustUnmarshal(extTrackStationsByte, &extTrackStation)

	return &types.QueryGetExtTrackStationResponse{
		Station: &extTrackStation,
	}, nil
}
