package keeper

import (
	"context"
	"encoding/json"
	"strconv"

	"cosmossdk.io/store/prefix"
	"github.com/ComputerKeeda/junction/x/junction/types"
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

	// till here, we have all the data required to store in the store
	// generate deterministic random number from the proof
	vrn, vrnErr := GenerateDeterministicRandomNumber(proof)
	if vrnErr != nil {
		return &types.MsgInitiateVrfResponse{
			Success: false,
		}, status.Error(codes.Internal, "error generating deterministic random number")
	}

	// find the station from id here
	_, err = k.getStationById(ctx, stationId)
	if err != nil {
		return &types.MsgInitiateVrfResponse{
			Success: false,
		}, status.Error(codes.NotFound, "station not found")
	}

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
	}

	vrfRecordByte := k.cdc.MustMarshal(&vrfRecord)
	vrfStore.Set(vrfStoreKeyByte, vrfRecordByte)

	return &types.MsgInitiateVrfResponse{
		Success: true,
	}, nil
}
