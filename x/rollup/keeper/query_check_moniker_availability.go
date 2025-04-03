package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/rollup/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CheckMonikerAvailability verifies whether the provided moniker is available.
// If the request is nil or the moniker is empty, it returns an error.
// It unwraps the SDK context from the given context, queries the store to check if
// the moniker already exists, and returns a response where IsAvailable is true if
// the moniker does not exist.
func (k Keeper) CheckMonikerAvailability(goCtx context.Context, req *types.QueryCheckMonikerAvailabilityRequest) (*types.QueryCheckMonikerAvailabilityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request: request cannot be nil")
	}

	if req.Moniker == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request: moniker cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the moniker already exists.
	monikerExists := k.GetMoniker(ctx, req.Moniker)

	// Moniker is available if it does not already exist.
	isAvailable := !monikerExists

	return &types.QueryCheckMonikerAvailabilityResponse{
		IsAvailable: isAvailable,
	}, nil
}
