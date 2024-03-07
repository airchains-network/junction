package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/ComputerKeeda/junction/x/junction/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) ValidateVrf(goCtx context.Context, msg *types.MsgValidateVrf) (*types.MsgValidateVrfResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	var creator = msg.Creator
	var stationId = msg.StationId
	var podNumber = msg.PodNumber
	var serializedRc = msg.SerializedRc

	// check if station exists
	station, err := k.getStationById(ctx, stationId)
	if err != nil {
		return &types.MsgValidateVrfResponse{
			Success: false,
		}, status.Error(codes.NotFound, "station not found")
	}

	// check if sender is track member
	isTrackMember := false
	for _, value := range station.Tracks {
		if value == creator {
			isTrackMember = true
		}
	}
	if !isTrackMember {
		// sender is not the Track Member
		return &types.MsgValidateVrfResponse{
			Success: false,
		}, sdkerrors.ErrInvalidAddress
	}

	// check if pod number is correct or not
	if station.LatestPod != podNumber {
		return &types.MsgValidateVrfResponse{
			Success: false,
		}, sdkerrors.ErrInvalidHeight
	}

	// !come here after completing the vrf validation logic
	/*
		here we have to first of all change the stations related messages and queries to the new structure
		after that we have to add a new key in station type
		which will tell us about the latest pod validated or not
		after that we have to check here if the pod is already validated or not
		if not then we have to validate the pod and then update the station
		else we have to return the error
	*/
	// check if pod is already validated
	// if station.LatestPodValidated {
	// 	return &types.MsgValidateVrfResponse{
	// 		Success: false,
	// 	}, sdkerrors.ErrUnauthorized
	// }

	// get the vrf details from the store
	vrfStoreKey, vrfStoreKeyByte := GetVRFKeyByte(stationId, podNumber) // "vrf/{stationId}/{podNumber}
	vrfStore := prefix.NewStore(storeAdapter, types.KeyPrefix(vrfStoreKey))

	vrfDetailsByte := vrfStore.Get(vrfStoreKeyByte)

	// if the vrf details are not present then we have to return the error
	if vrfDetailsByte == nil {
		return &types.MsgValidateVrfResponse{
			Success: false,
		}, status.Error(codes.NotFound, "vrf details not found")
	}

	var vrfDetails types.VrfRecord
	k.cdc.MustUnmarshal(vrfDetailsByte, &vrfDetails)
	
	if vrfDetails.IsVerified == true {
		// already verified 
		return &types.MsgValidateVrfResponse{
			Success: false,
		}, status.Error(codes.AlreadyExists, "vrf details already verified")
	}

	// if the vrf details are present then we have to validate the vrf

	hexPublicKey := vrfDetails.CreatorsVrfKey
	proof := vrfDetails.Proof
	vrfOutput := vrfDetails.VrfOutput

	valid, err := VerifyVRFProof(hexPublicKey, serializedRc, proof, int64(podNumber), vrfOutput)
	if err != nil {
		return &types.MsgValidateVrfResponse{
			Success: false,
		}, status.Error(codes.Internal, "error verifying proof")
	}

	if !valid {
		return &types.MsgValidateVrfResponse{
			Success: false,
		}, status.Error(codes.Internal, "error verifying proof")
	} else {
		return &types.MsgValidateVrfResponse{
			Success: true,
		}, nil
	}

}
