package keeper

import (
	"context"
	"fmt"
	"github.com/ComputerKeeda/junction/x/junction/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
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

	// submit pod
	Error := k.SubmitPodHelper(
		ctx,
		msg,
	)
	if Error != nil {
		return &types.MsgSubmitPodResponse{
			PodStatus: true,
		}, Error
	}

	// update pod submitted count
	figureDBStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FiguresDBPath))
	podSubmittedCountKey := fmt.Sprintf("pod-submitted-count__%s", station.Id)
	newPodCount := strconv.FormatUint(msg.PodNumber, 10)
	figureDBStore.Set([]byte(podSubmittedCountKey), []byte(newPodCount))

	return &types.MsgSubmitPodResponse{
		PodStatus: true,
	}, nil
}
