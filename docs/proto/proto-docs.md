<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [junction/wasm/v1/types.proto](#junction/wasm/v1/types.proto)
    - [AbsoluteTxPosition](#junction.wasm.v1.AbsoluteTxPosition)
    - [AccessConfig](#junction.wasm.v1.AccessConfig)
    - [AccessTypeParam](#junction.wasm.v1.AccessTypeParam)
    - [CodeInfo](#junction.wasm.v1.CodeInfo)
    - [ContractCodeHistoryEntry](#junction.wasm.v1.ContractCodeHistoryEntry)
    - [ContractInfo](#junction.wasm.v1.ContractInfo)
    - [Model](#junction.wasm.v1.Model)
    - [Params](#junction.wasm.v1.Params)
  
    - [AccessType](#junction.wasm.v1.AccessType)
    - [ContractCodeHistoryOperationType](#junction.wasm.v1.ContractCodeHistoryOperationType)
  
- [junction/wasm/v1/authz.proto](#junction/wasm/v1/authz.proto)
    - [AcceptedMessageKeysFilter](#junction.wasm.v1.AcceptedMessageKeysFilter)
    - [AcceptedMessagesFilter](#junction.wasm.v1.AcceptedMessagesFilter)
    - [AllowAllMessagesFilter](#junction.wasm.v1.AllowAllMessagesFilter)
    - [CodeGrant](#junction.wasm.v1.CodeGrant)
    - [CombinedLimit](#junction.wasm.v1.CombinedLimit)
    - [ContractExecutionAuthorization](#junction.wasm.v1.ContractExecutionAuthorization)
    - [ContractGrant](#junction.wasm.v1.ContractGrant)
    - [ContractMigrationAuthorization](#junction.wasm.v1.ContractMigrationAuthorization)
    - [MaxCallsLimit](#junction.wasm.v1.MaxCallsLimit)
    - [MaxFundsLimit](#junction.wasm.v1.MaxFundsLimit)
    - [StoreCodeAuthorization](#junction.wasm.v1.StoreCodeAuthorization)
  
- [junction/wasm/v1/genesis.proto](#junction/wasm/v1/genesis.proto)
    - [Code](#junction.wasm.v1.Code)
    - [Contract](#junction.wasm.v1.Contract)
    - [GenesisState](#junction.wasm.v1.GenesisState)
    - [Sequence](#junction.wasm.v1.Sequence)
  
- [junction/wasm/v1/ibc.proto](#junction/wasm/v1/ibc.proto)
    - [MsgIBCCloseChannel](#junction.wasm.v1.MsgIBCCloseChannel)
    - [MsgIBCSend](#junction.wasm.v1.MsgIBCSend)
    - [MsgIBCSendResponse](#junction.wasm.v1.MsgIBCSendResponse)
    - [MsgIBCWriteAcknowledgementResponse](#junction.wasm.v1.MsgIBCWriteAcknowledgementResponse)
  
- [junction/wasm/v1/proposal_legacy.proto](#junction/wasm/v1/proposal_legacy.proto)
    - [AccessConfigUpdate](#junction.wasm.v1.AccessConfigUpdate)
    - [ClearAdminProposal](#junction.wasm.v1.ClearAdminProposal)
    - [ExecuteContractProposal](#junction.wasm.v1.ExecuteContractProposal)
    - [InstantiateContract2Proposal](#junction.wasm.v1.InstantiateContract2Proposal)
    - [InstantiateContractProposal](#junction.wasm.v1.InstantiateContractProposal)
    - [MigrateContractProposal](#junction.wasm.v1.MigrateContractProposal)
    - [PinCodesProposal](#junction.wasm.v1.PinCodesProposal)
    - [StoreAndInstantiateContractProposal](#junction.wasm.v1.StoreAndInstantiateContractProposal)
    - [StoreCodeProposal](#junction.wasm.v1.StoreCodeProposal)
    - [SudoContractProposal](#junction.wasm.v1.SudoContractProposal)
    - [UnpinCodesProposal](#junction.wasm.v1.UnpinCodesProposal)
    - [UpdateAdminProposal](#junction.wasm.v1.UpdateAdminProposal)
    - [UpdateInstantiateConfigProposal](#junction.wasm.v1.UpdateInstantiateConfigProposal)
  
- [junction/wasm/v1/query.proto](#junction/wasm/v1/query.proto)
    - [CodeInfoResponse](#junction.wasm.v1.CodeInfoResponse)
    - [QueryAllContractStateRequest](#junction.wasm.v1.QueryAllContractStateRequest)
    - [QueryAllContractStateResponse](#junction.wasm.v1.QueryAllContractStateResponse)
    - [QueryBuildAddressRequest](#junction.wasm.v1.QueryBuildAddressRequest)
    - [QueryBuildAddressResponse](#junction.wasm.v1.QueryBuildAddressResponse)
    - [QueryCodeRequest](#junction.wasm.v1.QueryCodeRequest)
    - [QueryCodeResponse](#junction.wasm.v1.QueryCodeResponse)
    - [QueryCodesRequest](#junction.wasm.v1.QueryCodesRequest)
    - [QueryCodesResponse](#junction.wasm.v1.QueryCodesResponse)
    - [QueryContractHistoryRequest](#junction.wasm.v1.QueryContractHistoryRequest)
    - [QueryContractHistoryResponse](#junction.wasm.v1.QueryContractHistoryResponse)
    - [QueryContractInfoRequest](#junction.wasm.v1.QueryContractInfoRequest)
    - [QueryContractInfoResponse](#junction.wasm.v1.QueryContractInfoResponse)
    - [QueryContractsByCodeRequest](#junction.wasm.v1.QueryContractsByCodeRequest)
    - [QueryContractsByCodeResponse](#junction.wasm.v1.QueryContractsByCodeResponse)
    - [QueryContractsByCreatorRequest](#junction.wasm.v1.QueryContractsByCreatorRequest)
    - [QueryContractsByCreatorResponse](#junction.wasm.v1.QueryContractsByCreatorResponse)
    - [QueryParamsRequest](#junction.wasm.v1.QueryParamsRequest)
    - [QueryParamsResponse](#junction.wasm.v1.QueryParamsResponse)
    - [QueryPinnedCodesRequest](#junction.wasm.v1.QueryPinnedCodesRequest)
    - [QueryPinnedCodesResponse](#junction.wasm.v1.QueryPinnedCodesResponse)
    - [QueryRawContractStateRequest](#junction.wasm.v1.QueryRawContractStateRequest)
    - [QueryRawContractStateResponse](#junction.wasm.v1.QueryRawContractStateResponse)
    - [QuerySmartContractStateRequest](#junction.wasm.v1.QuerySmartContractStateRequest)
    - [QuerySmartContractStateResponse](#junction.wasm.v1.QuerySmartContractStateResponse)
  
    - [Query](#junction.wasm.v1.Query)
  
- [junction/wasm/v1/tx.proto](#junction/wasm/v1/tx.proto)
    - [MsgAddCodeUploadParamsAddresses](#junction.wasm.v1.MsgAddCodeUploadParamsAddresses)
    - [MsgAddCodeUploadParamsAddressesResponse](#junction.wasm.v1.MsgAddCodeUploadParamsAddressesResponse)
    - [MsgClearAdmin](#junction.wasm.v1.MsgClearAdmin)
    - [MsgClearAdminResponse](#junction.wasm.v1.MsgClearAdminResponse)
    - [MsgExecuteContract](#junction.wasm.v1.MsgExecuteContract)
    - [MsgExecuteContractResponse](#junction.wasm.v1.MsgExecuteContractResponse)
    - [MsgInstantiateContract](#junction.wasm.v1.MsgInstantiateContract)
    - [MsgInstantiateContract2](#junction.wasm.v1.MsgInstantiateContract2)
    - [MsgInstantiateContract2Response](#junction.wasm.v1.MsgInstantiateContract2Response)
    - [MsgInstantiateContractResponse](#junction.wasm.v1.MsgInstantiateContractResponse)
    - [MsgMigrateContract](#junction.wasm.v1.MsgMigrateContract)
    - [MsgMigrateContractResponse](#junction.wasm.v1.MsgMigrateContractResponse)
    - [MsgPinCodes](#junction.wasm.v1.MsgPinCodes)
    - [MsgPinCodesResponse](#junction.wasm.v1.MsgPinCodesResponse)
    - [MsgRemoveCodeUploadParamsAddresses](#junction.wasm.v1.MsgRemoveCodeUploadParamsAddresses)
    - [MsgRemoveCodeUploadParamsAddressesResponse](#junction.wasm.v1.MsgRemoveCodeUploadParamsAddressesResponse)
    - [MsgStoreAndInstantiateContract](#junction.wasm.v1.MsgStoreAndInstantiateContract)
    - [MsgStoreAndInstantiateContractResponse](#junction.wasm.v1.MsgStoreAndInstantiateContractResponse)
    - [MsgStoreAndMigrateContract](#junction.wasm.v1.MsgStoreAndMigrateContract)
    - [MsgStoreAndMigrateContractResponse](#junction.wasm.v1.MsgStoreAndMigrateContractResponse)
    - [MsgStoreCode](#junction.wasm.v1.MsgStoreCode)
    - [MsgStoreCodeResponse](#junction.wasm.v1.MsgStoreCodeResponse)
    - [MsgSudoContract](#junction.wasm.v1.MsgSudoContract)
    - [MsgSudoContractResponse](#junction.wasm.v1.MsgSudoContractResponse)
    - [MsgUnpinCodes](#junction.wasm.v1.MsgUnpinCodes)
    - [MsgUnpinCodesResponse](#junction.wasm.v1.MsgUnpinCodesResponse)
    - [MsgUpdateAdmin](#junction.wasm.v1.MsgUpdateAdmin)
    - [MsgUpdateAdminResponse](#junction.wasm.v1.MsgUpdateAdminResponse)
    - [MsgUpdateContractLabel](#junction.wasm.v1.MsgUpdateContractLabel)
    - [MsgUpdateContractLabelResponse](#junction.wasm.v1.MsgUpdateContractLabelResponse)
    - [MsgUpdateInstantiateConfig](#junction.wasm.v1.MsgUpdateInstantiateConfig)
    - [MsgUpdateInstantiateConfigResponse](#junction.wasm.v1.MsgUpdateInstantiateConfigResponse)
    - [MsgUpdateParams](#junction.wasm.v1.MsgUpdateParams)
    - [MsgUpdateParamsResponse](#junction.wasm.v1.MsgUpdateParamsResponse)
  
    - [Msg](#junction.wasm.v1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="junction/wasm/v1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## junction/wasm/v1/types.proto



<a name="junction.wasm.v1.AbsoluteTxPosition"></a>

### AbsoluteTxPosition
AbsoluteTxPosition is a unique transaction position that allows for global
ordering of transactions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `block_height` | [uint64](#uint64) |  | BlockHeight is the block the contract was created at |
| `tx_index` | [uint64](#uint64) |  | TxIndex is a monotonic counter within the block (actual transaction index, or gas consumed) |






<a name="junction.wasm.v1.AccessConfig"></a>

### AccessConfig
AccessConfig access control type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `permission` | [AccessType](#junction.wasm.v1.AccessType) |  |  |
| `addresses` | [string](#string) | repeated |  |






<a name="junction.wasm.v1.AccessTypeParam"></a>

### AccessTypeParam
AccessTypeParam


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `value` | [AccessType](#junction.wasm.v1.AccessType) |  |  |






<a name="junction.wasm.v1.CodeInfo"></a>

### CodeInfo
CodeInfo is data for the uploaded contract WASM code


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_hash` | [bytes](#bytes) |  | CodeHash is the unique identifier created by wasmvm |
| `creator` | [string](#string) |  | Creator address who initially stored the code |
| `instantiate_config` | [AccessConfig](#junction.wasm.v1.AccessConfig) |  | InstantiateConfig access control to apply on contract creation, optional |






<a name="junction.wasm.v1.ContractCodeHistoryEntry"></a>

### ContractCodeHistoryEntry
ContractCodeHistoryEntry metadata to a contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `operation` | [ContractCodeHistoryOperationType](#junction.wasm.v1.ContractCodeHistoryOperationType) |  |  |
| `code_id` | [uint64](#uint64) |  | CodeID is the reference to the stored WASM code |
| `updated` | [AbsoluteTxPosition](#junction.wasm.v1.AbsoluteTxPosition) |  | Updated Tx position when the operation was executed. |
| `msg` | [bytes](#bytes) |  |  |






<a name="junction.wasm.v1.ContractInfo"></a>

### ContractInfo
ContractInfo stores a WASM contract instance


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_id` | [uint64](#uint64) |  | CodeID is the reference to the stored Wasm code |
| `creator` | [string](#string) |  | Creator address who initially instantiated the contract |
| `admin` | [string](#string) |  | Admin is an optional address that can execute migrations |
| `label` | [string](#string) |  | Label is optional metadata to be stored with a contract instance. |
| `created` | [AbsoluteTxPosition](#junction.wasm.v1.AbsoluteTxPosition) |  | Created Tx position when the contract was instantiated. |
| `ibc_port_id` | [string](#string) |  |  |
| `extension` | [google.protobuf.Any](#google.protobuf.Any) |  | Extension is an extension point to store custom metadata within the persistence model. |






<a name="junction.wasm.v1.Model"></a>

### Model
Model is a struct that holds a KV pair


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  | hex-encode key to read it better (this is often ascii) |
| `value` | [bytes](#bytes) |  | base64-encode raw value |






<a name="junction.wasm.v1.Params"></a>

### Params
Params defines the set of wasm parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_upload_access` | [AccessConfig](#junction.wasm.v1.AccessConfig) |  |  |
| `instantiate_default_permission` | [AccessType](#junction.wasm.v1.AccessType) |  |  |





 <!-- end messages -->


<a name="junction.wasm.v1.AccessType"></a>

### AccessType
AccessType permission types

| Name | Number | Description |
| ---- | ------ | ----------- |
| ACCESS_TYPE_UNSPECIFIED | 0 | AccessTypeUnspecified placeholder for empty value |
| ACCESS_TYPE_NOBODY | 1 | AccessTypeNobody forbidden |
| ACCESS_TYPE_EVERYBODY | 3 | AccessTypeEverybody unrestricted |
| ACCESS_TYPE_ANY_OF_ADDRESSES | 4 | AccessTypeAnyOfAddresses allow any of the addresses |



<a name="junction.wasm.v1.ContractCodeHistoryOperationType"></a>

### ContractCodeHistoryOperationType
ContractCodeHistoryOperationType actions that caused a code change

| Name | Number | Description |
| ---- | ------ | ----------- |
| CONTRACT_CODE_HISTORY_OPERATION_TYPE_UNSPECIFIED | 0 | ContractCodeHistoryOperationTypeUnspecified placeholder for empty value |
| CONTRACT_CODE_HISTORY_OPERATION_TYPE_INIT | 1 | ContractCodeHistoryOperationTypeInit on chain contract instantiation |
| CONTRACT_CODE_HISTORY_OPERATION_TYPE_MIGRATE | 2 | ContractCodeHistoryOperationTypeMigrate code migration |
| CONTRACT_CODE_HISTORY_OPERATION_TYPE_GENESIS | 3 | ContractCodeHistoryOperationTypeGenesis based on genesis data |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="junction/wasm/v1/authz.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## junction/wasm/v1/authz.proto



<a name="junction.wasm.v1.AcceptedMessageKeysFilter"></a>

### AcceptedMessageKeysFilter
AcceptedMessageKeysFilter accept only the specific contract message keys in
the json object to be executed.
Since: wasmd 0.30


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keys` | [string](#string) | repeated | Messages is the list of unique keys |






<a name="junction.wasm.v1.AcceptedMessagesFilter"></a>

### AcceptedMessagesFilter
AcceptedMessagesFilter accept only the specific raw contract messages to be
executed.
Since: wasmd 0.30


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `messages` | [bytes](#bytes) | repeated | Messages is the list of raw contract messages |






<a name="junction.wasm.v1.AllowAllMessagesFilter"></a>

### AllowAllMessagesFilter
AllowAllMessagesFilter is a wildcard to allow any type of contract payload
message.
Since: wasmd 0.30






<a name="junction.wasm.v1.CodeGrant"></a>

### CodeGrant
CodeGrant a granted permission for a single code


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_hash` | [bytes](#bytes) |  | CodeHash is the unique identifier created by wasmvm Wildcard "*" is used to specify any kind of grant. |
| `instantiate_permission` | [AccessConfig](#junction.wasm.v1.AccessConfig) |  | InstantiatePermission is the superset access control to apply on contract creation. Optional |






<a name="junction.wasm.v1.CombinedLimit"></a>

### CombinedLimit
CombinedLimit defines the maximal amounts that can be sent to a contract and
the maximal number of calls executable. Both need to remain >0 to be valid.
Since: wasmd 0.30


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `calls_remaining` | [uint64](#uint64) |  | Remaining number that is decremented on each execution |
| `amounts` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Amounts is the maximal amount of tokens transferable to the contract. |






<a name="junction.wasm.v1.ContractExecutionAuthorization"></a>

### ContractExecutionAuthorization
ContractExecutionAuthorization defines authorization for wasm execute.
Since: wasmd 0.30


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `grants` | [ContractGrant](#junction.wasm.v1.ContractGrant) | repeated | Grants for contract executions |






<a name="junction.wasm.v1.ContractGrant"></a>

### ContractGrant
ContractGrant a granted permission for a single contract
Since: wasmd 0.30


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract` | [string](#string) |  | Contract is the bech32 address of the smart contract |
| `limit` | [google.protobuf.Any](#google.protobuf.Any) |  | Limit defines execution limits that are enforced and updated when the grant is applied. When the limit lapsed the grant is removed. |
| `filter` | [google.protobuf.Any](#google.protobuf.Any) |  | Filter define more fine-grained control on the message payload passed to the contract in the operation. When no filter applies on execution, the operation is prohibited. |






<a name="junction.wasm.v1.ContractMigrationAuthorization"></a>

### ContractMigrationAuthorization
ContractMigrationAuthorization defines authorization for wasm contract
migration. Since: wasmd 0.30


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `grants` | [ContractGrant](#junction.wasm.v1.ContractGrant) | repeated | Grants for contract migrations |






<a name="junction.wasm.v1.MaxCallsLimit"></a>

### MaxCallsLimit
MaxCallsLimit limited number of calls to the contract. No funds transferable.
Since: wasmd 0.30


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `remaining` | [uint64](#uint64) |  | Remaining number that is decremented on each execution |






<a name="junction.wasm.v1.MaxFundsLimit"></a>

### MaxFundsLimit
MaxFundsLimit defines the maximal amounts that can be sent to the contract.
Since: wasmd 0.30


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amounts` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Amounts is the maximal amount of tokens transferable to the contract. |






<a name="junction.wasm.v1.StoreCodeAuthorization"></a>

### StoreCodeAuthorization
StoreCodeAuthorization defines authorization for wasm code upload.
Since: wasmd 0.42


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `grants` | [CodeGrant](#junction.wasm.v1.CodeGrant) | repeated | Grants for code upload |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="junction/wasm/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## junction/wasm/v1/genesis.proto



<a name="junction.wasm.v1.Code"></a>

### Code
Code struct encompasses CodeInfo and CodeBytes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_id` | [uint64](#uint64) |  |  |
| `code_info` | [CodeInfo](#junction.wasm.v1.CodeInfo) |  |  |
| `code_bytes` | [bytes](#bytes) |  |  |
| `pinned` | [bool](#bool) |  | Pinned to wasmvm cache |






<a name="junction.wasm.v1.Contract"></a>

### Contract
Contract struct encompasses ContractAddress, ContractInfo, and ContractState


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |
| `contract_info` | [ContractInfo](#junction.wasm.v1.ContractInfo) |  |  |
| `contract_state` | [Model](#junction.wasm.v1.Model) | repeated |  |
| `contract_code_history` | [ContractCodeHistoryEntry](#junction.wasm.v1.ContractCodeHistoryEntry) | repeated |  |






<a name="junction.wasm.v1.GenesisState"></a>

### GenesisState
GenesisState - genesis state of x/wasm


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#junction.wasm.v1.Params) |  |  |
| `codes` | [Code](#junction.wasm.v1.Code) | repeated |  |
| `contracts` | [Contract](#junction.wasm.v1.Contract) | repeated |  |
| `sequences` | [Sequence](#junction.wasm.v1.Sequence) | repeated |  |






<a name="junction.wasm.v1.Sequence"></a>

### Sequence
Sequence key and value of an id generation counter


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id_key` | [bytes](#bytes) |  |  |
| `value` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="junction/wasm/v1/ibc.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## junction/wasm/v1/ibc.proto



<a name="junction.wasm.v1.MsgIBCCloseChannel"></a>

### MsgIBCCloseChannel
MsgIBCCloseChannel port and channel need to be owned by the contract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `channel` | [string](#string) |  |  |






<a name="junction.wasm.v1.MsgIBCSend"></a>

### MsgIBCSend
MsgIBCSend


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `channel` | [string](#string) |  | the channel by which the packet will be sent |
| `timeout_height` | [uint64](#uint64) |  | Timeout height relative to the current block height. The timeout is disabled when set to 0. |
| `timeout_timestamp` | [uint64](#uint64) |  | Timeout timestamp (in nanoseconds) relative to the current block timestamp. The timeout is disabled when set to 0. |
| `data` | [bytes](#bytes) |  | Data is the payload to transfer. We must not make assumption what format or content is in here. |






<a name="junction.wasm.v1.MsgIBCSendResponse"></a>

### MsgIBCSendResponse
MsgIBCSendResponse


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sequence` | [uint64](#uint64) |  | Sequence number of the IBC packet sent |






<a name="junction.wasm.v1.MsgIBCWriteAcknowledgementResponse"></a>

### MsgIBCWriteAcknowledgementResponse
MsgIBCWriteAcknowledgementResponse





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="junction/wasm/v1/proposal_legacy.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## junction/wasm/v1/proposal_legacy.proto



<a name="junction.wasm.v1.AccessConfigUpdate"></a>

### AccessConfigUpdate
AccessConfigUpdate contains the code id and the access config to be
applied.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_id` | [uint64](#uint64) |  | CodeID is the reference to the stored WASM code to be updated |
| `instantiate_permission` | [AccessConfig](#junction.wasm.v1.AccessConfig) |  | InstantiatePermission to apply to the set of code ids |






<a name="junction.wasm.v1.ClearAdminProposal"></a>

### ClearAdminProposal
Deprecated: Do not use. Since wasmd v0.40, there is no longer a need for
an explicit ClearAdminProposal. To clear the admin of a contract,
a simple MsgClearAdmin can be invoked from the x/gov module via
a v1 governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `contract` | [string](#string) |  | Contract is the address of the smart contract |






<a name="junction.wasm.v1.ExecuteContractProposal"></a>

### ExecuteContractProposal
Deprecated: Do not use. Since wasmd v0.40, there is no longer a need for
an explicit ExecuteContractProposal. To call execute on a contract,
a simple MsgExecuteContract can be invoked from the x/gov module via
a v1 governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `run_as` | [string](#string) |  | RunAs is the address that is passed to the contract's environment as sender |
| `contract` | [string](#string) |  | Contract is the address of the smart contract |
| `msg` | [bytes](#bytes) |  | Msg json encoded message to be passed to the contract as execute |
| `funds` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Funds coins that are transferred to the contract on instantiation |






<a name="junction.wasm.v1.InstantiateContract2Proposal"></a>

### InstantiateContract2Proposal
Deprecated: Do not use. Since wasmd v0.40, there is no longer a need for
an explicit InstantiateContract2Proposal. To instantiate contract 2,
a simple MsgInstantiateContract2 can be invoked from the x/gov module via
a v1 governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `run_as` | [string](#string) |  | RunAs is the address that is passed to the contract's enviroment as sender |
| `admin` | [string](#string) |  | Admin is an optional address that can execute migrations |
| `code_id` | [uint64](#uint64) |  | CodeID is the reference to the stored WASM code |
| `label` | [string](#string) |  | Label is optional metadata to be stored with a constract instance. |
| `msg` | [bytes](#bytes) |  | Msg json encode message to be passed to the contract on instantiation |
| `funds` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Funds coins that are transferred to the contract on instantiation |
| `salt` | [bytes](#bytes) |  | Salt is an arbitrary value provided by the sender. Size can be 1 to 64. |
| `fix_msg` | [bool](#bool) |  | FixMsg include the msg value into the hash for the predictable address. Default is false |






<a name="junction.wasm.v1.InstantiateContractProposal"></a>

### InstantiateContractProposal
Deprecated: Do not use. Since wasmd v0.40, there is no longer a need for
an explicit InstantiateContractProposal. To instantiate a contract,
a simple MsgInstantiateContract can be invoked from the x/gov module via
a v1 governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `run_as` | [string](#string) |  | RunAs is the address that is passed to the contract's environment as sender |
| `admin` | [string](#string) |  | Admin is an optional address that can execute migrations |
| `code_id` | [uint64](#uint64) |  | CodeID is the reference to the stored WASM code |
| `label` | [string](#string) |  | Label is optional metadata to be stored with a constract instance. |
| `msg` | [bytes](#bytes) |  | Msg json encoded message to be passed to the contract on instantiation |
| `funds` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Funds coins that are transferred to the contract on instantiation |






<a name="junction.wasm.v1.MigrateContractProposal"></a>

### MigrateContractProposal
Deprecated: Do not use. Since wasmd v0.40, there is no longer a need for
an explicit MigrateContractProposal. To migrate a contract,
a simple MsgMigrateContract can be invoked from the x/gov module via
a v1 governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text

Note: skipping 3 as this was previously used for unneeded run_as |
| `contract` | [string](#string) |  | Contract is the address of the smart contract |
| `code_id` | [uint64](#uint64) |  | CodeID references the new WASM code |
| `msg` | [bytes](#bytes) |  | Msg json encoded message to be passed to the contract on migration |






<a name="junction.wasm.v1.PinCodesProposal"></a>

### PinCodesProposal
Deprecated: Do not use. Since wasmd v0.40, there is no longer a need for
an explicit PinCodesProposal. To pin a set of code ids in the wasmvm
cache, a simple MsgPinCodes can be invoked from the x/gov module via
a v1 governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `code_ids` | [uint64](#uint64) | repeated | CodeIDs references the new WASM codes |






<a name="junction.wasm.v1.StoreAndInstantiateContractProposal"></a>

### StoreAndInstantiateContractProposal
Deprecated: Do not use. Since wasmd v0.40, there is no longer a need for
an explicit StoreAndInstantiateContractProposal. To store and instantiate
the contract, a simple MsgStoreAndInstantiateContract can be invoked from
the x/gov module via a v1 governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `run_as` | [string](#string) |  | RunAs is the address that is passed to the contract's environment as sender |
| `wasm_byte_code` | [bytes](#bytes) |  | WASMByteCode can be raw or gzip compressed |
| `instantiate_permission` | [AccessConfig](#junction.wasm.v1.AccessConfig) |  | InstantiatePermission to apply on contract creation, optional |
| `unpin_code` | [bool](#bool) |  | UnpinCode code on upload, optional |
| `admin` | [string](#string) |  | Admin is an optional address that can execute migrations |
| `label` | [string](#string) |  | Label is optional metadata to be stored with a constract instance. |
| `msg` | [bytes](#bytes) |  | Msg json encoded message to be passed to the contract on instantiation |
| `funds` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Funds coins that are transferred to the contract on instantiation |
| `source` | [string](#string) |  | Source is the URL where the code is hosted |
| `builder` | [string](#string) |  | Builder is the docker image used to build the code deterministically, used for smart contract verification |
| `code_hash` | [bytes](#bytes) |  | CodeHash is the SHA256 sum of the code outputted by builder, used for smart contract verification |






<a name="junction.wasm.v1.StoreCodeProposal"></a>

### StoreCodeProposal
Deprecated: Do not use. Since wasmd v0.40, there is no longer a need for
an explicit StoreCodeProposal. To submit WASM code to the system,
a simple MsgStoreCode can be invoked from the x/gov module via
a v1 governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `run_as` | [string](#string) |  | RunAs is the address that is passed to the contract's environment as sender |
| `wasm_byte_code` | [bytes](#bytes) |  | WASMByteCode can be raw or gzip compressed |
| `instantiate_permission` | [AccessConfig](#junction.wasm.v1.AccessConfig) |  | InstantiatePermission to apply on contract creation, optional |
| `unpin_code` | [bool](#bool) |  | UnpinCode code on upload, optional |
| `source` | [string](#string) |  | Source is the URL where the code is hosted |
| `builder` | [string](#string) |  | Builder is the docker image used to build the code deterministically, used for smart contract verification |
| `code_hash` | [bytes](#bytes) |  | CodeHash is the SHA256 sum of the code outputted by builder, used for smart contract verification |






<a name="junction.wasm.v1.SudoContractProposal"></a>

### SudoContractProposal
Deprecated: Do not use. Since wasmd v0.40, there is no longer a need for
an explicit SudoContractProposal. To call sudo on a contract,
a simple MsgSudoContract can be invoked from the x/gov module via
a v1 governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `contract` | [string](#string) |  | Contract is the address of the smart contract |
| `msg` | [bytes](#bytes) |  | Msg json encoded message to be passed to the contract as sudo |






<a name="junction.wasm.v1.UnpinCodesProposal"></a>

### UnpinCodesProposal
Deprecated: Do not use. Since wasmd v0.40, there is no longer a need for
an explicit UnpinCodesProposal. To unpin a set of code ids in the wasmvm
cache, a simple MsgUnpinCodes can be invoked from the x/gov module via
a v1 governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `code_ids` | [uint64](#uint64) | repeated | CodeIDs references the WASM codes |






<a name="junction.wasm.v1.UpdateAdminProposal"></a>

### UpdateAdminProposal
Deprecated: Do not use. Since wasmd v0.40, there is no longer a need for
an explicit UpdateAdminProposal. To set an admin for a contract,
a simple MsgUpdateAdmin can be invoked from the x/gov module via
a v1 governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `new_admin` | [string](#string) |  | NewAdmin address to be set |
| `contract` | [string](#string) |  | Contract is the address of the smart contract |






<a name="junction.wasm.v1.UpdateInstantiateConfigProposal"></a>

### UpdateInstantiateConfigProposal
Deprecated: Do not use. Since wasmd v0.40, there is no longer a need for
an explicit UpdateInstantiateConfigProposal. To update instantiate config
to a set of code ids, a simple MsgUpdateInstantiateConfig can be invoked from
the x/gov module via a v1 governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `access_config_updates` | [AccessConfigUpdate](#junction.wasm.v1.AccessConfigUpdate) | repeated | AccessConfigUpdate contains the list of code ids and the access config to be applied. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="junction/wasm/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## junction/wasm/v1/query.proto



<a name="junction.wasm.v1.CodeInfoResponse"></a>

### CodeInfoResponse
CodeInfoResponse contains code meta data from CodeInfo


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_id` | [uint64](#uint64) |  | id for legacy support |
| `creator` | [string](#string) |  |  |
| `data_hash` | [bytes](#bytes) |  |  |
| `instantiate_permission` | [AccessConfig](#junction.wasm.v1.AccessConfig) |  |  |






<a name="junction.wasm.v1.QueryAllContractStateRequest"></a>

### QueryAllContractStateRequest
QueryAllContractStateRequest is the request type for the
Query/AllContractState RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the address of the contract |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="junction.wasm.v1.QueryAllContractStateResponse"></a>

### QueryAllContractStateResponse
QueryAllContractStateResponse is the response type for the
Query/AllContractState RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `models` | [Model](#junction.wasm.v1.Model) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="junction.wasm.v1.QueryBuildAddressRequest"></a>

### QueryBuildAddressRequest
QueryBuildAddressRequest is the request type for the Query/BuildAddress RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_hash` | [string](#string) |  | CodeHash is the hash of the code |
| `creator_address` | [string](#string) |  | CreatorAddress is the address of the contract instantiator |
| `salt` | [string](#string) |  | Salt is a hex encoded salt |
| `init_args` | [bytes](#bytes) |  | InitArgs are optional json encoded init args to be used in contract address building if provided |






<a name="junction.wasm.v1.QueryBuildAddressResponse"></a>

### QueryBuildAddressResponse
QueryBuildAddressResponse is the response type for the Query/BuildAddress RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | Address is the contract address |






<a name="junction.wasm.v1.QueryCodeRequest"></a>

### QueryCodeRequest
QueryCodeRequest is the request type for the Query/Code RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_id` | [uint64](#uint64) |  | grpc-gateway_out does not support Go style CodID |






<a name="junction.wasm.v1.QueryCodeResponse"></a>

### QueryCodeResponse
QueryCodeResponse is the response type for the Query/Code RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_info` | [CodeInfoResponse](#junction.wasm.v1.CodeInfoResponse) |  |  |
| `data` | [bytes](#bytes) |  |  |






<a name="junction.wasm.v1.QueryCodesRequest"></a>

### QueryCodesRequest
QueryCodesRequest is the request type for the Query/Codes RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="junction.wasm.v1.QueryCodesResponse"></a>

### QueryCodesResponse
QueryCodesResponse is the response type for the Query/Codes RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_infos` | [CodeInfoResponse](#junction.wasm.v1.CodeInfoResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="junction.wasm.v1.QueryContractHistoryRequest"></a>

### QueryContractHistoryRequest
QueryContractHistoryRequest is the request type for the Query/ContractHistory
RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the address of the contract to query |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="junction.wasm.v1.QueryContractHistoryResponse"></a>

### QueryContractHistoryResponse
QueryContractHistoryResponse is the response type for the
Query/ContractHistory RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `entries` | [ContractCodeHistoryEntry](#junction.wasm.v1.ContractCodeHistoryEntry) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="junction.wasm.v1.QueryContractInfoRequest"></a>

### QueryContractInfoRequest
QueryContractInfoRequest is the request type for the Query/ContractInfo RPC
method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the address of the contract to query |






<a name="junction.wasm.v1.QueryContractInfoResponse"></a>

### QueryContractInfoResponse
QueryContractInfoResponse is the response type for the Query/ContractInfo RPC
method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the address of the contract |
| `contract_info` | [ContractInfo](#junction.wasm.v1.ContractInfo) |  |  |






<a name="junction.wasm.v1.QueryContractsByCodeRequest"></a>

### QueryContractsByCodeRequest
QueryContractsByCodeRequest is the request type for the Query/ContractsByCode
RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_id` | [uint64](#uint64) |  | grpc-gateway_out does not support Go style CodID |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="junction.wasm.v1.QueryContractsByCodeResponse"></a>

### QueryContractsByCodeResponse
QueryContractsByCodeResponse is the response type for the
Query/ContractsByCode RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contracts` | [string](#string) | repeated | contracts are a set of contract addresses |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="junction.wasm.v1.QueryContractsByCreatorRequest"></a>

### QueryContractsByCreatorRequest
QueryContractsByCreatorRequest is the request type for the
Query/ContractsByCreator RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator_address` | [string](#string) |  | CreatorAddress is the address of contract creator |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination defines an optional pagination for the request. |






<a name="junction.wasm.v1.QueryContractsByCreatorResponse"></a>

### QueryContractsByCreatorResponse
QueryContractsByCreatorResponse is the response type for the
Query/ContractsByCreator RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_addresses` | [string](#string) | repeated | ContractAddresses result set |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | Pagination defines the pagination in the response. |






<a name="junction.wasm.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="junction.wasm.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#junction.wasm.v1.Params) |  | params defines the parameters of the module. |






<a name="junction.wasm.v1.QueryPinnedCodesRequest"></a>

### QueryPinnedCodesRequest
QueryPinnedCodesRequest is the request type for the Query/PinnedCodes
RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="junction.wasm.v1.QueryPinnedCodesResponse"></a>

### QueryPinnedCodesResponse
QueryPinnedCodesResponse is the response type for the
Query/PinnedCodes RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_ids` | [uint64](#uint64) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="junction.wasm.v1.QueryRawContractStateRequest"></a>

### QueryRawContractStateRequest
QueryRawContractStateRequest is the request type for the
Query/RawContractState RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the address of the contract |
| `query_data` | [bytes](#bytes) |  |  |






<a name="junction.wasm.v1.QueryRawContractStateResponse"></a>

### QueryRawContractStateResponse
QueryRawContractStateResponse is the response type for the
Query/RawContractState RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `data` | [bytes](#bytes) |  | Data contains the raw store data |






<a name="junction.wasm.v1.QuerySmartContractStateRequest"></a>

### QuerySmartContractStateRequest
QuerySmartContractStateRequest is the request type for the
Query/SmartContractState RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the address of the contract |
| `query_data` | [bytes](#bytes) |  | QueryData contains the query data passed to the contract |






<a name="junction.wasm.v1.QuerySmartContractStateResponse"></a>

### QuerySmartContractStateResponse
QuerySmartContractStateResponse is the response type for the
Query/SmartContractState RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `data` | [bytes](#bytes) |  | Data contains the json data returned from the smart contract |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="junction.wasm.v1.Query"></a>

### Query
Query provides defines the gRPC querier service

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ContractInfo` | [QueryContractInfoRequest](#junction.wasm.v1.QueryContractInfoRequest) | [QueryContractInfoResponse](#junction.wasm.v1.QueryContractInfoResponse) | ContractInfo gets the contract meta data | GET|/junction/wasm/v1/contract/{address}|
| `ContractHistory` | [QueryContractHistoryRequest](#junction.wasm.v1.QueryContractHistoryRequest) | [QueryContractHistoryResponse](#junction.wasm.v1.QueryContractHistoryResponse) | ContractHistory gets the contract code history | GET|/junction/wasm/v1/contract/{address}/history|
| `ContractsByCode` | [QueryContractsByCodeRequest](#junction.wasm.v1.QueryContractsByCodeRequest) | [QueryContractsByCodeResponse](#junction.wasm.v1.QueryContractsByCodeResponse) | ContractsByCode lists all smart contracts for a code id | GET|/junction/wasm/v1/code/{code_id}/contracts|
| `AllContractState` | [QueryAllContractStateRequest](#junction.wasm.v1.QueryAllContractStateRequest) | [QueryAllContractStateResponse](#junction.wasm.v1.QueryAllContractStateResponse) | AllContractState gets all raw store data for a single contract | GET|/junction/wasm/v1/contract/{address}/state|
| `RawContractState` | [QueryRawContractStateRequest](#junction.wasm.v1.QueryRawContractStateRequest) | [QueryRawContractStateResponse](#junction.wasm.v1.QueryRawContractStateResponse) | RawContractState gets single key from the raw store data of a contract | GET|/junction/wasm/v1/contract/{address}/raw/{query_data}|
| `SmartContractState` | [QuerySmartContractStateRequest](#junction.wasm.v1.QuerySmartContractStateRequest) | [QuerySmartContractStateResponse](#junction.wasm.v1.QuerySmartContractStateResponse) | SmartContractState get smart query result from the contract | GET|/junction/wasm/v1/contract/{address}/smart/{query_data}|
| `Code` | [QueryCodeRequest](#junction.wasm.v1.QueryCodeRequest) | [QueryCodeResponse](#junction.wasm.v1.QueryCodeResponse) | Code gets the binary code and metadata for a singe wasm code | GET|/junction/wasm/v1/code/{code_id}|
| `Codes` | [QueryCodesRequest](#junction.wasm.v1.QueryCodesRequest) | [QueryCodesResponse](#junction.wasm.v1.QueryCodesResponse) | Codes gets the metadata for all stored wasm codes | GET|/junction/wasm/v1/code|
| `PinnedCodes` | [QueryPinnedCodesRequest](#junction.wasm.v1.QueryPinnedCodesRequest) | [QueryPinnedCodesResponse](#junction.wasm.v1.QueryPinnedCodesResponse) | PinnedCodes gets the pinned code ids | GET|/junction/wasm/v1/codes/pinned|
| `Params` | [QueryParamsRequest](#junction.wasm.v1.QueryParamsRequest) | [QueryParamsResponse](#junction.wasm.v1.QueryParamsResponse) | Params gets the module params | GET|/junction/wasm/v1/codes/params|
| `ContractsByCreator` | [QueryContractsByCreatorRequest](#junction.wasm.v1.QueryContractsByCreatorRequest) | [QueryContractsByCreatorResponse](#junction.wasm.v1.QueryContractsByCreatorResponse) | ContractsByCreator gets the contracts by creator | GET|/junction/wasm/v1/contracts/creator/{creator_address}|
| `BuildAddress` | [QueryBuildAddressRequest](#junction.wasm.v1.QueryBuildAddressRequest) | [QueryBuildAddressResponse](#junction.wasm.v1.QueryBuildAddressResponse) | BuildAddress builds a contract address | GET|/junction/wasm/v1/contract/build_address|

 <!-- end services -->



<a name="junction/wasm/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## junction/wasm/v1/tx.proto



<a name="junction.wasm.v1.MsgAddCodeUploadParamsAddresses"></a>

### MsgAddCodeUploadParamsAddresses
MsgAddCodeUploadParamsAddresses is the
MsgAddCodeUploadParamsAddresses request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | Authority is the address of the governance account. |
| `addresses` | [string](#string) | repeated |  |






<a name="junction.wasm.v1.MsgAddCodeUploadParamsAddressesResponse"></a>

### MsgAddCodeUploadParamsAddressesResponse
MsgAddCodeUploadParamsAddressesResponse defines the response
structure for executing a MsgAddCodeUploadParamsAddresses message.






<a name="junction.wasm.v1.MsgClearAdmin"></a>

### MsgClearAdmin
MsgClearAdmin removes any admin stored for a smart contract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  | Sender is the actor that signed the messages |
| `contract` | [string](#string) |  | Contract is the address of the smart contract |






<a name="junction.wasm.v1.MsgClearAdminResponse"></a>

### MsgClearAdminResponse
MsgClearAdminResponse returns empty data






<a name="junction.wasm.v1.MsgExecuteContract"></a>

### MsgExecuteContract
MsgExecuteContract submits the given message data to a smart contract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  | Sender is the that actor that signed the messages |
| `contract` | [string](#string) |  | Contract is the address of the smart contract |
| `msg` | [bytes](#bytes) |  | Msg json encoded message to be passed to the contract |
| `funds` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Funds coins that are transferred to the contract on execution |






<a name="junction.wasm.v1.MsgExecuteContractResponse"></a>

### MsgExecuteContractResponse
MsgExecuteContractResponse returns execution result data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `data` | [bytes](#bytes) |  | Data contains bytes to returned from the contract |






<a name="junction.wasm.v1.MsgInstantiateContract"></a>

### MsgInstantiateContract
MsgInstantiateContract create a new smart contract instance for the given
code id.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  | Sender is the that actor that signed the messages |
| `admin` | [string](#string) |  | Admin is an optional address that can execute migrations |
| `code_id` | [uint64](#uint64) |  | CodeID is the reference to the stored WASM code |
| `label` | [string](#string) |  | Label is optional metadata to be stored with a contract instance. |
| `msg` | [bytes](#bytes) |  | Msg json encoded message to be passed to the contract on instantiation |
| `funds` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Funds coins that are transferred to the contract on instantiation |






<a name="junction.wasm.v1.MsgInstantiateContract2"></a>

### MsgInstantiateContract2
MsgInstantiateContract2 create a new smart contract instance for the given
code id with a predictable address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  | Sender is the that actor that signed the messages |
| `admin` | [string](#string) |  | Admin is an optional address that can execute migrations |
| `code_id` | [uint64](#uint64) |  | CodeID is the reference to the stored WASM code |
| `label` | [string](#string) |  | Label is optional metadata to be stored with a contract instance. |
| `msg` | [bytes](#bytes) |  | Msg json encoded message to be passed to the contract on instantiation |
| `funds` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Funds coins that are transferred to the contract on instantiation |
| `salt` | [bytes](#bytes) |  | Salt is an arbitrary value provided by the sender. Size can be 1 to 64. |
| `fix_msg` | [bool](#bool) |  | FixMsg include the msg value into the hash for the predictable address. Default is false |






<a name="junction.wasm.v1.MsgInstantiateContract2Response"></a>

### MsgInstantiateContract2Response
MsgInstantiateContract2Response return instantiation result data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | Address is the bech32 address of the new contract instance. |
| `data` | [bytes](#bytes) |  | Data contains bytes to returned from the contract |






<a name="junction.wasm.v1.MsgInstantiateContractResponse"></a>

### MsgInstantiateContractResponse
MsgInstantiateContractResponse return instantiation result data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | Address is the bech32 address of the new contract instance. |
| `data` | [bytes](#bytes) |  | Data contains bytes to returned from the contract |






<a name="junction.wasm.v1.MsgMigrateContract"></a>

### MsgMigrateContract
MsgMigrateContract runs a code upgrade/ downgrade for a smart contract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  | Sender is the that actor that signed the messages |
| `contract` | [string](#string) |  | Contract is the address of the smart contract |
| `code_id` | [uint64](#uint64) |  | CodeID references the new WASM code |
| `msg` | [bytes](#bytes) |  | Msg json encoded message to be passed to the contract on migration |






<a name="junction.wasm.v1.MsgMigrateContractResponse"></a>

### MsgMigrateContractResponse
MsgMigrateContractResponse returns contract migration result data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `data` | [bytes](#bytes) |  | Data contains same raw bytes returned as data from the wasm contract. (May be empty) |






<a name="junction.wasm.v1.MsgPinCodes"></a>

### MsgPinCodes
MsgPinCodes is the MsgPinCodes request type.

Since: 0.40


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | Authority is the address of the governance account. |
| `code_ids` | [uint64](#uint64) | repeated | CodeIDs references the new WASM codes |






<a name="junction.wasm.v1.MsgPinCodesResponse"></a>

### MsgPinCodesResponse
MsgPinCodesResponse defines the response structure for executing a
MsgPinCodes message.

Since: 0.40






<a name="junction.wasm.v1.MsgRemoveCodeUploadParamsAddresses"></a>

### MsgRemoveCodeUploadParamsAddresses
MsgRemoveCodeUploadParamsAddresses is the
MsgRemoveCodeUploadParamsAddresses request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | Authority is the address of the governance account. |
| `addresses` | [string](#string) | repeated |  |






<a name="junction.wasm.v1.MsgRemoveCodeUploadParamsAddressesResponse"></a>

### MsgRemoveCodeUploadParamsAddressesResponse
MsgRemoveCodeUploadParamsAddressesResponse defines the response
structure for executing a MsgRemoveCodeUploadParamsAddresses message.






<a name="junction.wasm.v1.MsgStoreAndInstantiateContract"></a>

### MsgStoreAndInstantiateContract
MsgStoreAndInstantiateContract is the MsgStoreAndInstantiateContract
request type.

Since: 0.40


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | Authority is the address of the governance account. |
| `wasm_byte_code` | [bytes](#bytes) |  | WASMByteCode can be raw or gzip compressed |
| `instantiate_permission` | [AccessConfig](#junction.wasm.v1.AccessConfig) |  | InstantiatePermission to apply on contract creation, optional |
| `unpin_code` | [bool](#bool) |  | UnpinCode code on upload, optional. As default the uploaded contract is pinned to cache. |
| `admin` | [string](#string) |  | Admin is an optional address that can execute migrations |
| `label` | [string](#string) |  | Label is optional metadata to be stored with a constract instance. |
| `msg` | [bytes](#bytes) |  | Msg json encoded message to be passed to the contract on instantiation |
| `funds` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Funds coins that are transferred from the authority account to the contract on instantiation |
| `source` | [string](#string) |  | Source is the URL where the code is hosted |
| `builder` | [string](#string) |  | Builder is the docker image used to build the code deterministically, used for smart contract verification |
| `code_hash` | [bytes](#bytes) |  | CodeHash is the SHA256 sum of the code outputted by builder, used for smart contract verification |






<a name="junction.wasm.v1.MsgStoreAndInstantiateContractResponse"></a>

### MsgStoreAndInstantiateContractResponse
MsgStoreAndInstantiateContractResponse defines the response structure
for executing a MsgStoreAndInstantiateContract message.

Since: 0.40


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | Address is the bech32 address of the new contract instance. |
| `data` | [bytes](#bytes) |  | Data contains bytes to returned from the contract |






<a name="junction.wasm.v1.MsgStoreAndMigrateContract"></a>

### MsgStoreAndMigrateContract
MsgStoreAndMigrateContract is the MsgStoreAndMigrateContract
request type.

Since: 0.42


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | Authority is the address of the governance account. |
| `wasm_byte_code` | [bytes](#bytes) |  | WASMByteCode can be raw or gzip compressed |
| `instantiate_permission` | [AccessConfig](#junction.wasm.v1.AccessConfig) |  | InstantiatePermission to apply on contract creation, optional |
| `contract` | [string](#string) |  | Contract is the address of the smart contract |
| `msg` | [bytes](#bytes) |  | Msg json encoded message to be passed to the contract on migration |






<a name="junction.wasm.v1.MsgStoreAndMigrateContractResponse"></a>

### MsgStoreAndMigrateContractResponse
MsgStoreAndMigrateContractResponse defines the response structure
for executing a MsgStoreAndMigrateContract message.

Since: 0.42


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_id` | [uint64](#uint64) |  | CodeID is the reference to the stored WASM code |
| `checksum` | [bytes](#bytes) |  | Checksum is the sha256 hash of the stored code |
| `data` | [bytes](#bytes) |  | Data contains bytes to returned from the contract |






<a name="junction.wasm.v1.MsgStoreCode"></a>

### MsgStoreCode
MsgStoreCode submit Wasm code to the system


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  | Sender is the actor that signed the messages |
| `wasm_byte_code` | [bytes](#bytes) |  | WASMByteCode can be raw or gzip compressed |
| `instantiate_permission` | [AccessConfig](#junction.wasm.v1.AccessConfig) |  | InstantiatePermission access control to apply on contract creation, optional |






<a name="junction.wasm.v1.MsgStoreCodeResponse"></a>

### MsgStoreCodeResponse
MsgStoreCodeResponse returns store result data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_id` | [uint64](#uint64) |  | CodeID is the reference to the stored WASM code |
| `checksum` | [bytes](#bytes) |  | Checksum is the sha256 hash of the stored code |






<a name="junction.wasm.v1.MsgSudoContract"></a>

### MsgSudoContract
MsgSudoContract is the MsgSudoContract request type.

Since: 0.40


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | Authority is the address of the governance account. |
| `contract` | [string](#string) |  | Contract is the address of the smart contract |
| `msg` | [bytes](#bytes) |  | Msg json encoded message to be passed to the contract as sudo |






<a name="junction.wasm.v1.MsgSudoContractResponse"></a>

### MsgSudoContractResponse
MsgSudoContractResponse defines the response structure for executing a
MsgSudoContract message.

Since: 0.40


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `data` | [bytes](#bytes) |  | Data contains bytes to returned from the contract |






<a name="junction.wasm.v1.MsgUnpinCodes"></a>

### MsgUnpinCodes
MsgUnpinCodes is the MsgUnpinCodes request type.

Since: 0.40


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | Authority is the address of the governance account. |
| `code_ids` | [uint64](#uint64) | repeated | CodeIDs references the WASM codes |






<a name="junction.wasm.v1.MsgUnpinCodesResponse"></a>

### MsgUnpinCodesResponse
MsgUnpinCodesResponse defines the response structure for executing a
MsgUnpinCodes message.

Since: 0.40






<a name="junction.wasm.v1.MsgUpdateAdmin"></a>

### MsgUpdateAdmin
MsgUpdateAdmin sets a new admin for a smart contract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  | Sender is the that actor that signed the messages |
| `new_admin` | [string](#string) |  | NewAdmin address to be set |
| `contract` | [string](#string) |  | Contract is the address of the smart contract |






<a name="junction.wasm.v1.MsgUpdateAdminResponse"></a>

### MsgUpdateAdminResponse
MsgUpdateAdminResponse returns empty data






<a name="junction.wasm.v1.MsgUpdateContractLabel"></a>

### MsgUpdateContractLabel
MsgUpdateContractLabel sets a new label for a smart contract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  | Sender is the that actor that signed the messages |
| `new_label` | [string](#string) |  | NewLabel string to be set |
| `contract` | [string](#string) |  | Contract is the address of the smart contract |






<a name="junction.wasm.v1.MsgUpdateContractLabelResponse"></a>

### MsgUpdateContractLabelResponse
MsgUpdateContractLabelResponse returns empty data






<a name="junction.wasm.v1.MsgUpdateInstantiateConfig"></a>

### MsgUpdateInstantiateConfig
MsgUpdateInstantiateConfig updates instantiate config for a smart contract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  | Sender is the that actor that signed the messages |
| `code_id` | [uint64](#uint64) |  | CodeID references the stored WASM code |
| `new_instantiate_permission` | [AccessConfig](#junction.wasm.v1.AccessConfig) |  | NewInstantiatePermission is the new access control |






<a name="junction.wasm.v1.MsgUpdateInstantiateConfigResponse"></a>

### MsgUpdateInstantiateConfigResponse
MsgUpdateInstantiateConfigResponse returns empty data






<a name="junction.wasm.v1.MsgUpdateParams"></a>

### MsgUpdateParams
MsgUpdateParams is the MsgUpdateParams request type.

Since: 0.40


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | Authority is the address of the governance account. |
| `params` | [Params](#junction.wasm.v1.Params) |  | params defines the x/wasm parameters to update.

NOTE: All parameters must be supplied. |






<a name="junction.wasm.v1.MsgUpdateParamsResponse"></a>

### MsgUpdateParamsResponse
MsgUpdateParamsResponse defines the response structure for executing a
MsgUpdateParams message.

Since: 0.40





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="junction.wasm.v1.Msg"></a>

### Msg
Msg defines the wasm Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `StoreCode` | [MsgStoreCode](#junction.wasm.v1.MsgStoreCode) | [MsgStoreCodeResponse](#junction.wasm.v1.MsgStoreCodeResponse) | StoreCode to submit Wasm code to the system | |
| `InstantiateContract` | [MsgInstantiateContract](#junction.wasm.v1.MsgInstantiateContract) | [MsgInstantiateContractResponse](#junction.wasm.v1.MsgInstantiateContractResponse) | InstantiateContract creates a new smart contract instance for the given code id. | |
| `InstantiateContract2` | [MsgInstantiateContract2](#junction.wasm.v1.MsgInstantiateContract2) | [MsgInstantiateContract2Response](#junction.wasm.v1.MsgInstantiateContract2Response) | InstantiateContract2 creates a new smart contract instance for the given code id with a predictable address | |
| `ExecuteContract` | [MsgExecuteContract](#junction.wasm.v1.MsgExecuteContract) | [MsgExecuteContractResponse](#junction.wasm.v1.MsgExecuteContractResponse) | Execute submits the given message data to a smart contract | |
| `MigrateContract` | [MsgMigrateContract](#junction.wasm.v1.MsgMigrateContract) | [MsgMigrateContractResponse](#junction.wasm.v1.MsgMigrateContractResponse) | Migrate runs a code upgrade/ downgrade for a smart contract | |
| `UpdateAdmin` | [MsgUpdateAdmin](#junction.wasm.v1.MsgUpdateAdmin) | [MsgUpdateAdminResponse](#junction.wasm.v1.MsgUpdateAdminResponse) | UpdateAdmin sets a new admin for a smart contract | |
| `ClearAdmin` | [MsgClearAdmin](#junction.wasm.v1.MsgClearAdmin) | [MsgClearAdminResponse](#junction.wasm.v1.MsgClearAdminResponse) | ClearAdmin removes any admin stored for a smart contract | |
| `UpdateInstantiateConfig` | [MsgUpdateInstantiateConfig](#junction.wasm.v1.MsgUpdateInstantiateConfig) | [MsgUpdateInstantiateConfigResponse](#junction.wasm.v1.MsgUpdateInstantiateConfigResponse) | UpdateInstantiateConfig updates instantiate config for a smart contract | |
| `UpdateParams` | [MsgUpdateParams](#junction.wasm.v1.MsgUpdateParams) | [MsgUpdateParamsResponse](#junction.wasm.v1.MsgUpdateParamsResponse) | UpdateParams defines a governance operation for updating the x/wasm module parameters. The authority is defined in the keeper.

Since: 0.40 | |
| `SudoContract` | [MsgSudoContract](#junction.wasm.v1.MsgSudoContract) | [MsgSudoContractResponse](#junction.wasm.v1.MsgSudoContractResponse) | SudoContract defines a governance operation for calling sudo on a contract. The authority is defined in the keeper.

Since: 0.40 | |
| `PinCodes` | [MsgPinCodes](#junction.wasm.v1.MsgPinCodes) | [MsgPinCodesResponse](#junction.wasm.v1.MsgPinCodesResponse) | PinCodes defines a governance operation for pinning a set of code ids in the wasmvm cache. The authority is defined in the keeper.

Since: 0.40 | |
| `UnpinCodes` | [MsgUnpinCodes](#junction.wasm.v1.MsgUnpinCodes) | [MsgUnpinCodesResponse](#junction.wasm.v1.MsgUnpinCodesResponse) | UnpinCodes defines a governance operation for unpinning a set of code ids in the wasmvm cache. The authority is defined in the keeper.

Since: 0.40 | |
| `StoreAndInstantiateContract` | [MsgStoreAndInstantiateContract](#junction.wasm.v1.MsgStoreAndInstantiateContract) | [MsgStoreAndInstantiateContractResponse](#junction.wasm.v1.MsgStoreAndInstantiateContractResponse) | StoreAndInstantiateContract defines a governance operation for storing and instantiating the contract. The authority is defined in the keeper.

Since: 0.40 | |
| `RemoveCodeUploadParamsAddresses` | [MsgRemoveCodeUploadParamsAddresses](#junction.wasm.v1.MsgRemoveCodeUploadParamsAddresses) | [MsgRemoveCodeUploadParamsAddressesResponse](#junction.wasm.v1.MsgRemoveCodeUploadParamsAddressesResponse) | RemoveCodeUploadParamsAddresses defines a governance operation for removing addresses from code upload params. The authority is defined in the keeper. | |
| `AddCodeUploadParamsAddresses` | [MsgAddCodeUploadParamsAddresses](#junction.wasm.v1.MsgAddCodeUploadParamsAddresses) | [MsgAddCodeUploadParamsAddressesResponse](#junction.wasm.v1.MsgAddCodeUploadParamsAddressesResponse) | AddCodeUploadParamsAddresses defines a governance operation for adding addresses to code upload params. The authority is defined in the keeper. | |
| `StoreAndMigrateContract` | [MsgStoreAndMigrateContract](#junction.wasm.v1.MsgStoreAndMigrateContract) | [MsgStoreAndMigrateContractResponse](#junction.wasm.v1.MsgStoreAndMigrateContractResponse) | StoreAndMigrateContract defines a governance operation for storing and migrating the contract. The authority is defined in the keeper.

Since: 0.42 | |
| `UpdateContractLabel` | [MsgUpdateContractLabel](#junction.wasm.v1.MsgUpdateContractLabel) | [MsgUpdateContractLabelResponse](#junction.wasm.v1.MsgUpdateContractLabelResponse) | UpdateContractLabel sets a new label for a smart contract

Since: 0.43 | |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

