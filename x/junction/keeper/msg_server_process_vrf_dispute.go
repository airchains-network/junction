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
/* 
ProcessVrfDispute is a function used to handle disputes regarding a Verifiable Random Function (VRF).
It takes in the context and a message of type MsgProcessVrfDispute, and returns a response of type MsgProcessVrfDisputeResponse and an error.

@Problem:
When a dispute arises, all tracks of the station need to vote on the dispute. This involves creating a serialized_rc and comparing it with the 
serialized_rc of the VRF currently in use, obtained from the fetch_vrn function. If the serialized_rc matches, the track votes false, 
indicating the dispute is not valid. If not, the track votes true, indicating the dispute is valid.

@Solution:
- First, the function verifies if the station exists.
- Then it checks if the sender is a track of the station.
- Next, it verifies if the pod number matches the latest pod number of the station.
- It also checks if the dispute already exists.
- After verifying the dispute signatures, it proceeds to count the votes and stores the result.
- If the dispute result is true (indicating the VRN is invalid), the result is in favor of the verifier, and the creator is removed.
- Conversely, if the dispute result is false (indicating the VRN is valid), the result is in favor of the creator, and the verifier is removed.
- Finally, a request for a new VRN Generation is sent to the station.
*/
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
		// remove the 
		// generate a new VRN
	} else {
		// Do something when disputeResult is false
		// generate a new VRN
	}

	return &types.MsgProcessVrfDisputeResponse{}, nil
}
