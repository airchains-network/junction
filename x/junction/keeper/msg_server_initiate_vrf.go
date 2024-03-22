package keeper

import (
	"context"
	"encoding/json"
	"math/big"
	"strconv"

	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/junction/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) InitiateVrf(goCtx context.Context, msg *types.MsgInitiateVrf) (*types.MsgInitiateVrfResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	vrfCreator := msg.Creator
	podNumber := msg.PodNumber
	stationId := msg.StationId
	occupancy := msg.Occupancy
	creatorsVrfKey := msg.CreatorsVrfKey
	extraArg := msg.ExtraArg

	var unmarshalledExtraArg ExtraArg
	err := json.Unmarshal(extraArg, &unmarshalledExtraArg)
	if err != nil {
		return &types.MsgInitiateVrfResponse{
			Success: false,
		}, status.Error(codes.InvalidArgument, "incorrect extra data argument")
	}

	serializedRc := unmarshalledExtraArg.SerializedRc
	proof := unmarshalledExtraArg.Proof
	vrfOutput := unmarshalledExtraArg.VrfOutput

	// check if the pod number associated with the station is valid
	if podNumber < 1 {
		return &types.MsgInitiateVrfResponse{
			Success: false,
		}, status.Error(codes.InvalidArgument, "invalid pod number")
	}

	// find the station from id here
	station, err := k.getStationById(ctx, stationId)
	if err != nil {
		return &types.MsgInitiateVrfResponse{
			Success: false,
		}, status.Error(codes.NotFound, "station not found")
	}
	if station.LatestPod+1 != podNumber {
		return &types.MsgInitiateVrfResponse{
			Success: false,
		}, status.Error(codes.InvalidArgument, "invalid pod number")
	}
	// checking if the creator is valid track member or not
	vtmFound := false
	validTrackMembers := station.Tracks
	for _, trackMember := range validTrackMembers {
		if trackMember == vrfCreator {
			vtmFound = true
			break
		}
	}

	if !vtmFound {
		return &types.MsgInitiateVrfResponse{
			Success: false,
		}, status.Error(codes.PermissionDenied, "creator is not a valid track member")
	}

	// till here, we have all the data required to store in the store
	// *generate deterministic random number from the proof*
	vrn, vrnErr := GenerateDeterministicRandomNumber(proof)
	if vrnErr != nil {
		return &types.MsgInitiateVrfResponse{
			Success: false,
		}, status.Error(codes.Internal, "error generating deterministic random number")
	}
	vrnBigInt := new(big.Int).SetBytes(vrn)

	// store the data of this vrf, so that verifier can verify the data in next transaction
	vrfStoreKey, vrfStoreKeyByte := GetVRFKeyByte(stationId, podNumber) // "vrf/{stationId}/{podNumber}
	vrfStore := prefix.NewStore(storeAdapter, types.KeyPrefix(vrfStoreKey))

	vrfDetailsByte := vrfStore.Get(vrfStoreKeyByte)
	if vrfDetailsByte != nil {
		// vrf details already present
		return &types.MsgInitiateVrfResponse{
			Success: false,
		}, status.Error(codes.AlreadyExists, "vrf details already present")
	}

	occupancyBigInt := new(big.Int).SetUint64(occupancy)
	selectedTrackIndex := new(big.Int).Mod(vrnBigInt, occupancyBigInt).Uint64()

	// data is not present in the store, so we can store the data
	vrfRecord := types.VrfRecord{
		VrfCreatorAddr:           vrfCreator,
		VrfVerifierAddr:          "",
		PodNumber:                strconv.FormatUint(podNumber, 10),
		StationId:                stationId,
		Occupancy:                occupancy,
		CreatorsVrfKey:           creatorsVrfKey,
		SerializedRcFromCreator:  serializedRc,
		SerializedRcFromVerifier: nil,
		Proof:                    proof,
		VrfOutput:                vrfOutput,
		IsVerified:               false,
		Vrn:                      vrn,
		SelectedTrackIndex:       selectedTrackIndex,
	}

	vrfRecordByte := k.cdc.MustMarshal(&vrfRecord)
	vrfStore.Set(vrfStoreKeyByte, vrfRecordByte)

	return &types.MsgInitiateVrfResponse{
		Success: true,
	}, nil
}
