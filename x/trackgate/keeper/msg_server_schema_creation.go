package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SchemaCreation handles the creation of a new schema for a specified external track station.
func (k msgServer) SchemaCreation(goCtx context.Context, msg *types.MsgSchemaCreation) (*types.MsgSchemaCreationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	creator := msg.Creator
	extTrackStationId := msg.ExtTrackStationId
	version := msg.Version
	schema := msg.Schema

	if ValidateVersion(version) == false {
		return nil, status.Error(codes.InvalidArgument, "invalid version")
	}

	// Initialising the accurate store for access
	dbPath := BuildExtTrackSchemaPath(extTrackStationId)
	trackSchemaStore := prefix.NewStore(storeAdapter, types.KeyPrefix(dbPath))
	key := []byte(version)
	// Checking whether the schema for this version is already created or not
	uniqueSchema := trackSchemaStore.Get(key)
	if uniqueSchema != nil {
		return nil, status.Error(codes.AlreadyExists, "schema already exists for this version")
	}

	extTrackStationsDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExtTrackStationsDataStoreKey))
	extTrackStationIdByte := []byte(extTrackStationId)
	stationDataByte := extTrackStationsDataDB.Get(extTrackStationIdByte)
	if stationDataByte == nil {
		return &types.MsgSchemaCreationResponse{
			SchemaKey: "",
			Status:    false,
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
	// Generate schemaKey
	schemaKeyBytes := sha256.Sum256([]byte(creator + extTrackStationId + version))
	schemaKey := hex.EncodeToString(schemaKeyBytes[:])

	newExtTrackSchema := types.ExtTrackSchema{
		TrackName:         trackName,
		ExtTrackStationId: extTrackStationId,
		Version:           version,
		SchemaKey:         schemaKey,
		Schema:            schema,
	}

	if stationData.StationSchemaKey == "none" {
		// update codes
		updateStationDetails := types.ExtTrackStations{
			Operators:            stationData.Operators,
			LatestPod:            stationData.LatestPod,
			LatestMerkleRootHash: stationData.LatestMerkleRootHash,
			Name:                 stationData.Name,
			Id:                   stationData.Id,
			StationType:          stationData.StationType,
			FheEnabled:           stationData.FheEnabled,
			SequencerDetails:     stationData.SequencerDetails,
			DaDetails:            stationData.DaDetails,
			ProverDetails:        stationData.ProverDetails,
			StationSchemaKey:     schemaKey,
		}
		byteStationId := []byte(stationData.Id)
		byteUpdateStationDetails := k.cdc.MustMarshal(&updateStationDetails)
		extTrackStationsDataDB.Set(byteStationId, byteUpdateStationDetails)
	}

	/*
		Here we will add the version to the database of version finder
		Here is what it will look like
		dbPath will be ext_track_version_finder/{station-ID}
		key: schemaKey and value:version
	*/
	versionFinderDbPath := BuildExtTrackVersionFinderStoreKey(extTrackStationId)
	versionFinderStore := prefix.NewStore(storeAdapter, types.KeyPrefix(versionFinderDbPath))
	versionFinderKey := []byte(schemaKey)
	versionFinderValue := []byte(version)
	versionFinderStore.Set(versionFinderKey, versionFinderValue)
	/*
		Storing pattern
		ext_track_schema/ext_track_station_id : and in this database this is how data will be stored
		key: version
	*/
	storingData := k.cdc.MustMarshal(&newExtTrackSchema)
	trackSchemaStore.Set(key, storingData)

	return &types.MsgSchemaCreationResponse{
		SchemaKey: schemaKey,
		Status:    true,
	}, nil
}
