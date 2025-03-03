package vrf

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/airchains-network/junction/api/junction/vrf"
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
					RpcMethod:      "FetchVrn",
					Use:            "fetch-vrn [key-index] [collection-id]",
					Short:          "Query fetch_vrn",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "keyIndex"}, {ProtoField: "collectionId"}},
				},

				{
					RpcMethod:      "FetchCollection",
					Use:            "fetch-collection [collection-id]",
					Short:          "Query fetch_collection",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "collectionId"}},
				},

				{
					RpcMethod:      "GetCollections",
					Use:            "get-collections",
					Short:          "Query get_collections",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "FetchCollectionDetails",
					Use:            "fetch-collection-details [collection-id]",
					Short:          "Query fetch_collection_details",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "collectionId"}},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true,
				},
				{
					RpcMethod:      "RegisterCollection",
					Use:            "register-collection [collection-name] [collection-id] [members] [offset]",
					Short:          "Send a register_collection tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "collectionName"}, {ProtoField: "collectionId"}, {ProtoField: "members"}, {ProtoField: "offset"}},
				},
				{
					RpcMethod:      "InitiateVrf",
					Use:            "initiate-vrf [key] [collection-id] [upper-bound] [origin-digest]",
					Short:          "Send a initiate_vrf tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "key"}, {ProtoField: "collectionId"}, {ProtoField: "upperBound"}, {ProtoField: "originDigest"}},
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"serializedRc": {Name: "serialized-rc-file", Usage: "Path to file containing serialized RC bytes"},
						"proof":        {Name: "proof-file", Usage: "Path to file containing proof bytes"},
						"vrfPubkey":    {Name: "vrf-pubkey-file", Usage: "Path to file containing VRF public key bytes"},
					},
				},
				{
					RpcMethod:      "ValidateVrf",
					Use:            "validate-vrf [key] [collection-id]",
					Short:          "Send a validate_vrf tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "key"}, {ProtoField: "collectionId"}},
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"serializedRc": {Name: "serialized-rc-file", Usage: "Path to file containing serialized RC bytes"},
					},
				},
				{
					RpcMethod:      "ProcessVrfDispute",
					Use:            "process-vrf-dispute [key] [collection-id] [members] [votes]",
					Short:          "Send a process_vrf_dispute tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "key"}, {ProtoField: "collectionId"}, {ProtoField: "members"}, {ProtoField: "votes"}},
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"signatures": {Name: "signature-files", Usage: "Paths to files containing signatures (repeatable)"},
						"pubKeys":    {Name: "pubkey-files", Usage: "Paths to files containing public keys (repeatable)"},
					},
				},
			},
		},
	}
}
