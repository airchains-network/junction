package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ComputerKeeda/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetStationDetailsByAddress(goCtx context.Context, req *types.QueryGetStationDetailsByAddressRequest) (*types.QueryGetStationDetailsByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	address := req.Address
	stationId, found := k.GetStationIdByAddressHelper(ctx, address)
	if stationId == "nil" || !found {
		// No station found on this address
		return nil, sdkerrors.ErrKeyNotFound
	}
	station, err := k.getStationById(ctx, stationId)
	if err != nil {
		return nil, err
	}
	return &types.QueryGetStationDetailsByAddressResponse{
		Station: &station,
	}, nil
}
