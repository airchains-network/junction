package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"encoding/binary"
	"github.com/airchains-network/junction/x/vrf/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FetchVrn(goCtx context.Context, req *types.QueryFetchVrnRequest) (*types.QueryFetchVrnResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var keyIndex = req.KeyIndex
	var collectionId = req.CollectionId

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	vrfStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey+"/"+collectionId))
	vrfStoreKeyByte := make([]byte, 4)
	binary.LittleEndian.PutUint32(vrfStoreKeyByte, keyIndex)

	vrfDetailsByte := vrfStore.Get(vrfStoreKeyByte)
	// check if vrf details exists
	if vrfDetailsByte == nil {
		return &types.QueryFetchVrnResponse{
			Details: nil,
		}, status.Error(codes.NotFound, "vrf details not found")
	}

	var vrfDetails types.VrfRecord
	k.cdc.MustUnmarshal(vrfDetailsByte, &vrfDetails)
	return &types.QueryFetchVrnResponse{
		Details: &vrfDetails,
	}, nil

}
