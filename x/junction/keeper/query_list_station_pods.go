package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/airchains-network/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListStationPods(goCtx context.Context, req *types.QueryListStationPodsRequest) (*types.QueryListStationPodsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	stationId := req.StationId
	var pods []types.Pods

	podStoreKey := "pods/" + stationId
	podStore := prefix.NewStore(storeAdapter, types.KeyPrefix(podStoreKey))

	pageRes, err := query.Paginate(podStore, req.Pagination, func(key []byte, value []byte) error {
		var singlePodData types.Pods
		if err := k.cdc.Unmarshal(value, &singlePodData); err != nil {
			return err
		}

		pods = append(pods, singlePodData)
		return nil
	})
	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListStationPodsResponse{
		Pods:       pods,
		Pagination: pageRes,
	}, nil
}
