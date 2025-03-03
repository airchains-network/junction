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

func (k msgServer) RegisterCollection(goCtx context.Context, msg *types.MsgRegisterCollection) (*types.MsgRegisterCollectionResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	collectionDb := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey))

	collectionKey := []byte(msg.CollectionId)

	// Check if the collection already exists
	if collectionDetailsBytes := collectionDb.Get(collectionKey); collectionDetailsBytes != nil {
		errMsg := fmt.Sprintf("Collection '%s' already exists", msg.CollectionId)
		return nil, status.Error(codes.AlreadyExists, errMsg)
	}

	// Create and populate collection details
	collectionDetails := types.Collection{
		CollectionOwner:    msg.CollectionOwner,
		CollectionName:     msg.CollectionName,
		CollectionId:       msg.CollectionId,
		Members:            msg.Members,
		Offset:             msg.Offset,
		Index:              msg.Offset,
		GreyListed:         false,
		DisputeVoteYesHash: "none",
		DisputeVoteNoHash:  "none",
	}

	// Marshal collection details
	collectionDetailsBytes, err := k.cdc.Marshal(&collectionDetails)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to marshal collection details for collection '%s': %v", msg.CollectionId, err)
		return nil, status.Errorf(codes.Internal, errMsg)
	}

	// Save collection details
	collectionDb.Set(collectionKey, collectionDetailsBytes)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent("collection_registered",
			sdk.NewAttribute("collection_id", msg.CollectionId),
			sdk.NewAttribute("collection_owner", msg.CollectionOwner),
			sdk.NewAttribute("collection_name", msg.CollectionName),
			sdk.NewAttribute("members", fmt.Sprintf("%v", msg.Members)),
			sdk.NewAttribute("offset", fmt.Sprintf("%d", msg.Offset)),
		),
	)

	return &types.MsgRegisterCollectionResponse{
		Result: true,
	}, nil

}
