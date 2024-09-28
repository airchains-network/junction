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

func (k Keeper) ListSchemas(goCtx context.Context, req *types.QueryListSchemasRequest) (*types.QueryListSchemasResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	stationId := req.ExtTrackStationId
	var schemas []types.ExtTrackSchema
	// Initialising the accurate store for access
	dbPath := BuildExtTrackSchemaPath(stationId)
	trackSchemaStore := prefix.NewStore(storeAdapter, types.KeyPrefix(dbPath))

	pageResponse, err := query.Paginate(trackSchemaStore, req.Pagination, func(key []byte, value []byte) error {
		var singleSchemaData types.ExtTrackSchema
		if err := k.cdc.Unmarshal(value, &singleSchemaData); err != nil {
			return err
		}

		schemas = append(schemas, singleSchemaData)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListSchemasResponse{Schemas: schemas, Pagination: pageResponse}, nil
}
