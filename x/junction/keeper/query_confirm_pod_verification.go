package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ConfirmPodVerification(goCtx context.Context, req *types.QueryConfirmPodVerificationRequest) (*types.QueryConfirmPodVerificationResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	Error := k.ConfirmPodVerificationHelper(
		ctx,
		req,
	)
	if Error != nil {
		return &types.QueryConfirmPodVerificationResponse{
			IsVerified: false,
			Message:    Error.Error(),
		}, Error
	}

	return &types.QueryConfirmPodVerificationResponse{
		IsVerified: true,
		Message:    "pod verified successfully",
	}, nil

}
