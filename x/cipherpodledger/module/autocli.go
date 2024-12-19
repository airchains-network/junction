package cipherpodledger

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/airchains-network/junction/api/junction/cipherpodledger"
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
					RpcMethod:      "RegisterFhvm",
					Use:            "register-fhvm [chain-id] [chain-name] [proof-type] [proving-network-verification-key] [da-provider] [da-blob-id] [relayer-g-address] [relayer-asc-address] [pic-contract-address] [acl-contract-address] [tfhe-executor-contract-address] [kms-verifier-contract-address] [gateway-contract-address] [asc-child-contract-address]",
					Short:          "Send a register_fhvm tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "chainId"}, {ProtoField: "chainName"}, {ProtoField: "proofType"}, {ProtoField: "provingNetworkVerificationKey"}, {ProtoField: "daProvider"}, {ProtoField: "daBlobId"}, {ProtoField: "relayerGaddress"}, {ProtoField: "relayerAscAddress"}, {ProtoField: "picContractAddress"}, {ProtoField: "aclContractAddress"}, {ProtoField: "tfheExecutorContractAddress"}, {ProtoField: "kmsVerifierContractAddress"}, {ProtoField: "gatewayContractAddress"}, {ProtoField: "ascChildContractAddress"}},
				},
				{
					RpcMethod:      "SubmitPod",
					Use:            "submit-pod [asc-contract-address] [pod-number] [station-id] [da-blob-id] [timestamp] [proving-network] [zk-fhe-proof] [zk-fhe-public-witness]",
					Short:          "Send a submit_pod tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "ascContractAddress"}, {ProtoField: "podNumber"}, {ProtoField: "stationId"}, {ProtoField: "daBlobId"}, {ProtoField: "timestamp"}, {ProtoField: "provingNetwork"}, {ProtoField: "zkFheproof"}, {ProtoField: "zkFhepublicWitness"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
