package junction

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/airchains-network/junction/api/junction/junction"
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

				{
					RpcMethod:      "GetLatestSubmittedPodNumber",
					Use:            "get-latest-submitted-pod-number [station-id]",
					Short:          "Query get_latest_submitted_pod_number",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stationId"}},
				},

				{
					RpcMethod:      "GetLatestVerifiedPodNumber",
					Use:            "get-latest-verified-pod-number [station-id]",
					Short:          "Query get_latest_verified_pod_number",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stationId"}},
				},

				{
					RpcMethod:      "FetchVrn",
					Use:            "fetch-vrn [pod-number] [station-id]",
					Short:          "Query fetch_vrn",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "podNumber"}, {ProtoField: "stationId"}},
				},

				{
					RpcMethod:      "GetTracks",
					Use:            "get-tracks [station-id]",
					Short:          "Query get_tracks",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stationId"}},
				},

				{
					RpcMethod:      "IsTrackMember",
					Use:            "is-track-member [track-address] [station-id]",
					Short:          "Query is_track_member",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "trackAddress"}, {ProtoField: "stationId"}},
				},

				{
					RpcMethod:      "ListStationPods",
					Use:            "list-station-pods [station-id]",
					Short:          "Query list-station-pods",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stationId"}},
				},

				{
					RpcMethod:      "ListStationPodsCustomPagination",
					Use:            "list-station-pods-custom-pagination [station-id] [offset] [limit] [order]",
					Short:          "Query list-station-pods-custom-pagination",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stationId"}, {ProtoField: "offset"}, {ProtoField: "limit"}, {ProtoField: "order"}},
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
				{
					RpcMethod:      "VerifyPod",
					Use:            "verify-pod [station-id] [pod-number] [merkle-root-hash] [previous-merkle-root-hash] [zk-proof]",
					Short:          "Send a verify_pod tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stationId"}, {ProtoField: "podNumber"}, {ProtoField: "merkleRootHash"}, {ProtoField: "previousMerkleRootHash"}, {ProtoField: "zkProof"}},
				},
				{
					RpcMethod:      "InitiateVrf",
					Use:            "initiate-vrf [pod-number] [station-id] [occupancy] [creators-vrf-key] [extra-arg]",
					Short:          "Send a initiate_vrf tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "podNumber"}, {ProtoField: "stationId"}, {ProtoField: "occupancy"}, {ProtoField: "creatorsVrfKey"}, {ProtoField: "extraArg"}},
				},
				{
					RpcMethod:      "ValidateVrf",
					Use:            "validate-vrf [station-id] [pod-number] [serialized-rc]",
					Short:          "Send a validate_vrf tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stationId"}, {ProtoField: "podNumber"}, {ProtoField: "serializedRc"}},
				},
				{
					RpcMethod:      "ProcessVrfDispute",
					Use:            "process-vrf-dispute [pod-number] [station-id] [signatures] [votes] [public-keys]",
					Short:          "Send a process_vrf_dispute tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "podNumber"}, {ProtoField: "stationId"}, {ProtoField: "signatures"}, {ProtoField: "votes"}, {ProtoField: "publicKeys"}},
				},
				{
					RpcMethod:      "AddNewTrack",
					Use:            "add-new-track [station-id] [new-track-address] [new-track-voting-power] [signatures] [votes] [public-keys]",
					Short:          "Send a add_new_track tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stationId"}, {ProtoField: "newTrackAddress"}, {ProtoField: "newTrackVotingPower"}, {ProtoField: "signatures"}, {ProtoField: "votes"}, {ProtoField: "publicKeys"}},
				},
				{
					RpcMethod:      "RemoveTrack",
					Use:            "remove-track [station-id] [track-address] [signatures] [votes] [public-keys]",
					Short:          "Send a remove_track tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stationId"}, {ProtoField: "trackAddress"}, {ProtoField: "signatures"}, {ProtoField: "votes"}, {ProtoField: "publicKeys"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
