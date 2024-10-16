package keeper

import (
	"context"
	"fmt"
	"strconv"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListTrackEngagementsCustomPagination(goCtx context.Context, req *types.QueryListTrackEngagementsCustomPaginationRequest) (*types.QueryListTrackEngagementsCustomPaginationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	extTrackStationId := req.ExtTrackStationId
	schemaEngagementDbPath := BuildExtTrackSchemaEngagementsStoreKey(extTrackStationId)
	schemaEngagementStore := prefix.NewStore(storeAdapter, types.KeyPrefix(schemaEngagementDbPath))

	extTrackStationIdByte := []byte(extTrackStationId)
	// Station Metrics Check
	stationMetricsDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TrackGateFigureStoreKey))
	stationMetricsDataBytes := stationMetricsDB.Get(extTrackStationIdByte)
	if stationMetricsDataBytes == nil {
		return nil, status.Errorf(codes.NotFound, "ext track station metrics %s not found", extTrackStationId)
	}

	var stationMetrics types.StationMetrics
	k.cdc.MustUnmarshal(stationMetricsDataBytes, &stationMetrics)

	totalPodDataCount := stationMetrics.TotalPodCount
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
		return &types.QueryListTrackEngagementsCustomPaginationResponse{
			Engagements: []types.ExtTrackSchemaEngagement{},
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

	var schemaEngagements []types.ExtTrackSchemaEngagement

	if order == "asc" {
		var loopError error
		for i := offset; i < offset+limit && i < totalPodDataCount; i++ {
			key := fmt.Sprintf("%d", i+1)
			keyByte := []byte(key)
			value := schemaEngagementStore.Get(keyByte)
			if value == nil {
				loopError = status.Errorf(codes.NotFound, "ext track schema engagement %s not found", key)
				break
			}

			var engagement types.ExtTrackSchemaEngagement
			if err := k.cdc.Unmarshal(value, &engagement); err != nil {
				return nil, err
			}

			schemaEngagements = append(schemaEngagements, engagement)
		}
		if loopError != nil {
			return nil, loopError
		}
	} else if order == "desc" {
		var loopError error
		startIndex := totalPodDataCount - offset
		if startIndex < 1 {
			// No data to return
			return &types.QueryListTrackEngagementsCustomPaginationResponse{
				Engagements: []types.ExtTrackSchemaEngagement{},
			}, nil
		}
		endIndex := startIndex - limit + 1
		if endIndex < 1 {
			endIndex = 1 // Assuming keys start from 1
		}

		for i := startIndex; i >= endIndex; i-- {
			key := strconv.FormatUint(i, 10)
			keyByte := []byte(key)
			value := schemaEngagementStore.Get(keyByte)
			if value == nil {
				loopError = status.Errorf(codes.NotFound, "ext track schema engagement %s not found", key)
				break
			}

			var engagement types.ExtTrackSchemaEngagement
			if err := k.cdc.Unmarshal(value, &engagement); err != nil {
				loopError = err
				break
			}

			schemaEngagements = append(schemaEngagements, engagement)
		}
		if loopError != nil {
			return nil, loopError
		}
	} else {
		return nil, status.Error(codes.InvalidArgument, "invalid order")
	}

	paginationResponse := types.TrackgatePaginationResponse{
		Offset: offset,
		Limit:  limit,
		Order:  order,
		Total:  totalPodDataCount,
	}

	return &types.QueryListTrackEngagementsCustomPaginationResponse{
		Engagements: schemaEngagements,
		Pagination:  &paginationResponse,
	}, nil
}
