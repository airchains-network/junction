package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/zksequencer/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CollegeAll(ctx context.Context, req *types.QueryAllCollegeRequest) (*types.QueryAllCollegeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var colleges []types.College

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	collegeStore := prefix.NewStore(store, types.KeyPrefix(types.CollegeKey))

	pageRes, err := query.Paginate(collegeStore, req.Pagination, func(key []byte, value []byte) error {
		var college types.College
		if err := k.cdc.Unmarshal(value, &college); err != nil {
			return err
		}

		colleges = append(colleges, college)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCollegeResponse{College: colleges, Pagination: pageRes}, nil
}

func (k Keeper) College(ctx context.Context, req *types.QueryGetCollegeRequest) (*types.QueryGetCollegeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	college, found := k.GetCollege(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetCollegeResponse{College: college}, nil
}
