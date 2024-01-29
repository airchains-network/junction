package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) VerifyPod(goCtx context.Context, msg *types.MsgVerifyPod) (*types.MsgVerifyPodResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	Error := k.VerifyPodHelper(
		ctx,
		msg,
	)
	if Error != nil {
		return &types.MsgVerifyPodResponse{
			IsVerified: false,
			Message:    Error.Error(),
		}, Error
	}

	return &types.MsgVerifyPodResponse{
		IsVerified: true,
		Message:    "",
	}, nil
}
