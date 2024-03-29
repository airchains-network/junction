package keeper

import (
	"encoding/json"
	"fmt"
	"strconv"

	bls12381 "github.com/airchains-network/gnark/backend/groth16/bls12-381"
	"github.com/airchains-network/junction/x/junction/types"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetPodKeyByte(stationId string, podNumber uint64) (string, []byte) {
	podStoreKey := "pods/" + stationId
	podNumberString := strconv.FormatUint(podNumber, 10)
	podStoreKeyByte := []byte(podNumberString)
	return podStoreKey, podStoreKeyByte
}

func (k Keeper) SubmitPodHelper(ctx sdk.Context, msg *types.MsgSubmitPod) *sdkerrors.Error {

	var stationId = msg.StationId
	var podNumber = msg.PodNumber
	var merkleRootHash = msg.MerkleRootHash
	var previousMerkleRootHash = msg.PreviousMerkleRootHash
	var publicWitness = msg.PublicWitness
	var timestamp = msg.Timestamp

	// check if witness format is correct
	var witness fr.Vector
	witnessCheck := json.Unmarshal(publicWitness, &witness)
	if witnessCheck != nil {
		return sdkerrors.ErrInvalidRequest
	}

	podStoreKey, podStoreKeyByte := GetPodKeyByte(stationId, podNumber) // "pods/{stationId}/{podNumber}
	podStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(podStoreKey))

	// check whether the previous merkle root hash and the incoming hash is same or not
	if podNumber > 1 {
		_, previousPodNumber := GetPodKeyByte(stationId, podNumber-1)
		previousPodDetailsByte := podStore.Get(previousPodNumber)
		var previousPodDetails types.Pods
		k.cdc.MustUnmarshal(previousPodDetailsByte, &previousPodDetails)

		if previousPodDetails.MerkleRootHash != previousMerkleRootHash {
			return sdkerrors.ErrInvalidRequest
		}
	}

	newPod := types.Pods{
		PodNumber:              podNumber,
		MerkleRootHash:         merkleRootHash,
		PreviousMerkleRootHash: previousMerkleRootHash,
		ZkProof:                []byte(""),
		Witness:                publicWitness,
		Timestamp:              timestamp,
		IsVerified:             false,
	}

	// store new pod details
	storingData := k.cdc.MustMarshal(&newPod)
	podStore.Set(podStoreKeyByte, storingData) // "pods/{stationId}/{podNumberByte}

	// update pod-submitted-count
	figureDBStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FiguresDBPath))
	podSubmittedCountKey := fmt.Sprintf("pod-submitted-count__%s", stationId)
	podNumberString := strconv.FormatUint(podNumber, 10)
	figureDBStore.Set([]byte(podSubmittedCountKey), []byte(podNumberString))

	return nil
}

func (k Keeper) GetPodHelper(ctx sdk.Context, stationId string, podNumber uint64) (pods types.Pods, sdkErr *sdkerrors.Error) {

	podStoreKey, podStoreKeyByte := GetPodKeyByte(stationId, podNumber) // "pods/{stationId}/{podNumber}
	podStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(podStoreKey))
	podDetailsByte := podStore.Get(podStoreKeyByte)

	if podDetailsByte == nil {
		return pods, sdkerrors.ErrKeyNotFound
	}

	var podDetails types.Pods
	k.cdc.MustUnmarshal(podDetailsByte, &podDetails)

	return podDetails, nil
}

func (k Keeper) VerifyPodHelper(ctx sdk.Context, msg *types.MsgVerifyPod) error {

	stationId := msg.StationId
	podNumber := msg.PodNumber
	merkleRootHash := msg.MerkleRootHash
	previousMerkleRootHash := msg.PreviousMerkleRootHash
	zkProof := msg.ZkProof

	// get station details by id
	station, err := k.getStationById(ctx, stationId)
	if err != nil {
		return status.Error(codes.NotFound, "station not found")
	}

	// check if sender is track member
	isTrackMember := false
	for _, value := range station.Tracks {
		if value == msg.Creator {
			isTrackMember = true
		}
	}
	if !isTrackMember {
		// sender is not the Track Member
		return status.Error(codes.Unauthenticated, "sender is not the track member")
	}

	// Check if pod number is correct or not
	if station.LatestPod+1 != podNumber {
		return status.Error(codes.Unavailable, "incorrect pod number")
	}

	// get latest submitted pod == podNumber
	figureDBStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FiguresDBPath))
	podSubmittedCountKey := fmt.Sprintf("pod-submitted-count__%s", stationId)
	submittedPodNumberByte := figureDBStore.Get([]byte(podSubmittedCountKey))
	submittedPodNumString := string(submittedPodNumberByte)
	submittedPodCount, err := strconv.ParseUint(submittedPodNumString, 10, 64)
	if err != nil {
		return status.Error(codes.InvalidArgument, "error in formatting pod number")
	}
	if submittedPodCount != podNumber {
		return status.Error(codes.InvalidArgument, "incorrect pod number")
	}

	podStoreKey, podStoreKeyByte := GetPodKeyByte(stationId, podNumber) // "pods/{stationId}/{podNumber}
	podStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(podStoreKey))
	podDetailsByte := podStore.Get(podStoreKeyByte)

	if podDetailsByte == nil {
		return status.Error(codes.DataLoss, "pod detail byte conversion failed")
	}

	var currentlyStoredPod types.Pods
	k.cdc.MustUnmarshal(podDetailsByte, &currentlyStoredPod)
	if podNumber > 1 {
		if currentlyStoredPod.PreviousMerkleRootHash != previousMerkleRootHash {
			return status.Error(codes.InvalidArgument, "incorrect previous merkle root hash")
		}
		if currentlyStoredPod.MerkleRootHash != merkleRootHash {
			return status.Error(codes.InvalidArgument, "incorrect merkle root hash")
		}
	}
	// check if pod is already verified
	if currentlyStoredPod.IsVerified {
		return nil // already verified
	}

	// Verification Variables requirement and unmarshal codes below
	var proof *bls12381.Proof
	var witness fr.Vector
	var vk bls12381.VerifyingKey
	proofErr := json.Unmarshal(zkProof, &proof)
	if proofErr != nil {
		return status.Error(codes.InvalidArgument, "invalid proof provided in argument")
	}
	podWitness := currentlyStoredPod.Witness
	witnessErr := json.Unmarshal(podWitness, &witness)
	if witnessErr != nil {
		return status.Error(codes.Unavailable, "error in unmarshalling witness")
	}

	currentStationVerificationKey := station.VerificationKey
	unmarshalVkError := json.Unmarshal(currentStationVerificationKey, &vk)
	if unmarshalVkError != nil {
		return status.Error(codes.Unavailable, "error in unmarshalling verification key")
	}

	// unmarshalling of all the three variables are done now we can log this data and see the proof, witness and vk
	//if err := logData(proof, &witness, &vk); err != nil {
	//	return status.Error(codes.Internal, "unable to log files")
	//}
	verifyErr := bls12381.Verify(proof, &vk, witness)
	if verifyErr != nil {
		return status.Error(codes.Aborted, "verification failed"+verifyErr.Error())
	}

	newPod := types.Pods{
		PodNumber:              currentlyStoredPod.PodNumber,
		MerkleRootHash:         currentlyStoredPod.MerkleRootHash,
		PreviousMerkleRootHash: currentlyStoredPod.PreviousMerkleRootHash,
		ZkProof:                zkProof,
		Witness:                currentlyStoredPod.Witness,
		Timestamp:              currentlyStoredPod.Timestamp,
		IsVerified:             true,
	}
	storingData := k.cdc.MustMarshal(&newPod)
	podStoreKey, podStoreKeyByte = GetPodKeyByte(stationId, currentlyStoredPod.PodNumber) // "pods/{stationId}/{podNumber}
	podStore = prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(podStoreKey))
	podStore.Set(podStoreKeyByte, storingData)

	// update latest_verified_pod_count
	podSubmittedCountKey = fmt.Sprintf("pod-verified-count__%s", stationId)
	podNumberString := strconv.FormatUint(podNumber, 10)
	figureDBStoreKey := []byte(podSubmittedCountKey)
	figureDBStoreValue := []byte(podNumberString)
	figureDBStore.Set(figureDBStoreKey, figureDBStoreValue)

	// update station's latest merkle root hash and latest pod number
	updatedStationDetails := types.Stations{
		Tracks:               station.Tracks,
		VotingPower:          station.VotingPower,
		LatestPod:            podNumber,
		LatestMerkleRootHash: newPod.MerkleRootHash,
		VerificationKey:      station.VerificationKey,
		StationInfo:          station.StationInfo,
		Id:                   station.Id,
		Creator:              station.Creator,
	}
	stationDataDB := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StationDataKey))
	byteStation := k.cdc.MustMarshal(&updatedStationDetails)
	byteStationId := []byte(station.Id)
	stationDataDB.Set(byteStationId, byteStation)

	return nil
}

func (k Keeper) ConfirmPodVerificationHelper(ctx sdk.Context, request *types.QueryConfirmPodVerificationRequest) error {

	stationId := request.StationId
	podNumber := request.PodNumber
	merkleRootHash := request.MerkleRootHash
	previousMerkleRootHash := request.PreviousMerkleRootHash
	zkProof := request.ZkProof

	// get station details by id
	station, err := k.getStationById(ctx, stationId)
	if err != nil {
		return status.Error(codes.NotFound, "station not found")
	}

	podStoreKey, podStoreKeyByte := GetPodKeyByte(stationId, podNumber) // "pods/{stationId}/{podNumber}
	podStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(podStoreKey))
	podDetailsByte := podStore.Get(podStoreKeyByte)

	if podDetailsByte == nil {
		return status.Error(codes.DataLoss, "pod detail byte conversion failed")
	}

	var currentlyStoredPod types.Pods
	k.cdc.MustUnmarshal(podDetailsByte, &currentlyStoredPod)
	if podNumber > 1 {
		if currentlyStoredPod.PreviousMerkleRootHash != previousMerkleRootHash {
			return status.Error(codes.InvalidArgument, "incorrect previous merkle root hash")
		}
		if currentlyStoredPod.MerkleRootHash != merkleRootHash {
			return status.Error(codes.InvalidArgument, "incorrect merkle root hash")
		}
	}

	// Verification Variables requirement and unmarshal codes below
	var proof *bls12381.Proof
	var witness fr.Vector
	var vk bls12381.VerifyingKey

	proofErr := json.Unmarshal(zkProof, &proof)
	if proofErr != nil {
		return status.Error(codes.InvalidArgument, "invalid proof provided in argument")
	}

	podWitness := currentlyStoredPod.Witness
	witnessErr := json.Unmarshal(podWitness, &witness)
	if witnessErr != nil {
		return status.Error(codes.Unavailable, "error in unmarshalling witness")
	}

	currentStationVerificationKey := station.VerificationKey
	unmarshalVkError := json.Unmarshal(currentStationVerificationKey, &vk)
	if unmarshalVkError != nil {
		return status.Error(codes.Unavailable, "error in unmarshalling verification key")
	}

	verifyErr := bls12381.Verify(proof, &vk, witness)
	if verifyErr != nil {
		return status.Error(codes.Aborted, "verification failed"+verifyErr.Error())
	}

	return nil
}
