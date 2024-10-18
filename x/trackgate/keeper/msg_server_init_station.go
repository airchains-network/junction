package keeper

import (
	"context"
	"encoding/json"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) InitStation(goCtx context.Context, msg *types.MsgInitStation) (*types.MsgInitStationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	submitter := msg.Submitter
	stationId := msg.StationId
	stationInfo := msg.StationInfo
	operators := msg.Operators

	var stationInfoDetails types.StationInfoDetails
	stationInfoUnmarshalErr := json.Unmarshal(stationInfo, &stationInfoDetails)
	if stationInfoUnmarshalErr != nil {
		return &types.MsgInitStationResponse{
			Status: false,
		}, stationInfoUnmarshalErr
	}

	byteSequencerDetails, err := json.Marshal(stationInfoDetails.SequencerDetails)
	if err != nil {
		return &types.MsgInitStationResponse{
			Status: false,
		}, status.Error(codes.InvalidArgument, "incorrect extra data argument")
	}
	byteDaDetails, err := json.Marshal(stationInfoDetails.DADetails)
	if err != nil {
		return &types.MsgInitStationResponse{
			Status: false,
		}, status.Error(codes.InvalidArgument, "incorrect extra data argument")
	}
	byteProverDetails, err := json.Marshal(stationInfoDetails.ProverDetails)
	if err != nil {
		return &types.MsgInitStationResponse{
			Status: false,
		}, status.Error(codes.InvalidArgument, "incorrect extra data argument")
	}
	var newExtTrackStation = types.ExtTrackStations{
		Creator:                   submitter,
		Operators:                 operators,
		LatestPod:                 0,
		LatestAcknowledgementHash: "0",
		Name:                      stationInfoDetails.StationName,
		Id:                        stationId,
		StationType:               stationInfoDetails.Type,
		FheEnabled:                stationInfoDetails.FheEnabled,
		SequencerDetails:          byteSequencerDetails,
		DaDetails:                 byteDaDetails,
		ProverDetails:             byteProverDetails,
		StationSchemaKey:          "none",
	}

	newStationMetrics := types.StationMetrics{
		TotalPodCount:         0,
		TotalSchemaCount:      0,
		TotalMigrationCount:   0,
		TotalVerifiedPodCount: 0,
	}

	stationMetricsDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TrackGateFigureStoreKey))
	stationMetricsKey := []byte(stationId)
	stationMetricsValue := k.cdc.MustMarshal(&newStationMetrics)
	stationMetricsDB.Set(stationMetricsKey, stationMetricsValue)

	extTrackStationsDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExtTrackStationsDataStoreKey))
	stationIdBytes := []byte(stationId)
	uniqueStationIDCheck := extTrackStationsDataDB.Get(stationIdBytes)
	if uniqueStationIDCheck != nil {
		return &types.MsgInitStationResponse{
			Status: false,
		}, status.Error(codes.InvalidArgument, "station id already exists")
	}

	byteStation := k.cdc.MustMarshal(&newExtTrackStation)
	byteStationID := []byte(stationId)

	extTrackStationsDataDB.Set(byteStationID, byteStation)

	return &types.MsgInitStationResponse{
		Status: true,
	}, nil
}
