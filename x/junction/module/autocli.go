package junction

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/ComputerKeeda/junction/api/junction/junction"
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
					RpcMethod:      "GetStation",
					Use:            "get-station [id]",
					Short:          "Query get_station",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListStations",
					Use:            "list-stations",
					Short:          "Query list_stations",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "GetStationDetailsByAddress",
					Use:            "get-station-details-by-address [address]",
					Short:          "Query get_station_details_by_address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},

				{
					RpcMethod:      "GetPod",
					Use:            "get-pod [station-id] [pod-number]",
					Short:          "Query get_pod",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stationId"}, {ProtoField: "podNumber"}},
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
					Use:            "init-station [tracks] [verification-key] [station-id] [station-info]",
					Short:          "Send a init_station tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "tracks"}, {ProtoField: "verificationKey"}, {ProtoField: "stationId"}, {ProtoField: "stationInfo"}},
				},
				{
					RpcMethod:      "SubmitPod",
					Use:            "submit-pod [station-id] [pod-number] [merkle-root-hash] [previous-merkle-root-hash] [public-witness] [timestamp]",
					Short:          "Send a submit_pod tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stationId"}, {ProtoField: "podNumber"}, {ProtoField: "merkleRootHash"}, {ProtoField: "previousMerkleRootHash"}, {ProtoField: "publicWitness"}, {ProtoField: "timestamp"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
