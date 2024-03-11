package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	"github.com/ComputerKeeda/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) ProcessVrfDispute(goCtx context.Context, msg *types.MsgProcessVrfDispute) (*types.MsgProcessVrfDisputeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var creator = msg.Creator
	var podNumber = msg.PodNumber
	var stationId = msg.StationId
	var signatures = msg.Signatures
	var publicKeys = msg.PublicKeys
	var votes = msg.Votes

	station, err := k.getStationById(ctx, stationId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "station not found")
	}

	tracks := station.Tracks
	var found bool
	for _, track := range tracks {
		if track == creator {
			found = true
			break
		}
	}
	if !found {
		return nil, status.Error(codes.NotFound, "sender not found in tracks")
	}

	if podNumber != station.LatestPod {
		return nil, status.Error(codes.InvalidArgument, "pod number is incorrect")
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	vrfDisputeStoreKey, vrfDisputeStoreKeyByte := GetVRFKeyByte(stationId, podNumber)
	vrfDisputeStore := prefix.NewStore(storeAdapter, types.KeyPrefix(vrfDisputeStoreKey))
	vrfDisputeDetailsByte := vrfDisputeStore.Get(vrfDisputeStoreKeyByte)
	if vrfDisputeDetailsByte != nil {
		return nil, status.Error(codes.AlreadyExists, "dispute already exists")
	}

	var voteBytes [][]byte
	for i := 0; i < len(votes); i++ {
		// voteBytes[i] = []byte{boolToByte(votes[i])}
		voteBytes = append(voteBytes, []byte{boolToByte(votes[i])})
	}

	votes_bool, signers, success, err := VerifyVrfDisputeSignatures(tracks, signatures, publicKeys, voteBytes)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if !success {
		return nil, status.Error(codes.InvalidArgument, "signature verification failed")
	}

	voteResult := countDisputeVotes(votes_bool, signers)
	disputeResult := voteResult.Result
	byteResult := k.cdc.MustMarshal(&voteResult)
	vrfDisputeStore.Set(vrfDisputeStoreKeyByte, byteResult)

	/* 
		voteResult.result == true then VRN is invalid and the dispute result is inclined in favor of the verifier
		else VRN is valid and the dispute result is inclined in favor of the creator
	*/

	if disputeResult {
		// Do something when disputeResult is true
		// remove the creator since the dispute is inclined in favor of the verifier
		// generate a new VRN
	} else {
		// Do something when disputeResult is false
		// remove the verifier since the dispute is inclined in favor of the creator
		// generate a new VRN
	}

	return &types.MsgProcessVrfDisputeResponse{}, nil
}
