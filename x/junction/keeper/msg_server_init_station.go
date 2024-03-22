package keeper

import (
	"context"
	"encoding/json"
	"github.com/airchains-network/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) InitStation(goCtx context.Context, msg *types.MsgInitStation) (*types.MsgInitStationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var verificationKey = msg.VerificationKey
	var stationID = msg.StationId
	var stationInfo = msg.StationInfo
	var extraArgs = msg.ExtraArg

	var stationArgs types.StationArg
	stationArgsUnmarshalErr := json.Unmarshal(extraArgs, &stationArgs)
	if stationArgsUnmarshalErr != nil {
		return &types.MsgInitStationResponse{
			StationId: "nil",
			Status:    false,
		}, stationArgsUnmarshalErr
	}

	var newStation = types.Stations{
		Tracks:               msg.Tracks,
		VotingPower:          msg.TracksVotingPower,
		LatestPod:            0,
		LatestMerkleRootHash: "0",
		VerificationKey:      verificationKey,
		StationInfo:          stationInfo,
		Id:                   stationID,
		Creator:              msg.Creator,
		Spsp:                 "nil",
		DaType:               stationArgs.DaType,
		TrackType:            stationArgs.TrackType,
		Prover:               stationArgs.Prover,
	}

	Error := k.initStationHelper(ctx, newStation, msg.Creator)
	if Error != nil {
		return &types.MsgInitStationResponse{
			StationId: "nil",
			Status:    false,
		}, Error
	}

	return &types.MsgInitStationResponse{
		StationId: stationID,
		Status:    true,
	}, nil
}
