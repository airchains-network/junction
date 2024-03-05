package keeper

import (
	"context"

	"github.com/ComputerKeeda/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) InitStation(goCtx context.Context, msg *types.MsgInitStation) (*types.MsgInitStationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var verificationKey = msg.VerificationKey
	var stationID = msg.StationId
	var stationInfo = msg.StationInfo

	var newStation = types.Stations{
		Tracks:               msg.Tracks,
		VotingPower:          []uint64{100},
		LatestPod:            0,
		LatestMerkleRootHash: "0",
		VerificationKey:      verificationKey,
		StationInfo:          stationInfo,
		Id:                   stationID,
		Creator:              msg.Creator,
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
