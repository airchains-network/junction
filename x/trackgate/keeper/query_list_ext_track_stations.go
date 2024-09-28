package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListExtTrackStations(goCtx context.Context, req *types.QueryListExtTrackStationsRequest) (*types.QueryListExtTrackStationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	var extTrackStation []types.ExtTrackStations
	extTrackStationsDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExtTrackStationsDataStoreKey))

	pageResponse, err := query.Paginate(extTrackStationsDataDB, req.Pagination, func(key []byte, value []byte) error {
		var singleStationData types.ExtTrackStations
		if err := k.cdc.Unmarshal(value, &singleStationData); err != nil {
			return err
		}

		extTrackStation = append(extTrackStation, singleStationData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryListExtTrackStationsResponse{
		StationsList: extTrackStation,
		Pagination:   pageResponse,
	}, nil
}
