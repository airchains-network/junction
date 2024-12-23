package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/cipherpodledger/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetPod(goCtx context.Context, req *types.QueryGetPodRequest) (*types.QueryGetPodResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	var podNumber = req.PodNumber
	var stationID = req.StationId

	// check if station id exists
	fhvmMetadataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FhvmKeyStoreKey))
	fhvmMetadataKey := []byte(stationID)
	fhvmMetadataDataBytes := fhvmMetadataDB.Get(fhvmMetadataKey)
	if fhvmMetadataDataBytes == nil {
		return nil, status.Error(codes.NotFound, "stationId is not registered")
	}

	var fhvmMetadata types.FhvmsMeta
	k.cdc.MustUnmarshal(fhvmMetadataDataBytes, &fhvmMetadata)

	// get the pod details
	podStorePath, podStoreKeyByte := k.GetPodKeyByte(stationID, podNumber)
	podStore := prefix.NewStore(storeAdapter, types.KeyPrefix(podStorePath))
	podDataBytes := podStore.Get(podStoreKeyByte)
	if podDataBytes == nil {
		return nil, status.Error(codes.NotFound, "pod not found")
	}

	var podData types.PodData
	k.cdc.MustUnmarshal(podDataBytes, &podData)

	return &types.QueryGetPodResponse{
		Pod: &podData,
	}, nil
}
