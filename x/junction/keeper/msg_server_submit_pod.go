package keeper

import (
	"context"
	"github.com/ComputerKeeda/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SubmitPod(goCtx context.Context, msg *types.MsgSubmitPod) (*types.MsgSubmitPodResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var stationId = msg.StationId

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
		}, sdkerrors.ErrInvalidHeight
	}

	Error := k.SubmitPodHelper(
		ctx,
		msg,
	)

	if Error != nil {
		return &types.MsgSubmitPodResponse{
			PodStatus: true,
		}, Error
	}

	return &types.MsgSubmitPodResponse{
		PodStatus: true,
	}, nil
}
