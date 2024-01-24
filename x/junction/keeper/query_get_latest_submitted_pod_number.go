package keeper

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"

	"github.com/ComputerKeeda/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetLatestSubmittedPodNumber(goCtx context.Context, req *types.QueryGetLatestSubmittedPodNumberRequest) (*types.QueryGetLatestSubmittedPodNumberResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	figureDBStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FiguresDBPath))
	podSubmittedCountKey := fmt.Sprintf("pod-submitted-count__%s", req.StationId)
	podNumberByte := figureDBStore.Get([]byte(podSubmittedCountKey))
	if podNumberByte == nil {
		// in case of no pod count
		return &types.QueryGetLatestSubmittedPodNumberResponse{
			PodNumber: 0,
			Message:   "no submitted pods detected",
		}, nil
	}

	// Convert the byte slice to a string
	podNumString := string(podNumberByte)
	podNumber, err := strconv.ParseUint(podNumString, 10, 64)
	if err != nil {
		return &types.QueryGetLatestSubmittedPodNumberResponse{
			PodNumber: 0,
			Message:   err.Error(),
		}, nil
	}

	return &types.QueryGetLatestSubmittedPodNumberResponse{
		PodNumber: podNumber,
		Message:   "successfully retrieved latest submitted pod count",
	}, nil
}
