package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/rollup/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRollups(goCtx context.Context, req *types.QueryGetRollupsRequest) (*types.QueryGetRollupsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rollups []types.RollupMetadata

	rollupDataStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RollupDataKey))

	pageRes, err := query.Paginate(rollupDataStore, req.Pagination, func(key []byte, value []byte) error {
		var rollupInfo types.RollupMetadata
		if err := k.cdc.Unmarshal(value, &rollupInfo); err != nil {
			return err
		}

		// Clear proverVerificationKey before adding to the response
		rollupInfo.ProverVerificationKey = nil

		rollups = append(rollups, rollupInfo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetRollupsResponse{Rollups: rollups, Pagination: pageRes}, nil
}
