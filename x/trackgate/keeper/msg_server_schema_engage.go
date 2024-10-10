package keeper

import (
	"context"
	"encoding/json"
	"strconv"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SchemaEngage(goCtx context.Context, msg *types.MsgSchemaEngage) (*types.MsgSchemaEngageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	operator := msg.Operator
	extTrackStationId := msg.ExtTrackStationId

	schemaObject := msg.SchemaObject
	acknowledgementHash := msg.AcknowledgementHash
	podNumber := msg.PodNumber
	sequencerDetails := msg.SequencerDetails

	extTrackStationsDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExtTrackStationsDataStoreKey))
	extTrackStationIdByte := []byte(extTrackStationId)
	stationDataByte := extTrackStationsDataDB.Get(extTrackStationIdByte)
	if stationDataByte == nil {
		return nil, status.Errorf(codes.NotFound, "ext track station %s not found", extTrackStationId)
	}

	var stationData types.ExtTrackStations
	k.cdc.MustUnmarshal(stationDataByte, &stationData)

	// Station Metrics Check
	stationMetricsDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TrackGateFigureStoreKey))
	stationMetricsDataBytes := stationMetricsDB.Get(extTrackStationIdByte)
	if stationMetricsDataBytes == nil {
		return nil, status.Errorf(codes.NotFound, "ext track station metrics %s not found", extTrackStationId)
	}

	var stationMetrics types.StationMetrics
	k.cdc.MustUnmarshal(stationMetricsDataBytes, &stationMetrics)

	schemaKey := stationData.StationSchemaKey
	// check operator
	operators := stationData.Operators
	if !ContainsOperator(operators, operator) {
		return nil, status.Errorf(codes.PermissionDenied, "operator %s not found in ext track station %s", operator, extTrackStationId)
	}

	// check pod number
	if stationData.LatestPod+1 != podNumber {
		return nil, status.Errorf(codes.PermissionDenied, "pod number %d not match with ext track station %s", podNumber, extTrackStationId)
	}

	versionFinderDbPath := BuildExtTrackVersionFinderStoreKey(extTrackStationId)
	versionFinderStore := prefix.NewStore(storeAdapter, types.KeyPrefix(versionFinderDbPath))
	versionFinderKey := []byte(schemaKey)
	versionByte := versionFinderStore.Get(versionFinderKey)
	if versionByte == nil {
		return nil, status.Errorf(codes.NotFound, "schema %s not found in ext track station %s", schemaKey, extTrackStationId)
	}

	dbPath := BuildExtTrackSchemaPath(extTrackStationId)
	trackSchemaStore := prefix.NewStore(storeAdapter, types.KeyPrefix(dbPath))

	var schema types.ExtTrackSchema
	schemaDataByte := trackSchemaStore.Get(versionByte)

	if schemaDataByte == nil {
		return nil, status.Errorf(codes.NotFound, "schema %s not found in ext track station %s", schemaKey, extTrackStationId)
	}

	k.cdc.MustUnmarshal(schemaDataByte, &schema)

	var schemaDef types.SchemaDef
	if err := json.Unmarshal(schema.Schema, &schemaDef); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to unmarshal schema %s", schemaKey)
	}

	schemaObjectString := string(schemaObject)

	_, err := DynamicUnmarshal(schemaDef, schemaObjectString)
	if err != nil {
		return nil, err
	}

	newSchemaEngagement := types.ExtTrackSchemaEngagement{
		ExtTrackStationId:   extTrackStationId,
		EngageBy:            operator,
		AcknowledgementHash: acknowledgementHash,
		PodNumber:           podNumber,
		StationName:         stationData.Name,
		TrackName:           schema.TrackName,
		SchemaVersion:       schema.Version,
		SchemaKey:           schemaKey,
		SchemaObject:        schemaObject,
		SequencerDetails:    sequencerDetails,
		IsVerified:          false,
		VerifiedBy:          "none",
	}

	newSchemaEngagementDbPath := BuildExtTrackSchemaEngagementsStoreKey(extTrackStationId)
	newSchemaEngagementStore := prefix.NewStore(storeAdapter, types.KeyPrefix(newSchemaEngagementDbPath))
	podNumberString := strconv.FormatUint(podNumber, 10)
	newEngagementKey := []byte(podNumberString)

	check := newSchemaEngagementStore.Get(newEngagementKey)
	if check != nil {
		return nil, status.Errorf(codes.AlreadyExists, "schema engagement already exists for pod number %d", podNumber)
	}

	newEngagementValue := k.cdc.MustMarshal(&newSchemaEngagement)
	newSchemaEngagementStore.Set(newEngagementKey, newEngagementValue)
	// update codes
	updateStationDetails := types.ExtTrackStations{
		Creator:                   stationData.Creator,
		Operators:                 stationData.Operators,
		LatestPod:                 podNumber,
		LatestAcknowledgementHash: acknowledgementHash,
		Name:                      stationData.Name,
		Id:                        stationData.Id,
		StationType:               stationData.StationType,
		FheEnabled:                stationData.FheEnabled,
		SequencerDetails:          stationData.SequencerDetails,
		DaDetails:                 stationData.DaDetails,
		ProverDetails:             stationData.ProverDetails,
		StationSchemaKey:          schemaKey,
	}

	byteStationId := []byte(stationData.Id)
	byteUpdateStationDetails := k.cdc.MustMarshal(&updateStationDetails)

	extTrackStationsDataDB.Set(byteStationId, byteUpdateStationDetails)

	updatedStationMetrics := types.StationMetrics{
		TotalPodCount:         stationMetrics.TotalPodCount + 1,
		TotalSchemaCount:      stationMetrics.TotalSchemaCount,
		TotalMigrationCount:   stationMetrics.TotalMigrationCount,
		TotalVerifiedPodCount: stationMetrics.TotalVerifiedPodCount,
	}
	stationMetricsValue := k.cdc.MustMarshal(&updatedStationMetrics)
	stationMetricsDB.Set(extTrackStationIdByte, stationMetricsValue)

	return &types.MsgSchemaEngageResponse{
		Status: true,
	}, nil
}
