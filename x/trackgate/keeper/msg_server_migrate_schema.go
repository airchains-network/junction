package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MigrateSchema(goCtx context.Context, msg *types.MsgMigrateSchema) (*types.MsgMigrateSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	operator := msg.Operator
	extTrackStationId := msg.ExtTrackStationId
	newSchemaKey := msg.NewSchemaKey

	extTrackStationsDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExtTrackStationsDataStoreKey))
	extTrackStationIdByte := []byte(extTrackStationId)

	// Station Metrics Check
	stationMetricsDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TrackGateFigureStoreKey))
	stationMetricsDataBytes := stationMetricsDB.Get(extTrackStationIdByte)
	if stationMetricsDataBytes == nil {
		return nil, status.Errorf(codes.NotFound, "ext track station metrics %s not found", extTrackStationId)
	}

	var stationMetrics types.StationMetrics
	k.cdc.MustUnmarshal(stationMetricsDataBytes, &stationMetrics)

	stationDataByte := extTrackStationsDataDB.Get(extTrackStationIdByte)
	if stationDataByte == nil {
		return nil, status.Errorf(codes.NotFound, "ext track station %s not found", extTrackStationId)
	}
	var stationData types.ExtTrackStations
	k.cdc.MustUnmarshal(stationDataByte, &stationData)

	// check operator
	operators := stationData.Operators
	if !ContainsOperator(operators, operator) {
		return nil, status.Errorf(codes.PermissionDenied, "operator %s not found in ext track station %s", operator, extTrackStationId)
	}

	// Checking whether the new-scheme key is valid or not
	versionFinderDbPath := BuildExtTrackVersionFinderStoreKey(extTrackStationId)
	versionFinderStore := prefix.NewStore(storeAdapter, types.KeyPrefix(versionFinderDbPath))
	versionFinderKey := []byte(newSchemaKey)
	versionByte := versionFinderStore.Get(versionFinderKey)
	if versionByte == nil {
		return nil, status.Errorf(codes.NotFound, "schema %s not found in ext track station %s", newSchemaKey, extTrackStationId)
	}

	dbPath := BuildExtTrackSchemaPath(extTrackStationId)
	trackSchemaStore := prefix.NewStore(storeAdapter, types.KeyPrefix(dbPath))

	schemaDataByte := trackSchemaStore.Get(versionByte)

	if schemaDataByte == nil {
		return nil, status.Errorf(codes.NotFound, "schema %s not found in ext track station %s", newSchemaKey, extTrackStationId)
	}

	// update codes
	updateStationDetails := types.ExtTrackStations{
		Operators:                 stationData.Operators,
		LatestPod:                 stationData.LatestPod,
		LatestAcknowledgementHash: stationData.LatestAcknowledgementHash,
		Name:                      stationData.Name,
		Id:                        stationData.Id,
		StationType:               stationData.StationType,
		FheEnabled:                stationData.FheEnabled,
		SequencerDetails:          stationData.SequencerDetails,
		DaDetails:                 stationData.DaDetails,
		ProverDetails:             stationData.ProverDetails,
		StationSchemaKey:          newSchemaKey,
	}

	byteStationId := []byte(stationData.Id)
	byteUpdateStationDetails := k.cdc.MustMarshal(&updateStationDetails)

	extTrackStationsDataDB.Set(byteStationId, byteUpdateStationDetails)

	updatedStationMetrics := types.StationMetrics{
		TotalPodCount:         stationMetrics.TotalPodCount,
		TotalSchemaCount:      stationMetrics.TotalSchemaCount,
		TotalMigrationCount:   stationMetrics.TotalMigrationCount + 1,
		TotalVerifiedPodCount: stationMetrics.TotalVerifiedPodCount,
	}
	stationMetricsValue := k.cdc.MustMarshal(&updatedStationMetrics)
	stationMetricsDB.Set(extTrackStationIdByte, stationMetricsValue)

	return &types.MsgMigrateSchemaResponse{
		Status:      true,
		Description: "schema migrated successfully",
	}, nil
}
