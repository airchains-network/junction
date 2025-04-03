package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"

	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/rollup/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// InitRollup initializes a new rollup with the provided parameters.
// It creates two separate stores: rollupDataStore and rollupMonikerStore.
//
// rollupDataStore is used to store the rollup metadata associated with the ChainId,
// allowing for efficient retrieval of rollup information based on the unique ChainId.
//
// rollupMonikerStore is used to ensure that each rollup has a unique moniker.
// This prevents the creation of multiple rollups with the same moniker, which could lead
// to confusion and data integrity issues. By checking this store before creating a new rollup,
// we can enforce this uniqueness constraint.
//
// Parameters:
// - goCtx: The context for the operation, which includes metadata about the current state.
// - msg: The message containing the details for the rollup to be initialized.
//
// Returns:
// - A pointer to MsgInitRollupResponse containing the RollupId and a status indicating success or failure.
// - An error if the rollup with the given moniker already exists or if any other issue occurs during initialization.

func (k msgServer) InitRollup(goCtx context.Context, msg *types.MsgInitRollup) (*types.MsgInitRollupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// rollupStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RollupKey))
	rollupDataStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RollupDataKey))
	rollupMonikerStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RollupMonikerKey))

	var creator = msg.Creator
	var moniker = msg.Moniker
	var chainId = msg.ChainId
	var denomName = msg.DenomName
	var keys = msg.Keys
	var supply = msg.Supply
	var daType = msg.DaType
	var aclContractAddress = msg.AclContractAddress
	var kmsVerifierAddress = msg.KmsVerifierAddress
	var tfheExecutorAddress = msg.TfheExecutorAddress
	var gatewayContractAddress = msg.GatewayContractAddress
	var ascContractAddress = msg.AscContractAddress
	var relayerGAddress = msg.RelayerGAddress
	var relayerASCAddress = msg.RelayerASCAddress
	// Generate RollupId using SHA-256
	blockNumber := ctx.BlockHeight()
	timestamp := ctx.BlockTime().String()
	hash := sha256.New()
	hash.Write([]byte(moniker + strconv.FormatInt(blockNumber, 10) + timestamp))
	rollupId := hex.EncodeToString(hash.Sum(nil))

	// check if the passed wallet address is a valid ethereum address
	if !IsValidEthereumAddress(aclContractAddress) || !IsValidEthereumAddress(kmsVerifierAddress) || !IsValidEthereumAddress(tfheExecutorAddress) || !IsValidEthereumAddress(gatewayContractAddress) || !IsValidEthereumAddress(relayerGAddress) {
		return nil, status.Error(codes.InvalidArgument, "invalid ethereum address")
	}

	// check if the passed relayer asc address is a valid cosmos address with the prefix "air"
	if !IsValidCosmosWalletAndContractAddress(relayerASCAddress) {
		return nil, status.Error(codes.InvalidArgument, "invalid relayer asc address")
	}

	// check if the passed ascContractAddress is a valid cosmos wasm contract address with the prefix "air"
	if !IsValidCosmosWalletAndContractAddress(ascContractAddress) {
		return nil, status.Error(codes.InvalidArgument, "invalid asc contract address")
	}
	// write

	var rollup = types.RollupMetadata{
		CreatedBy:                    creator,
		RollupId:                     rollupId,
		RollupLatestBatchNo:          0,
		RollupLatestFinalizedBatchNo: 0,
		Moniker:                      moniker,
		ChainId:                      chainId,
		DenomName:                    denomName,
		Keys:                         keys,
		Supply:                       supply,
		DaType:                       daType,
		DaLatestHash:                 "none",
		ProverType:                   "none",
		ProverVerificationKey:        nil,
		ProverEndpoint:               "none",
		AclContractAddress:           aclContractAddress,
		KmsVerifierAddress:           kmsVerifierAddress,
		TfheExecutorAddress:          tfheExecutorAddress,
		GatewayContractAddress:       gatewayContractAddress,
		AscContractAddress:           ascContractAddress,
		RelayerGAddress:              relayerGAddress,
		RelayerASCAddress:            relayerASCAddress,
	}

	if rollupMonikerStore.Has([]byte(moniker)) {
		return nil, status.Error(codes.AlreadyExists, "rollup already exists with this moniker")
	}

	// Store rollup metadata
	rollupBytes := k.cdc.MustMarshal(&rollup)
	rollupDataStore.Set([]byte(rollup.RollupId), rollupBytes)
	rollupMonikerStore.Set([]byte(rollup.Moniker), []byte(rollup.RollupId))

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		"rollup-initialized",
		sdk.NewAttribute("creator", rollup.CreatedBy),
		sdk.NewAttribute("rollup-id", rollupId),
		sdk.NewAttribute("moniker", rollup.Moniker),
		sdk.NewAttribute("chain-id", rollup.ChainId),
		sdk.NewAttribute("denom-name", rollup.DenomName),
		sdk.NewAttribute("da-type", rollup.DaType),
		sdk.NewAttribute("keys", strings.Join(rollup.Keys, ",")),
		sdk.NewAttribute("supply", strings.Join(uint64SliceToStringSlice(rollup.Supply), ",")),
		sdk.NewAttribute("acl-contract-address", rollup.AclContractAddress),
		sdk.NewAttribute("kms-verifier-address", rollup.KmsVerifierAddress),
		sdk.NewAttribute("tfhe-executor-address", rollup.TfheExecutorAddress),
		sdk.NewAttribute("gateway-contract-address", rollup.GatewayContractAddress),
		sdk.NewAttribute("asc-contract-address", rollup.AscContractAddress),
		sdk.NewAttribute("relayer-g-address", rollup.RelayerGAddress),
	))

	return &types.MsgInitRollupResponse{
		RollupId: rollupId,
		Status:   true,
	}, nil
}
