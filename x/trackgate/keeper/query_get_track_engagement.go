package keeper

import (
	"context"
	"strconv"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetTrackEngagement(goCtx context.Context, req *types.QueryGetTrackEngagementRequest) (*types.QueryGetTrackEngagementResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	extTrackStationId := req.ExtTrackStationId
	podNumber := req.PodNumber

	schemaEngagementDbPath := BuildExtTrackSchemaEngagementsStoreKey(extTrackStationId)
	schemaEngagementStore := prefix.NewStore(storeAdapter, types.KeyPrefix(schemaEngagementDbPath))
	podNumberString := strconv.FormatUint(podNumber, 10)
	engagementKey := []byte(podNumberString)

	if !schemaEngagementStore.Has(engagementKey) {
		return nil, status.Error(codes.NotFound, "schema not found")
	}

	var schemaEngagement types.ExtTrackSchemaEngagement
	schemaEngagementByte := schemaEngagementStore.Get(engagementKey)

	k.cdc.MustUnmarshal(schemaEngagementByte, &schemaEngagement)

	return &types.QueryGetTrackEngagementResponse{
		Engagement: &schemaEngagement,
	}, nil
}
