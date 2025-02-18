package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/zksequencer/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QueryVrfStudent2(goCtx context.Context, req *types.QueryQueryVrfStudent2Request) (*types.QueryQueryVrfStudent2Response, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	student, err := k.GetStudentFromVRF(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryQueryVrfStudent2Response{
		Name: student.Name, Details: student.Details,
	}, nil
}
