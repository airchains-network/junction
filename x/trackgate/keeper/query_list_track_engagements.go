package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListTrackEngagements(goCtx context.Context, req *types.QueryListTrackEngagementsRequest) (*types.QueryListTrackEngagementsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	extTrackStationId := req.ExtTrackStationId
	schemaEngagementDbPath := BuildExtTrackSchemaEngagementsStoreKey(extTrackStationId)
	schemaEngagementStore := prefix.NewStore(storeAdapter, types.KeyPrefix(schemaEngagementDbPath))

	var schemaEngagements []types.ExtTrackSchemaEngagement
	pageResponse, err := query.Paginate(schemaEngagementStore, req.Pagination, func(key []byte, value []byte) error {
		var singleSchemaEngagements types.ExtTrackSchemaEngagement
		if err := k.cdc.Unmarshal(value, &singleSchemaEngagements); err != nil {
			return err
		}

		schemaEngagements = append(schemaEngagements, singleSchemaEngagements)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListTrackEngagementsResponse{
		Engagements: schemaEngagements,
		Pagination:  pageResponse,
	}, nil
}
