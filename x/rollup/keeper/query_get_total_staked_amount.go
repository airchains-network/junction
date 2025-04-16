package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/rollup/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetTotalStakedAmount(goCtx context.Context, req *types.QueryGetTotalStakedAmountRequest) (*types.QueryGetTotalStakedAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	totalStakedAmountResponse, err := k.Inner_GetTotalStakedAmount(ctx)
	if err != nil {
		return nil, err
	}

	response := &types.QueryGetTotalStakedAmountResponse{
		TotalStakedAmount: int64(totalStakedAmountResponse.TotalStakedAmount),
	}

	// Map creators to proto message format
	creators := make([]*types.Creator, 0, len(totalStakedAmountResponse.Creators))
	for _, creator := range totalStakedAmountResponse.Creators {
		protoCreator := &types.Creator{
			CreatorAddress: creator.CreatorAddress,
		}

		rollups := make([]*types.RollupStake, 0, len(creator.Rollups))
		for _, rollup := range creator.Rollups {
			rollups = append(rollups, &types.RollupStake{
				RollupId:     rollup.RollupID,
				AmountStaked: int64(rollup.AmountStaked),
				Denom:        rollup.Denom,
			})
		}

		protoCreator.Rollups = rollups
		creators = append(creators, protoCreator)
	}

	response.Creators = creators

	return response, nil
}
