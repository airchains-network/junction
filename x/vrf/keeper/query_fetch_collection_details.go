package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"fmt"
	"github.com/airchains-network/junction/x/vrf/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FetchCollectionDetails(goCtx context.Context, req *types.QueryFetchCollectionDetailsRequest) (*types.QueryFetchCollectionDetailsResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	collectionStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey))
	vrfStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey+"/"+req.CollectionId))

	collectionKey := []byte(req.CollectionId)

	// Check if the collection exists.
	collectionBytes := collectionStore.Get(collectionKey)
	if collectionBytes == nil {
		errMsg := fmt.Sprintf("Collection '%s' does not exist", req.CollectionId)
		return &types.QueryFetchCollectionDetailsResponse{}, status.Error(codes.NotFound, errMsg)
	}

	// Unmarshal the collection details.
	var collectionDetails types.Collection
	k.cdc.MustUnmarshal(collectionBytes, &collectionDetails)

	iterator := vrfStore.Iterator(nil, nil)
	defer func(iterator storetypes.Iterator) {
		err := iterator.Close()
		if err != nil {
			k.Logger().Error("failed to close iterator", "error", err)
		}
	}(iterator)

	var vrfDetails []types.VrfRecord
	// Iterate over all stored vrf records.
	for ; iterator.Valid(); iterator.Next() {
		var vrfRecord types.VrfRecord
		k.cdc.MustUnmarshal(iterator.Value(), &vrfRecord)
		vrfDetails = append(vrfDetails, vrfRecord)
	}

	return &types.QueryFetchCollectionDetailsResponse{
		Details:    &collectionDetails,
		VrfRecords: vrfDetails,
	}, nil
	
}
