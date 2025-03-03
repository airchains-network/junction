package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/vrf/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetCollections(goCtx context.Context, req *types.QueryGetCollectionsRequest) (*types.QueryGetCollectionsResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	collectionStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey))

	iterator := collectionStore.Iterator(nil, nil)
	defer func() {
		if err := iterator.Close(); err != nil {
			k.Logger().Error("failed to close iterator", "error", err)
		}
	}()

	var collections []types.Collection
	// Iterate over all stored collections.
	for ; iterator.Valid(); iterator.Next() {
		var collection types.Collection
		k.cdc.MustUnmarshal(iterator.Value(), &collection)
		collections = append(collections, collection)
	}

	return &types.QueryGetCollectionsResponse{Collections: collections}, nil

}
