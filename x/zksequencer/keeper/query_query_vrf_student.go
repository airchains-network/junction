package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/zksequencer/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QueryVrfStudent(goCtx context.Context, req *types.QueryQueryVrfStudentRequest) (*types.QueryQueryVrfStudentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryQueryVrfStudentResponse{}, nil
}
