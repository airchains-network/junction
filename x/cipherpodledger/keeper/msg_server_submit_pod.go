package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"

	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/cipherpodledger/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) SubmitPod(goCtx context.Context, msg *types.MsgSubmitPod) (*types.MsgSubmitPodResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	submittedBy := msg.SubmittedBy
	ascChildContractAddress := msg.AscChildContractAddress
	podNumber := msg.PodNumber
	stationId := msg.StationId
	daBlobId := msg.DaBlobId
	timestamp := msg.Timestamp
	provingNetwork := msg.ProvingNetwork
	decodedZkFhepublicWitness := msg.ZkFhepublicWitness

	fhvmMetadataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FhvmKeyStoreKey))
	fhvmMetadataKey := []byte(stationId)
	fhvmMetadataDataBytes := fhvmMetadataDB.Get(fhvmMetadataKey)
	if fhvmMetadataDataBytes == nil {
		return nil, status.Error(codes.NotFound, "stationId is not registered")
	}

	var fhvmMetadata types.FhvmsMeta
	k.cdc.MustUnmarshal(fhvmMetadataDataBytes, &fhvmMetadata)

	// check if the station id is correct or not
	chainIdHash := sha256.Sum256([]byte(ascChildContractAddress))
	if stationId != hex.EncodeToString(chainIdHash[:]) {
		return nil, status.Error(codes.InvalidArgument, "chainId is not valid")
	}

	// check if the pod number is correct or not
	if fhvmMetadata.LatestPodNumber+1 != podNumber {
		return nil, status.Error(codes.FailedPrecondition, "pod number is not correct")
	}

	newPod := types.PodData{
		AscContractAddress: ascChildContractAddress,
		PodNumber:          podNumber,
		DaBlobId:           daBlobId,
		SubmittedBy:        submittedBy,
		Status:             "pending",
		Timestamp:          timestamp,
		ProvingNetwork:     provingNetwork,
		ZkFHEProof:         nil,
		ZkFHEPublicWitness: decodedZkFhepublicWitness,
		IsProofVerified:    false,
	}

	// store new pod details
	storingData := k.cdc.MustMarshal(&newPod)
	podStorePath, podStoreKeyByte := k.GetPodKeyByte(stationId, podNumber)
	podStore := prefix.NewStore(storeAdapter, types.KeyPrefix(podStorePath))
	podStore.Set(podStoreKeyByte, storingData)

	// update pod-submitted-count
	figureDBStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FiguresDBPath))
	podSubmittedCountKey := fmt.Sprintf("zkFHE-pod-submitted-count__%s", stationId)
	podNumberString := strconv.FormatUint(podNumber, 10)
	figureDBStore.Set([]byte(podSubmittedCountKey), []byte(podNumberString))

	return &types.MsgSubmitPodResponse{
		Status: true,
	}, nil
}
