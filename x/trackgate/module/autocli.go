package trackgate

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/airchains-network/junction/api/junction/trackgate"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "GetExtTrackStation",
					Use:            "get-ext-track-station [id]",
					Short:          "Query get_ext_track_station",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListExtTrackStations",
					Use:            "list-ext-track-stations",
					Short:          "Query list_ext_track_stations",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "RetrieveSchemaKey",
					Use:            "retrieve-schema-key [ext-track-station-id] [schema-version]",
					Short:          "Query retrieve-schema-key",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "extTrackStationId"}, {ProtoField: "schemaVersion"}},
				},

				{
					RpcMethod:      "ListSchemas",
					Use:            "list-schemas [ext-track-station-id]",
					Short:          "Query list-schemas",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "extTrackStationId"}},
				},

				{
					RpcMethod:      "GetTrackEngagement",
					Use:            "get-track-engagement [ext-track-station-id] [pod-number]",
					Short:          "Query get-track-engagement",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "extTrackStationId"}, {ProtoField: "podNumber"}},
				},

				{
					RpcMethod:      "ListTrackEngagements",
					Use:            "list-track-engagements [ext-track-station-id]",
					Short:          "Query list-track-engagements",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "extTrackStationId"}},
				},

				{
					RpcMethod:      "GetSchemas",
					Use:            "get-schemas [ext-track-station-id] [schema-version]",
					Short:          "Query get-schemas",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "extTrackStationId"}, {ProtoField: "schemaVersion"}},
				},

				{
					RpcMethod:      "GetStationMetrics",
					Use:            "get-station-metrics [ext-track-station-id]",
					Short:          "Query get-station-metrics",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "extTrackStationId"}},
				},

				{
					RpcMethod:      "ListTrackEngagementsCustomPagination",
					Use:            "list-track-engagements-custom-pagination [ext-track-station-id] [offset] [limit] [order]",
					Short:          "Query list-track-engagements-custom-pagination",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "extTrackStationId"}, {ProtoField: "offset"}, {ProtoField: "limit"}, {ProtoField: "order"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "InitStation",
					Use:            "init-station [station-id] [station-info] [operators]",
					Short:          "Send a init_station tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stationId"}, {ProtoField: "stationInfo"}, {ProtoField: "operators"}},
				},
				{
					RpcMethod:      "SchemaCreation",
					Use:            "schema-creation [ext-track-station-id] [version] [schema]",
					Short:          "Send a schema-creation tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "extTrackStationId"}, {ProtoField: "version"}, {ProtoField: "schema"}},
				},
				{
					RpcMethod:      "SchemaEngage",
					Use:            "schema-engage [ext-track-station-id] [schema-object] [state-root] [pod-number]",
					Short:          "Send a schema-engage tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "extTrackStationId"}, {ProtoField: "schemaObject"}, {ProtoField: "acknowledgementHash"}, {ProtoField: "podNumber"}, {ProtoField: "sequencerDetails"}},
				},
				{
					RpcMethod:      "MigrateSchema",
					Use:            "migrate-schema [ext-track-station-id] [new-schema-key]",
					Short:          "Send a migrate-schema tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "extTrackStationId"}, {ProtoField: "newSchemaKey"}},
				},
				{
					RpcMethod:      "AuditSequencer",
					Use:            "audit-sequencer [sequencer-checks]",
					Short:          "Send a audit-sequencer tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "sequencerChecks"}},
				},
				{
					RpcMethod:      "LogBlobData",
					Use:            "log-blob-data [da-data]",
					Short:          "Send a log-blob-data tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "daData"}},
				},
				{
					RpcMethod:      "IntegrityCheck",
					Use:            "integrity-check [da-commitment]",
					Short:          "Send a integrity-check tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "daCommitment"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
