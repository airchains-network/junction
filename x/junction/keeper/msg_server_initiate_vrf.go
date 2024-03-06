package keeper

import (
	"context"
	"encoding/json"

	"github.com/ComputerKeeda/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) InitiateVrf(goCtx context.Context, msg *types.MsgInitiateVrf) (*types.MsgInitiateVrfResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

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

	vrn , vrnErr := GenerateDeterministicRandomNumber(proof)
	if vrnErr != nil {
		return &types.MsgInitiateVrfResponse{
			Success: false,
		}, status.Error(codes.Internal, "error generating deterministic random number")
	}

	// 

	// find the station from id here
	
	//  add code which will update the stations latest pod submission details

	_ = ctx

	return &types.MsgInitiateVrfResponse{}, nil
}
