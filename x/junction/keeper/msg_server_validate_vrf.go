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

	if station.LatestPod+1 != podNumber {
		return &types.MsgValidateVrfResponse{
			Success: false,
		}, sdkerrors.ErrInvalidHeight
	}

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
		//update the vrf details
		var updateVrfRecord types.VrfRecord
		updateVrfRecord.VrfCreatorAddr = vrfDetails.VrfCreatorAddr
		updateVrfRecord.VrfVerifierAddr = creator
		updateVrfRecord.PodNumber = vrfDetails.PodNumber
		updateVrfRecord.StationId = vrfDetails.StationId
		updateVrfRecord.Occupancy = vrfDetails.Occupancy
		updateVrfRecord.CreatorsVrfKey = vrfDetails.CreatorsVrfKey
		updateVrfRecord.SerializedRcFromCreator = vrfDetails.SerializedRcFromCreator
		updateVrfRecord.SerializedRcFromVerifier = serializedRc
		updateVrfRecord.Proof = vrfDetails.Proof
		updateVrfRecord.VrfOutput = vrfDetails.VrfOutput
		updateVrfRecord.IsVerified = true
		updateVrfRecord.Vrn = vrfDetails.Vrn
		updateVrfRecord.SelectedTrackIndex = vrfDetails.SelectedTrackIndex

		vrfRecordByte := k.cdc.MustMarshal(&updateVrfRecord)
		vrfStore.Set(vrfStoreKeyByte, vrfRecordByte)

		stationTrack := station.Tracks
		spsp := stationTrack[vrfDetails.SelectedTrackIndex]

		updateStation := types.Stations{
			Tracks:               stationTrack,
			VotingPower:          station.VotingPower,
			LatestPod:            station.LatestPod,
			LatestMerkleRootHash: station.LatestMerkleRootHash,
			VerificationKey:      station.VerificationKey,
			StationInfo:          station.StationInfo,
			Id:                   station.Id,
			Creator:              station.Creator,
			Spsp:                 spsp,
		}
		stationDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StationDataKey))
		byteStationId := []byte(station.Id)
		byteStation := k.cdc.MustMarshal(&updateStation)
		stationDataDB.Set(byteStationId, byteStation)

		return &types.MsgValidateVrfResponse{
			Success: true,
		}, nil
	}

}
