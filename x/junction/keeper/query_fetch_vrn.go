package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/junction/types"
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

	var podNumber = req.PodNumber
	var stationId = req.StationId

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	vrfStoreKey, vrfStoreKeyByte := GetVRFKeyByte(stationId, podNumber) // "vrf/{stationId}/{podNumber}
	vrfStore := prefix.NewStore(storeAdapter, types.KeyPrefix(vrfStoreKey))

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
