package types

const (
	EventTypeRollupProverInitialized = "rollup-prover-initialized"
	EventTypeTokenLocked             = "token-locked"
	EventTypeRollupInitialized       = "rollup-initialized"
	EventTypeBatchMetadataSubmitted  = "batch-metadata-submitted"
)

const (
	AttributeKeyRollupId               = "rollup-id"
	AttributeKeyRollupMoniker          = "rollup-moniker"
	AttributeKeyMoniker                = "moniker" // TODO: should merge with 'rollup-moniker' ?
	AttributeKeyProverType             = "prover-type"
	AttributeKeyProverEndpoint         = "prover-endpoint"
	AttributeKeyCreator                = "creator"
	AttributeKeyDenom                  = "denom"
	AttributeKeyDenomName              = "denom-name"
	AttributeKeyChainId                = "chain-id"
	AttributeKeyDaType                 = "da-type"
	AttributeKeyDaName                 = "da-name" // TODO: should merge with 'da-type' ?
	AttributeKeyKeys                   = "keys"
	AttributeKeySupply                 = "supply"
	AttributeKeyTotalSupply            = "total-supply"
	AttributeKeyAclContractAddress     = "acl-contract-address"
	AttributeKeyKmsVerifierAddress     = "kms-verifier-address"
	AttributeKeyTfheExecutorAddress    = "tfhe-executor-address"
	AttributeKeyGatewayContractAddress = "gateway-contract-address"
	AttributeKeyAscContractAddress     = "asc-contract-address"
	AttributeKeyRelayerGAddress        = "relayer-g-address"
	AttributeKeySubmitter              = "submitter"
	AttributeKeyBatchNo                = "batch-no"
	AttributeKeyTimestamp              = "timestamp"
	AttributeKeyDaCommitment           = "da-commitment"
	AttributeKeyDaHash                 = "da-hash"
	AttributeKeyDaPointer              = "da-pointer"
	AttributeKeyDaNamespace            = "da-namespace"
	AttributeKeyPreviousMerkleRootHash = "previous-merkle-root-hash"
	AttributeKeyIsFinalized            = "is-finalized"
)
