package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"fmt"
	"github.com/airchains-network/junction/x/vrf/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FetchCollection(goCtx context.Context, req *types.QueryFetchCollectionRequest) (*types.QueryFetchCollectionResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	// Connect to collection store
	collectionStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey))

	collectionKey := []byte(req.CollectionId)

	// Check if the collection exists.
	collectionBytes := collectionStore.Get(collectionKey)
	if collectionBytes == nil {
		errMsg := fmt.Sprintf("Collection '%s' does not exist", req.CollectionId)
		return &types.QueryFetchCollectionResponse{}, status.Error(codes.NotFound, errMsg)
	}

	// Unmarshal the collection details.
	var collection types.Collection
	k.cdc.MustUnmarshal(collectionBytes, &collection)

	return &types.QueryFetchCollectionResponse{
		Details: &collection,
	}, nil
	
}
