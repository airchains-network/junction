package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"encoding/json"
	"github.com/airchains-network/junction/x/trackgate/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) AuditSequencer(goCtx context.Context, msg *types.MsgAuditSequencer) (*types.MsgAuditSequencerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	verifier := msg.Verifier
	if verifier == "" {
		return &types.MsgAuditSequencerResponse{
			Status:      false,
			Description: "Verifier address is missing or invalid.",
		}, status.Error(codes.InvalidArgument, "verifier address cannot be empty")
	}

	sequencerChecks := msg.SequencerChecks
	if sequencerChecks == nil {
		return &types.MsgAuditSequencerResponse{
			Status:      false,
			Description: "Sequencer checks are missing or invalid.",
		}, status.Error(codes.InvalidArgument, "sequencer checks cannot be empty")
	}

	for _, check := range sequencerChecks {
		podNumber := check.PodNumber
		extTrackStationId := check.ExtTrackStationId
		namespace := check.Namespace
		verificationStatus := check.VerificationStatus

		//extTrackStationsDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExtTrackStationsDataStoreKey))
		extTrackStationIdByte := []byte(extTrackStationId)
		//stationDataByte := extTrackStationsDataDB.Get(extTrackStationIdByte)
		//if stationDataByte == nil {
		//	return nil, status.Errorf(codes.NotFound, "ext track station %s not found", extTrackStationId)
		//}

		//var stationData types.ExtTrackStations
		//k.cdc.MustUnmarshal(stationDataByte, &stationData)

		newSchemaEngagementDbPath := BuildExtTrackSchemaEngagementsStoreKey(extTrackStationId)
		newSchemaEngagementStore := prefix.NewStore(storeAdapter, types.KeyPrefix(newSchemaEngagementDbPath))
		engagementKeyByte := []byte(podNumber)
		engagementDataByte := newSchemaEngagementStore.Get(engagementKeyByte)
		if engagementDataByte == nil {
			return nil, status.Errorf(codes.NotFound, "ext track engagement %s not found", podNumber)
		}

		var engagementData types.ExtTrackSchemaEngagement
		k.cdc.MustUnmarshal(engagementDataByte, &engagementData)

		sequencerDetails := engagementData.SequencerDetails
		if sequencerDetails == nil {
			return nil, status.Errorf(codes.NotFound, "ext track engagement sequencer details %s not found", podNumber)
		}

		var engagementSequencerDetails types.EngagementSequencerDetails
		err := json.Unmarshal(sequencerDetails, &engagementSequencerDetails)
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "ext track engagement sequencer details %s not found", podNumber)
		}

		verifyNamespace := engagementSequencerDetails.NameSpace

		if verifyNamespace != namespace {
			return nil, status.Errorf(codes.NotFound, "ext track engagement sequencer details %s not found", podNumber)
		}

		updateEngagementData := types.ExtTrackSchemaEngagement{
			ExtTrackStationId:   engagementData.ExtTrackStationId,
			EngageBy:            engagementData.EngageBy,
			AcknowledgementHash: engagementData.AcknowledgementHash,
			PodNumber:           engagementData.PodNumber,
			StationName:         engagementData.StationName,
			TrackName:           engagementData.TrackName,
			SchemaVersion:       engagementData.SchemaVersion,
			SchemaKey:           engagementData.SchemaKey,
			SchemaObject:        engagementData.SchemaObject,
			SequencerDetails:    engagementData.SequencerDetails,
			IsVerified:          verificationStatus,
			VerifiedBy:          verifier,
		}

		updateEngagementDataKey := []byte(podNumber)
		updateEngagementDataValue := k.cdc.MustMarshal(&updateEngagementData)
		newSchemaEngagementStore.Set(updateEngagementDataKey, updateEngagementDataValue)

		// Station Metrics Check
		stationMetricsDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TrackGateFigureStoreKey))
		stationMetricsDataBytes := stationMetricsDB.Get(extTrackStationIdByte)
		if stationMetricsDataBytes == nil {
			return nil, status.Errorf(codes.NotFound, "ext track station metrics %s not found", extTrackStationId)
		}

		var stationMetrics types.StationMetrics
		k.cdc.MustUnmarshal(stationMetricsDataBytes, &stationMetrics)

		updatedStationMetrics := types.StationMetrics{
			TotalPodCount:         stationMetrics.TotalPodCount,
			TotalSchemaCount:      stationMetrics.TotalSchemaCount,
			TotalMigrationCount:   stationMetrics.TotalMigrationCount,
			TotalVerifiedPodCount: stationMetrics.TotalVerifiedPodCount + 1,
		}
		stationMetricsValue := k.cdc.MustMarshal(&updatedStationMetrics)
		stationMetricsDB.Set(extTrackStationIdByte, stationMetricsValue)
	}

	return &types.MsgAuditSequencerResponse{
		Status:      true,
		Description: "Sequencer checks successfully audited.",
	}, nil
}
