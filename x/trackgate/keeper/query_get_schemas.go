package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetSchemas(goCtx context.Context, req *types.QueryGetSchemasRequest) (*types.QueryGetSchemasResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	stationId := req.ExtTrackStationId
	schemaVersion := req.SchemaVersion
	var schema types.ExtTrackSchema

	if ValidateVersion(schemaVersion) == false {
		return nil, status.Error(codes.InvalidArgument, "invalid version")
	}

	dbPath := BuildExtTrackSchemaPath(stationId)
	trackSchemaStore := prefix.NewStore(storeAdapter, types.KeyPrefix(dbPath))
	key := []byte(schemaVersion)

	schemaData := trackSchemaStore.Get(key)
	if schemaData == nil {
		return nil, status.Error(codes.NotFound, "schema not found")
	}

	k.cdc.MustUnmarshal(schemaData, &schema)

	return &types.QueryGetSchemasResponse{
		Schema: &schema,
	}, nil
}
