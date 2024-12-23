package keeper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

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

	// getting station metadata
	fhvmMetadataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FhvmKeyStoreKey))
	fhvmMetadataKey := []byte(stationId)
	fhvmMetadataDataBytes := fhvmMetadataDB.Get(fhvmMetadataKey)
	if fhvmMetadataDataBytes == nil {
		return nil, status.Error(codes.NotFound, "stationId is not registered")
	}

	var fhvmMetadata types.FhvmsMeta
	k.cdc.MustUnmarshal(fhvmMetadataDataBytes, &fhvmMetadata)

	finalityPodNumber := fhvmMetadata.FinalityPodNumber
	lowerBoundPodNumber := podRange[0]
	upperBoundPodNumber := podRange[1]

	if lowerBoundPodNumber == 0 {
		return nil, status.Error(codes.InvalidArgument, "lower bound pod number cannot be 0")
	}

	if upperBoundPodNumber < lowerBoundPodNumber {
		return nil, status.Error(codes.InvalidArgument, "pod range is not valid")
	}

	// checking if the pod range is valid
	if lowerBoundPodNumber < finalityPodNumber {
		return nil, status.Error(codes.InvalidArgument, "pod range is not valid")
	}

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

		podStorePath := k.GetPodKeyByteForStation(stationId)
		podStore := prefix.NewStore(storeAdapter, types.KeyPrefix(podStorePath))
		for pod := lowerBoundPodNumber; pod <= upperBoundPodNumber; pod++ {
			key := []byte(fmt.Sprintf("%d", pod))
			podDetailsBytes := podStore.Get(key)
			if podDetailsBytes == nil {
				return nil, status.Error(codes.NotFound, "pod not found")
			}
			var podData types.PodData
			k.cdc.MustUnmarshal(podDetailsBytes, &podData)

			updatePodData := types.PodData{
				AscContractAddress: podData.AscContractAddress,
				PodNumber:          pod,
				DaBlobId:           podData.DaBlobId,
				SubmittedBy:        podData.SubmittedBy,
				Status:             "da blob logged",
				Timestamp:          podData.Timestamp,
				ProvingNetwork:     podData.ProvingNetwork,
				ZkFHEProof:         podData.ZkFHEProof,
				ZkFHEPublicWitness: podData.ZkFHEPublicWitness,
				IsProofVerified:    podData.IsProofVerified,
			}

			storingData := k.cdc.MustMarshal(&updatePodData)
			podStore.Set(key, storingData)
		}

		return &types.MsgLogBlobDataResponse{
			Status: true,
		}, nil
	} else if daProvider == "avail" {
		decodedAvailPodBundle, err := k.DecodeAvailPodBundle(podBundle)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "failed to decode Avail pod bundle")
		}

		blobDataBytes := decodedAvailPodBundle.BlobData
		commitment := decodedAvailPodBundle.Commitment

		newAvailDABlobSpace := AvailDABlobSpace{
			BlobData:   blobDataBytes,
			Commitment: commitment,
			stationId:  stationId,
			podRange:   podRange,
		}

		newAvailDABlobSpaceBytes, err := json.Marshal(newAvailDABlobSpace)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "failed to marshal Avail DABlobSpace")
		}

		podStorePath := k.GetPodKeyByteForStation(stationId)
		podStore := prefix.NewStore(storeAdapter, types.KeyPrefix(podStorePath))
		for pod := lowerBoundPodNumber; pod <= upperBoundPodNumber; pod++ {
			key := []byte(fmt.Sprintf("%d", pod))
			podDetailsBytes := podStore.Get(key)
			if podDetailsBytes == nil {
				return nil, status.Error(codes.NotFound, "pod not found")
			}
			var podData types.PodData
			k.cdc.MustUnmarshal(podDetailsBytes, &podData)

			updatePodData := types.PodData{
				AscContractAddress: podData.AscContractAddress,
				PodNumber:          pod,
				DaBlobId:           podData.DaBlobId,
				SubmittedBy:        podData.SubmittedBy,
				Status:             "da blob logged",
				Timestamp:          podData.Timestamp,
				ProvingNetwork:     podData.ProvingNetwork,
				ZkFHEProof:         podData.ZkFHEProof,
				ZkFHEPublicWitness: podData.ZkFHEPublicWitness,
				IsProofVerified:    podData.IsProofVerified,
			}

			storingData := k.cdc.MustMarshal(&updatePodData)
			podStore.Set(key, storingData)
		}

		blobStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DABlobDataStoreKey))
		stationIdBytes := []byte(stationId)
		blobStore.Set(stationIdBytes, newAvailDABlobSpaceBytes)

		return &types.MsgLogBlobDataResponse{
			Status: true,
		}, nil
	} else if daProvider == "eigen" {
		decodedEigenPodBundle, err := k.DecodeEigenPodBundle(podBundle)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "failed to decode Eigen pod bundle")
		}

		blobDataBytes := decodedEigenPodBundle.BlobData
		commitment := decodedEigenPodBundle.Commitment

		newEigenDABlobSpace := EigenDABlobSpace{
			BlobData:   blobDataBytes,
			Commitment: commitment,
			stationId:  stationId,
			podRange:   podRange,
		}

		newEigenDABlobSpaceBytes, err := json.Marshal(newEigenDABlobSpace)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "failed to marshal Eigen DABlobSpace")
		}

		podStorePath := k.GetPodKeyByteForStation(stationId)
		podStore := prefix.NewStore(storeAdapter, types.KeyPrefix(podStorePath))
		for pod := lowerBoundPodNumber; pod <= upperBoundPodNumber; pod++ {
			key := []byte(fmt.Sprintf("%d", pod))
			podDetailsBytes := podStore.Get(key)
			if podDetailsBytes == nil {
				return nil, status.Error(codes.NotFound, "pod not found")
			}
			var podData types.PodData
			k.cdc.MustUnmarshal(podDetailsBytes, &podData)

			updatePodData := types.PodData{
				AscContractAddress: podData.AscContractAddress,
				PodNumber:          pod,
				DaBlobId:           podData.DaBlobId,
				SubmittedBy:        podData.SubmittedBy,
				Status:             "da blob logged",
				Timestamp:          podData.Timestamp,
				ProvingNetwork:     podData.ProvingNetwork,
				ZkFHEProof:         podData.ZkFHEProof,
				ZkFHEPublicWitness: podData.ZkFHEPublicWitness,
				IsProofVerified:    podData.IsProofVerified,
			}

			storingData := k.cdc.MustMarshal(&updatePodData)
			podStore.Set(key, storingData)
		}

		blobStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DABlobDataStoreKey))
		stationIdBytes := []byte(stationId)
		blobStore.Set(stationIdBytes, newEigenDABlobSpaceBytes)

		return &types.MsgLogBlobDataResponse{
			Status: true,
		}, nil
	} else {
		return nil, status.Error(codes.InvalidArgument, "invalid DA provider")
	}
}
