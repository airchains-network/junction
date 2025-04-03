package rollup

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/airchains-network/junction/api/junction/rollup"
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
					RpcMethod:      "CheckMonikerAvailability",
					Use:            "check-moniker-availability [moniker]",
					Short:          "Query check-moniker-availability",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "moniker"}},
				},

				{
					RpcMethod:      "GetRollupInfo",
					Use:            "get-rollup-info [rollup-id]",
					Short:          "Query GetRollupInfo",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "rollupId"}},
				},

				{
					RpcMethod:      "GetRollups",
					Use:            "get-rollups",
					Short:          "Query GetRollups",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "GetRollupByMoniker",
					Use:            "get-rollup-by-moniker [moniker]",
					Short:          "Query GetRollupByMoniker",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "moniker"}},
				},

				{
					RpcMethod:      "GetBatchInfo",
					Use:            "get-batch-info [rollup-id] [batch-no]",
					Short:          "Query GetBatchInfo",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "rollupId"}, {ProtoField: "batchNo"}},
				},

				{
					RpcMethod:      "GetAllBatches",
					Use:            "get-all-batches [rollup-id]",
					Short:          "Query get_all_batches",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "rollupId"}},
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
					RpcMethod:      "InitRollup",
					Use:            "init-rollup [moniker] [chain-id] [denom-name] [da-type] [keys] [supply] [acl-contract-address] [kms-verifier-address] [tfhe-executor-address] [gateway-contract-address] [asc-contract-address] [relayer-g-address] [relayer-asc-address]",
					Short:          "Send a InitRollup tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "moniker"}, {ProtoField: "chainId"}, {ProtoField: "denomName"}, {ProtoField: "daType"}, {ProtoField: "keys"}, {ProtoField: "supply"}, {ProtoField: "aclContractAddress"}, {ProtoField: "kmsVerifierAddress"}, {ProtoField: "tfheExecutorAddress"}, {ProtoField: "gatewayContractAddress"}, {ProtoField: "ascContractAddress"}, {ProtoField: "relayerGAddress"}, {ProtoField: "relayerASCAddress"}},
				},
				{
					RpcMethod:      "InitProver",
					Use:            "init-prover [rollup-id] [prover-verification-key] [prover-type] [prover-endpoint]",
					Short:          "Send a InitProver tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "rollupId"}, {ProtoField: "proverVerificationKey"}, {ProtoField: "proverType"}, {ProtoField: "proverEndpoint"}},
				},
				{
					RpcMethod:      "SubmitBatchMetadata",
					Use:            "submit-batch-metadata [batch-no] [rollup-id] [da-name] [da-commitment] [da-hash] [da-pointer] [da-namespace]",
					Short:          "Send a SubmitBatchMetadata tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "batchNo"}, {ProtoField: "rollupId"}, {ProtoField: "daName"}, {ProtoField: "daCommitment"}, {ProtoField: "daHash"}, {ProtoField: "daPointer"}, {ProtoField: "daNamespace"}},
				},
				{
					RpcMethod:      "SubmitBatch",
					Use:            "submit-batch [rollup-id] [batch-no] [merkle-root-hash] [previous-merkle-root-hash] [zk-proof] [public-witness]",
					Short:          "Send a SubmitBatch tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "rollupId"}, {ProtoField: "batchNo"}, {ProtoField: "merkleRootHash"}, {ProtoField: "previousMerkleRootHash"}, {ProtoField: "zkProof"}, {ProtoField: "publicWitness"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
