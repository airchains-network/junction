package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ComputerKeeda/junction/x/junction/types"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
)

func (k msgServer) SubmitPod(goCtx context.Context, msg *types.MsgSubmitPod) (*types.MsgSubmitPodResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var stationId = msg.StationId

	// check if station exists
	station, err := k.getStationById(ctx, stationId)
	if err != nil {
		return &types.MsgSubmitPodResponse{
			PodStatus: false,
		}, err
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
		return &types.MsgSubmitPodResponse{
			PodStatus: false,
		}, sdkerrors.ErrInvalidAddress
	}

	// check if pod number is correct or not
	if station.LatestPod+1 != msg.PodNumber {
		return &types.MsgSubmitPodResponse{
			PodStatus: false,
		}, sdkerrors.ErrInvalidHeight
	}

	// submit pod
	//Error := k.SubmitPodHelper(
	//	ctx,
	//	msg,
	//)
	//if Error != nil {
	//	return &types.MsgSubmitPodResponse{
	//		PodStatus: false,
	//	}, Error
	//}

	var podNumber = msg.PodNumber
	var previousMerkleRootHash = msg.PreviousMerkleRootHash
	var publicWitness = msg.PublicWitness

	// check if witness format is correct
	var witness fr.Vector
	witnessCheck := json.Unmarshal(publicWitness, &witness)
	if witnessCheck != nil {
		return &types.MsgSubmitPodResponse{
			PodStatus: false,
		}, sdkerrors.ErrInvalidRequest
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
			return &types.MsgSubmitPodResponse{
				PodStatus: false,
			}, sdkerrors.ErrInvalidRequest
		}
	}

	newPod := types.Pods{
		PodNumber:              msg.PodNumber,
		MerkleRootHash:         msg.MerkleRootHash,
		PreviousMerkleRootHash: msg.PreviousMerkleRootHash,
		ZkProof:                []byte(""),
		Witness:                msg.PublicWitness,
		Timestamp:              msg.Timestamp,
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

	return &types.MsgSubmitPodResponse{
		PodStatus: true,
	}, nil
}
