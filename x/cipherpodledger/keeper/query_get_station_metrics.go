package keeper

import (
	"context"
	"fmt"
	"strconv"

	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/cipherpodledger/types"
	"github.com/cosmos/cosmos-sdk/runtime"
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

	stationID := req.StationId

	fhvmMetadataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FhvmKeyStoreKey))
	fhvmMetadataKey := []byte(stationID)
	fhvmMetadataDataBytes := fhvmMetadataDB.Get(fhvmMetadataKey)
	if fhvmMetadataDataBytes == nil {
		return nil, status.Error(codes.NotFound, "stationId is not registered")
	}

	var fhvmMetadata types.FhvmsMeta
	k.cdc.MustUnmarshal(fhvmMetadataDataBytes, &fhvmMetadata)

	figureDBStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FiguresDBPath))
	podSubmittedCountKey := fmt.Sprintf("zkFHE-pod-submitted-count__%s", stationID)
	podFinalizedCountKey := fmt.Sprintf("zkFHE-pod-finalized-count__%s", stationID)
	podVerifiedCountKey := fmt.Sprintf("zkFHE-pod-verified-count__%s", stationID)

	podSubmittedCountKeyBytes := []byte(podSubmittedCountKey)
	podFinalizedCountKeyBytes := []byte(podFinalizedCountKey)
	podVerifiedCountKeyBytes := []byte(podVerifiedCountKey)

	podSubmittedCount := figureDBStore.Get(podSubmittedCountKeyBytes)
	podFinalizedCount := figureDBStore.Get(podFinalizedCountKeyBytes)
	podVerifiedCount := figureDBStore.Get(podVerifiedCountKeyBytes)

	totalFinalizedPodCount, _ := strconv.ParseUint(string(podFinalizedCount), 10, 64)
	totalSubmittedPodCount, _ := strconv.ParseUint(string(podSubmittedCount), 10, 64)
	totalVerifiedPodCount, _ := strconv.ParseUint(string(podVerifiedCount), 10, 64)

	return &types.QueryGetStationMetricsResponse{
		TotalFinalizedPodCount: totalFinalizedPodCount,
		TotalSubmittedPodCount: totalSubmittedPodCount,
		TotalVerifiedPodCount:  totalVerifiedPodCount,
		StationDetails:         &fhvmMetadata,
	}, nil
}
