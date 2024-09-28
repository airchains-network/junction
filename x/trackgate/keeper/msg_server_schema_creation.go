package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SchemaCreation handles the creation of a new schema for a specified external track station.
func (k msgServer) SchemaCreation(goCtx context.Context, msg *types.MsgSchemaCreation) (*types.MsgSchemaCreationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	extTrackStationId := msg.ExtTrackStationId
	version := msg.Version
	schema := msg.Schema

	if ValidateVersion(version) == false {
		return nil, status.Error(codes.InvalidArgument, "invalid version")
	}

	extTrackStationsDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExtTrackStationsDataStoreKey))
	stationDataByte := extTrackStationsDataDB.Get([]byte(extTrackStationId))
	if stationDataByte == nil {
		return &types.MsgSchemaCreationResponse{
			TrackKey: "",
			Status:   false,
		}, status.Error(codes.NotFound, "station not found")
	}

	var stationData types.ExtTrackStations
	k.cdc.MustUnmarshal(stationDataByte, &stationData)

	var stationTrackData types.SequencerDetails
	errCheck := json.Unmarshal(stationData.SequencerDetails, &stationTrackData)
	if errCheck != nil {
		return nil, status.Error(codes.Internal, "something went wrong")
	}

	trackName := stationTrackData.Name
	// Generate UUID
	trackKey := uuid.New().String()
	newExtTrackSchema := types.ExtTrackSchema{
		TrackName:         trackName,
		ExtTrackStationId: extTrackStationId,
		Version:           version,
		TrackKey:          trackKey,
		Schema:            schema,
	}

	/*
		Storing pattern
		ext_track_schema/ext_track_station_id : and in this database this is how data will be stored
		key: version
	*/
	storingData := k.cdc.MustMarshal(&newExtTrackSchema)
	dbPath := BuildExtTrackSchemaPath(extTrackStationId)
	trackSchemaStore := prefix.NewStore(storeAdapter, types.KeyPrefix(dbPath))
	key := []byte(version)
	trackSchemaStore.Set(key, storingData)

	return &types.MsgSchemaCreationResponse{
		TrackKey: trackKey,
		Status:   true,
	}, nil
}
