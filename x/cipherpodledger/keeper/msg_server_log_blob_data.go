package keeper

import (
	"bytes"
	"context"
	"encoding/json"

	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/cipherpodledger/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) LogBlobData(goCtx context.Context, msg *types.MsgLogBlobData) (*types.MsgLogBlobDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	stationId := msg.StationId
	// this will contain the blob data of the pods that are being logged, the structure of this bundle will be deferent for different networks
	podBundle := msg.PodBundle
	podRange := msg.PodRange
	daProvider := msg.DaProvider

	if daProvider == "celestia" {
		decodedCelestiaPodBundle, blob, err := k.DecodeCelestiaPodBundle(podBundle)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "failed to decode Celestia pod bundle")
		}

		blobDataBytes := decodedCelestiaPodBundle.BlobData
		commitment := decodedCelestiaPodBundle.Commitment
		namespaceID := decodedCelestiaPodBundle.NamespaceID
		height := decodedCelestiaPodBundle.Height

		// check if the commitment passed is valid
		if bytes.Equal(blob.Commitment, commitment) {
			return nil, status.Error(codes.InvalidArgument, "invalid commitment passed")
		}

		// check if the namespaceID passed is valid
		if bytes.Equal(blob.NamespaceId, namespaceID) {
			return nil, status.Error(codes.InvalidArgument, "invalid namespaceID passed")
		}

		newCelestiaDABlobSpace := CelestiaDABlobSpace{
			BlobData:    blobDataBytes,
			Commitment:  commitment,
			NamespaceID: namespaceID,
			Height:      height,
			stationId:   stationId,
			podRange:    podRange,
		}

		newCelestiaDABlobSpaceBytes, err := json.Marshal(newCelestiaDABlobSpace)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "failed to marshal Celestia DABlobSpace")
		}

		stationIdBytes := []byte(stationId)
		blobStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DABlobDataStoreKey))
		blobStore.Set(stationIdBytes, newCelestiaDABlobSpaceBytes)

		return &types.MsgLogBlobDataResponse{
			Status: true,
		}, nil
	} else if daProvider == "avail" {
		// TODO: Implement the logic for the avail DA provider

		return &types.MsgLogBlobDataResponse{
			Status: false,
		}, nil
	} else if daProvider == "eigen" {
		// TODO: Implement the logic for the eigen DA provider
		return &types.MsgLogBlobDataResponse{
			Status: false,
		}, nil
	} else {
		return nil, status.Error(codes.InvalidArgument, "invalid DA provider")
	}
}
