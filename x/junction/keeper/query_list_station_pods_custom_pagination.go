package keeper

import (
	"context"
	"fmt"
	"strconv"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	"github.com/airchains-network/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListStationPodsCustomPagination(goCtx context.Context, req *types.QueryListStationPodsCustomPaginationRequest) (*types.QueryListStationPodsCustomPaginationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	stationId := req.StationId

	podStoreKey := "pods/" + stationId
	podStore := prefix.NewStore(storeAdapter, types.KeyPrefix(podStoreKey))

	// find station
	// check if station exists
	station, err := k.getStationById(ctx, stationId)
	if err != nil {
		return &types.QueryListStationPodsCustomPaginationResponse{
			Pods:       nil,
			Pagination: nil,
		}, err
	}
	totalPodDataCount := station.LatestPod

	offset := req.Offset
	limit := req.Limit
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 100
	}
	if offset >= totalPodDataCount {
		// No data to return
		return &types.QueryListStationPodsCustomPaginationResponse{
			Pods:       nil,
			Pagination: nil,
		}, nil
	}

	// Adjust limit if it exceeds the total count
	if offset+limit > totalPodDataCount {
		limit = totalPodDataCount - offset
	}

	order := req.Order
	if order == "" {
		order = "desc"
	}
	var pods []types.Pods

	if order == "asc" {
		var loopError error
		for i := offset; i < offset+limit && i < totalPodDataCount; i++ {
			key := fmt.Sprintf("%d", i+1)
			keyByte := []byte(key)
			value := podStore.Get(keyByte)
			if value == nil {
				loopError = status.Errorf(codes.NotFound, "ext track schema engagement %s not found", key)
				break
			}

			var pod types.Pods
			if err := k.cdc.Unmarshal(value, &pod); err != nil {
				return nil, err
			}

			pods = append(pods, pod)
		}
		if loopError != nil {
			return nil, loopError
		}
	} else if order == "desc" {
		var loopError error
		startIndex := totalPodDataCount - offset
		if startIndex < 1 {
			// No data to return
			return &types.QueryListStationPodsCustomPaginationResponse{
				Pods:       nil,
				Pagination: nil,
			}, nil
		}
		endIndex := startIndex - limit + 1
		if endIndex < 1 {
			endIndex = 1 // Assuming keys start from 1
		}

		for i := startIndex; i >= endIndex; i-- {
			key := strconv.FormatUint(i, 10)
			keyByte := []byte(key)
			value := podStore.Get(keyByte)
			if value == nil {
				loopError = status.Errorf(codes.NotFound, "ext track schema engagement %s not found", key)
				break
			}

			var pod types.Pods
			if err := k.cdc.Unmarshal(value, &pod); err != nil {
				loopError = err
				break
			}

			pods = append(pods, pod)
		}
		if loopError != nil {
			return nil, loopError
		}
	} else {
		return nil, status.Error(codes.InvalidArgument, "invalid order")
	}

	paginationResponse := types.JunctionPaginationResponse{
		Offset: offset,
		Limit:  limit,
		Order:  order,
		Total:  totalPodDataCount,
	}
	return &types.QueryListStationPodsCustomPaginationResponse{
		Pods:       pods,
		Pagination: &paginationResponse,
	}, nil
}
