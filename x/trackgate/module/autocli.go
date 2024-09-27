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
					Use:            "init-station [station-id] [station-info]",
					Short:          "Send a init_station tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stationId"}, {ProtoField: "stationInfo"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
