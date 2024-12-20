package keeper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/cipherpodledger/types"
	"github.com/celestiaorg/celestia-openrpc/types/blob"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) IntegrityCheck(goCtx context.Context, msg *types.MsgIntegrityCheck) (*types.MsgIntegrityCheckResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	stationId := msg.StationId
	blobRef := msg.BlobRef
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

	// getting station blob space data
	stationIdBytes := []byte(stationId)
	blobStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DABlobDataStoreKey))
	blobSpaceStoreDataBytes := blobStore.Get(stationIdBytes)
	if blobSpaceStoreDataBytes == nil {
		return nil, status.Error(codes.NotFound, "stationId is not registered")
	}

	if daProvider == "celestia" {
		var blobSpace CelestiaDABlobSpace
		err := json.Unmarshal(blobSpaceStoreDataBytes, &blobSpace)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "failed to unmarshal Celestia DABlobSpace")
		}

		blobDataBytes := blobSpace.BlobData
		var blobData blob.Blob
		err = json.Unmarshal(blobDataBytes, &blobData)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "failed to unmarshal Celestia DABlobSpace")
		}

		var decodedBlobRef blob.Commitment
		err = json.Unmarshal(blobRef, &decodedBlobRef)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "failed to unmarshal blobRef")
		}

		result := bytes.Equal(decodedBlobRef, blobData.Commitment)
		if result {
			upperBondPod := blobSpace.podRange[1]
			lowerBondPod := blobSpace.podRange[0]

			podStorePath := k.GetPodKeyByteForStation(stationId)
			podStore := prefix.NewStore(storeAdapter, types.KeyPrefix(podStorePath))

			for pod := lowerBondPod; pod <= upperBondPod; pod++ {
				key := []byte(fmt.Sprintf("%d", pod))
				podDetailsBytes := podStore.Get(key)
				if podDetailsBytes == nil {
					return nil, status.Error(codes.NotFound, "pod not found")
				}
				var podData types.PodData
				k.cdc.MustUnmarshal(podDetailsBytes, &podData)

				updatePodData := types.PodData{
					AscContractAddress: podData.AscContractAddress,
					PodNumber:          podData.PodNumber,
					DaBlobId:           podData.DaBlobId,
					SubmittedBy:        podData.SubmittedBy,
					Status:             "finalized",
					Timestamp:          podData.Timestamp,
					ProvingNetwork:     podData.ProvingNetwork,
					ZkFHEProof:         podData.ZkFHEProof,
					ZkFHEPublicWitness: podData.ZkFHEPublicWitness,
					IsProofVerified:    podData.IsProofVerified,
				}

				storingData := k.cdc.MustMarshal(&updatePodData)
				podStore.Set(key, storingData)
			}

			// update the fhvm metadata
			updateFhvmMetadata := types.FhvmsMeta{
				ChainId:                       fhvmMetadata.ChainId,
				ChainName:                     fhvmMetadata.ChainName,
				Status:                        fhvmMetadata.Status,
				ProofType:                     fhvmMetadata.ProofType,
				ProvingNetworkVerificationKey: fhvmMetadata.ProvingNetworkVerificationKey,
				DaProvider:                    fhvmMetadata.DaProvider,
				DaBlobId:                      fhvmMetadata.DaBlobId,
				RelayerGAddress:               fhvmMetadata.RelayerGAddress,
				RelayerAscAddress:             fhvmMetadata.RelayerAscAddress,
				PicContractAddress:            fhvmMetadata.PicContractAddress,
				AclContractAddress:            fhvmMetadata.AclContractAddress,
				TfheExecutorContractAddress:   fhvmMetadata.TfheExecutorContractAddress,
				KmsVerifierContractAddress:    fhvmMetadata.KmsVerifierContractAddress,
				GatewayContractAddress:        fhvmMetadata.GatewayContractAddress,
				AscChildContractAddress:       fhvmMetadata.AscChildContractAddress,
				LatestPodNumber:               fhvmMetadata.LatestPodNumber,
				LatestVerifiedPodNumber:       fhvmMetadata.LatestVerifiedPodNumber,
				FinalityPodNumber:             upperBondPod,
			}

			byteFhvmMetadata := k.cdc.MustMarshal(&updateFhvmMetadata)
			fhvmMetadataDB.Set(fhvmMetadataKey, byteFhvmMetadata)

			// update the pod-finalized-count
			figureDBStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FiguresDBPath))
			podFinalizedCountKey := fmt.Sprintf("zkFHE-pod-finalized-count__%s", stationId)
			podNumberString := strconv.FormatUint(upperBondPod, 10)
			figureDBStore.Set([]byte(podFinalizedCountKey), []byte(podNumberString))

			return &types.MsgIntegrityCheckResponse{
				Status: true,
			}, nil
		} else {
			return nil, status.Error(codes.InvalidArgument, "invalid blobRef passed")
		}
	} else if daProvider == "avail" {
		var blobSpace AvailDABlobSpace
		err := json.Unmarshal(blobSpaceStoreDataBytes, &blobSpace)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "failed to unmarshal Avail DABlobSpace")
		}

		result := bytes.Equal(blobRef, blobSpace.Commitment)

		if result {
			upperBondPod := blobSpace.podRange[1]
			lowerBondPod := blobSpace.podRange[0]

			podStorePath := k.GetPodKeyByteForStation(stationId)
			podStore := prefix.NewStore(storeAdapter, types.KeyPrefix(podStorePath))

			for pod := lowerBondPod; pod <= upperBondPod; pod++ {
				key := []byte(fmt.Sprintf("%d", pod))
				podDetailsBytes := podStore.Get(key)
				if podDetailsBytes == nil {
					return nil, status.Error(codes.NotFound, "pod not found")
				}
				var podData types.PodData
				k.cdc.MustUnmarshal(podDetailsBytes, &podData)

				updatePodData := types.PodData{
					AscContractAddress: podData.AscContractAddress,
					PodNumber:          podData.PodNumber,
					DaBlobId:           podData.DaBlobId,
					SubmittedBy:        podData.SubmittedBy,
					Status:             "finalized",
					Timestamp:          podData.Timestamp,
					ProvingNetwork:     podData.ProvingNetwork,
					ZkFHEProof:         podData.ZkFHEProof,
					ZkFHEPublicWitness: podData.ZkFHEPublicWitness,
					IsProofVerified:    podData.IsProofVerified,
				}

				storingData := k.cdc.MustMarshal(&updatePodData)
				podStore.Set(key, storingData)
			}

			// update the fhvm metadata
			updateFhvmMetadata := types.FhvmsMeta{
				ChainId:                       fhvmMetadata.ChainId,
				ChainName:                     fhvmMetadata.ChainName,
				Status:                        fhvmMetadata.Status,
				ProofType:                     fhvmMetadata.ProofType,
				ProvingNetworkVerificationKey: fhvmMetadata.ProvingNetworkVerificationKey,
				DaProvider:                    fhvmMetadata.DaProvider,
				DaBlobId:                      fhvmMetadata.DaBlobId,
				RelayerGAddress:               fhvmMetadata.RelayerGAddress,
				RelayerAscAddress:             fhvmMetadata.RelayerAscAddress,
				PicContractAddress:            fhvmMetadata.PicContractAddress,
				AclContractAddress:            fhvmMetadata.AclContractAddress,
				TfheExecutorContractAddress:   fhvmMetadata.TfheExecutorContractAddress,
				KmsVerifierContractAddress:    fhvmMetadata.KmsVerifierContractAddress,
				GatewayContractAddress:        fhvmMetadata.GatewayContractAddress,
				AscChildContractAddress:       fhvmMetadata.AscChildContractAddress,
				LatestPodNumber:               fhvmMetadata.LatestPodNumber,
				LatestVerifiedPodNumber:       fhvmMetadata.LatestVerifiedPodNumber,
				FinalityPodNumber:             upperBondPod,
			}

			byteFhvmMetadata := k.cdc.MustMarshal(&updateFhvmMetadata)
			fhvmMetadataDB.Set(fhvmMetadataKey, byteFhvmMetadata)
			// update the pod-finalized-count
			figureDBStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FiguresDBPath))
			podFinalizedCountKey := fmt.Sprintf("zkFHE-pod-finalized-count__%s", stationId)
			podNumberString := strconv.FormatUint(upperBondPod, 10)
			figureDBStore.Set([]byte(podFinalizedCountKey), []byte(podNumberString))

			return &types.MsgIntegrityCheckResponse{
				Status: true,
			}, nil
		} else {
			return nil, status.Error(codes.InvalidArgument, "invalid blobRef passed")
		}

	} else if daProvider == "eigen" {
		var blobSpace EigenDABlobSpace
		err := json.Unmarshal(blobSpaceStoreDataBytes, &blobSpace)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "failed to unmarshal Eigen DABlobSpace")
		}

		result := bytes.Equal(blobRef, blobSpace.Commitment)
		if result {
			upperBondPod := blobSpace.podRange[1]
			lowerBondPod := blobSpace.podRange[0]

			podStorePath := k.GetPodKeyByteForStation(stationId)
			podStore := prefix.NewStore(storeAdapter, types.KeyPrefix(podStorePath))

			for pod := lowerBondPod; pod <= upperBondPod; pod++ {
				key := []byte(fmt.Sprintf("%d", pod))
				podDetailsBytes := podStore.Get(key)
				if podDetailsBytes == nil {
					return nil, status.Error(codes.NotFound, "pod not found")
				}
				var podData types.PodData
				k.cdc.MustUnmarshal(podDetailsBytes, &podData)

				updatePodData := types.PodData{
					AscContractAddress: podData.AscContractAddress,
					PodNumber:          podData.PodNumber,
					DaBlobId:           podData.DaBlobId,
					SubmittedBy:        podData.SubmittedBy,
					Status:             "finalized",
					Timestamp:          podData.Timestamp,
					ProvingNetwork:     podData.ProvingNetwork,
					ZkFHEProof:         podData.ZkFHEProof,
					ZkFHEPublicWitness: podData.ZkFHEPublicWitness,
					IsProofVerified:    podData.IsProofVerified,
				}

				storingData := k.cdc.MustMarshal(&updatePodData)
				podStore.Set(key, storingData)
			}

			// update the fhvm metadata
			updateFhvmMetadata := types.FhvmsMeta{
				ChainId:                       fhvmMetadata.ChainId,
				ChainName:                     fhvmMetadata.ChainName,
				Status:                        fhvmMetadata.Status,
				ProofType:                     fhvmMetadata.ProofType,
				ProvingNetworkVerificationKey: fhvmMetadata.ProvingNetworkVerificationKey,
				DaProvider:                    fhvmMetadata.DaProvider,
				DaBlobId:                      fhvmMetadata.DaBlobId,
				RelayerGAddress:               fhvmMetadata.RelayerGAddress,
				RelayerAscAddress:             fhvmMetadata.RelayerAscAddress,
				PicContractAddress:            fhvmMetadata.PicContractAddress,
				AclContractAddress:            fhvmMetadata.AclContractAddress,
				TfheExecutorContractAddress:   fhvmMetadata.TfheExecutorContractAddress,
				KmsVerifierContractAddress:    fhvmMetadata.KmsVerifierContractAddress,
				GatewayContractAddress:        fhvmMetadata.GatewayContractAddress,
				AscChildContractAddress:       fhvmMetadata.AscChildContractAddress,
				LatestPodNumber:               fhvmMetadata.LatestPodNumber,
				LatestVerifiedPodNumber:       fhvmMetadata.LatestVerifiedPodNumber,
				FinalityPodNumber:             upperBondPod,
			}

			byteFhvmMetadata := k.cdc.MustMarshal(&updateFhvmMetadata)
			fhvmMetadataDB.Set(fhvmMetadataKey, byteFhvmMetadata)
			// update the pod-finalized-count
			figureDBStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FiguresDBPath))
			podFinalizedCountKey := fmt.Sprintf("zkFHE-pod-finalized-count__%s", stationId)
			podNumberString := strconv.FormatUint(upperBondPod, 10)
			figureDBStore.Set([]byte(podFinalizedCountKey), []byte(podNumberString))

			return &types.MsgIntegrityCheckResponse{
				Status: true,
			}, nil
		} else {
			return nil, status.Error(codes.InvalidArgument, "invalid blobRef passed")
		}
	} else {
		return nil, status.Error(codes.InvalidArgument, "invalid DA provider")
	}
}
