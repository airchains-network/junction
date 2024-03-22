package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/junction/types"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SubmitPod(goCtx context.Context, msg *types.MsgSubmitPod) (*types.MsgSubmitPodResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	var stationId = msg.StationId
	var creator = msg.Creator

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
		}, status.Error(codes.FailedPrecondition, "pod number is not correct")
	}

	var podNumber = msg.PodNumber
	var previousMerkleRootHash = msg.PreviousMerkleRootHash
	var publicWitness = msg.PublicWitness

	vrfStoreKey, vrfStoreKeyByte := GetVRFKeyByte(stationId, podNumber) // "vrf/{stationId}/{podNumber}
	vrfStore := prefix.NewStore(storeAdapter, types.KeyPrefix(vrfStoreKey))
	vrfDetailsByte := vrfStore.Get(vrfStoreKeyByte)

	if vrfDetailsByte == nil {
		return &types.MsgSubmitPodResponse{
			PodStatus: false,
		}, status.Error(codes.NotFound, "vrf details not found")
	}

	var vrfDetails types.VrfRecord
	k.cdc.MustUnmarshal(vrfDetailsByte, &vrfDetails)

	if vrfDetails.IsVerified == false {
		return &types.MsgSubmitPodResponse{
			PodStatus: false,
		}, status.Error(codes.FailedPrecondition, "vrf not verified")
	}

	if station.Spsp != creator {
		return &types.MsgSubmitPodResponse{
			PodStatus: false,
		}, status.Error(codes.PermissionDenied, "not the selected pod submitter")
	}

	// check if witness format is correct
	var witness fr.Vector
	witnessCheck := json.Unmarshal(publicWitness, &witness)
	if witnessCheck != nil {
		return &types.MsgSubmitPodResponse{
			PodStatus: false,
		}, sdkerrors.ErrInvalidRequest
	}

	podStoreKey, podStoreKeyByte := GetPodKeyByte(stationId, podNumber) // "pods/{stationId}/{podNumber}
	podStore := prefix.NewStore(storeAdapter, types.KeyPrefix(podStoreKey))

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
	figureDBStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FiguresDBPath))
	podSubmittedCountKey := fmt.Sprintf("pod-submitted-count__%s", stationId)
	podNumberString := strconv.FormatUint(podNumber, 10)
	figureDBStore.Set([]byte(podSubmittedCountKey), []byte(podNumberString))

	return &types.MsgSubmitPodResponse{
		PodStatus: true,
	}, nil
}
