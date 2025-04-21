package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"

	"cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	rolluptypes "github.com/airchains-network/junction/x/rollup/types"
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

func (k msgServer) InitRollup(goCtx context.Context, msg *rolluptypes.MsgInitRollup) (*rolluptypes.MsgInitRollupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// rollupStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RollupKey))
	rollupDataStore := prefix.NewStore(storeAdapter, rolluptypes.KeyPrefix(rolluptypes.RollupDataKey))
	rollupMonikerStore := prefix.NewStore(storeAdapter, rolluptypes.KeyPrefix(rolluptypes.RollupMonikerKey))

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

	// find the total supply of the denom
	var totalSupply uint64
	for _, supply := range supply {
		totalSupply += supply
	}

	// Check if the creator address has more than totalSupply balance
	creatorAddr, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid creator address")
	}

	totalSupplyCoins := sdk.NewCoins(sdk.NewCoin("uamf", math.NewIntFromUint64(totalSupply)))

	balance := k.bankKeeper.GetBalance(ctx, creatorAddr, "uamf")
	if balance.Amount.Uint64() < totalSupply {
		return nil, status.Error(codes.FailedPrecondition, "creator address does not have sufficient balance")
	}

	// if msg.GenesisSupply.Denom != "uamf" {
	// 	return nil, status.Error(codes.FailedPrecondition, fmt.Sprintf("invalid genesis supply denom, expected: uamf, got: %s", msg.GenesisSupply.Denom))
	// }

	// if msg.GenesisSupply.Amount.Uint64() != totalSupply {
	// 	return nil, status.Error(codes.FailedPrecondition, fmt.Sprintf("invalid genesis supply amount, expected: %d, got: %d", totalSupply, msg.GenesisSupply.Amount.Uint64()))
	// }

	var rollup = rolluptypes.RollupMetadata{
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

	// send the total supply of the denom from the creator address to the module account
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, rolluptypes.ModuleName, totalSupplyCoins)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "failed to send coins from creator address to module account")
	}

	/*
		This section checks if a ledger entry for the creator address already exists.
		If it does, it retrieves the existing rollup IDs, appends the new rollup ID, and updates the store.
		If it doesn't exist, it creates a new entry with the rollup ID.
	*/

	ledgerEntryStore := prefix.NewStore(storeAdapter, rolluptypes.KeyPrefix(rolluptypes.LedgerEntryRollupCreatorKey))
	if ledgerEntryStore.Has([]byte(creatorAddr.String())) {
		rollupsIdBytes := ledgerEntryStore.Get([]byte(creatorAddr.String()))
		var rollupsIds []string
		err = json.Unmarshal(rollupsIdBytes, &rollupsIds)
		if err != nil {
			return nil, status.Error(codes.FailedPrecondition, "failed to unmarshal rollups ids")
		}

		rollupsIds = append(rollupsIds, rollupId)
		newRollupsIdBytes, err := json.Marshal(rollupsIds)
		if err != nil {
			return nil, status.Error(codes.FailedPrecondition, "failed to marshal rollups ids")
		}
		ledgerEntryStore.Set([]byte(creatorAddr.String()), newRollupsIdBytes)
	} else {
		rollupsIds := []string{rollupId}
		newRollupsIdBytes, err := json.Marshal(rollupsIds)
		if err != nil {
			return nil, status.Error(codes.FailedPrecondition, "failed to marshal rollups ids")
		}
		ledgerEntryStore.Set([]byte(creatorAddr.String()), newRollupsIdBytes)
	}

	// update the ledger entry for the creator address
	ledgerEntry := rolluptypes.LedgerEntry{
		CreatorAddress: creatorAddr.String(),
		AmountStaked:   totalSupply,
		Denom:          denomName,
		RollupId:       rollupId,
		Timestamp:      ctx.BlockTime().String(),
		BlockHeight:    uint64(ctx.BlockHeight()),
	}
	ledgerEntryBytes := k.cdc.MustMarshal(&ledgerEntry)

	rollupStakingInfoStore := prefix.NewStore(storeAdapter, rolluptypes.KeyPrefix(rolluptypes.RollupStakingInfoKey))
	rollupStakingInfoKey := []byte(rollupId)
	rollupStakingInfoBytes := rollupStakingInfoStore.Get(rollupStakingInfoKey)
	if rollupStakingInfoBytes == nil {
		rollupStakingInfoStore.Set(rollupStakingInfoKey, ledgerEntryBytes)
	} else {
		return nil, status.Error(codes.FailedPrecondition, "rollup staking info already exists")
	}

	// Store rollup metadata
	rollupBytes := k.cdc.MustMarshal(&rollup)
	rollupDataStore.Set([]byte(rollup.RollupId), rollupBytes)
	rollupMonikerStore.Set([]byte(rollup.Moniker), []byte(rollup.RollupId))

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		rolluptypes.EventTypeTokenLocked,
		sdk.NewAttribute(rolluptypes.AttributeKeyCreator, rollup.CreatedBy),
		sdk.NewAttribute(rolluptypes.AttributeKeyRollupId, rollupId),
		sdk.NewAttribute(sdk.AttributeKeyAmount, strconv.FormatUint(totalSupply, 10)),
		sdk.NewAttribute(rolluptypes.AttributeKeyDenom, rollup.DenomName),
	))

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		rolluptypes.EventTypeRollupInitialized,
		sdk.NewAttribute(rolluptypes.AttributeKeyCreator, rollup.CreatedBy),
		sdk.NewAttribute(rolluptypes.AttributeKeyRollupId, rollupId),
		sdk.NewAttribute(rolluptypes.AttributeKeyMoniker, rollup.Moniker),
		sdk.NewAttribute(rolluptypes.AttributeKeyChainId, rollup.ChainId),
		sdk.NewAttribute(rolluptypes.AttributeKeyDenomName, rollup.DenomName),
		sdk.NewAttribute(rolluptypes.AttributeKeyDaType, rollup.DaType),
		sdk.NewAttribute(rolluptypes.AttributeKeyKeys, strings.Join(rollup.Keys, ",")),
		sdk.NewAttribute(rolluptypes.AttributeKeySupply, strings.Join(uint64SliceToStringSlice(rollup.Supply), ",")),
		sdk.NewAttribute(rolluptypes.AttributeKeyTotalSupply, strconv.FormatUint(totalSupply, 10)),
		sdk.NewAttribute(rolluptypes.AttributeKeyAclContractAddress, rollup.AclContractAddress),
		sdk.NewAttribute(rolluptypes.AttributeKeyKmsVerifierAddress, rollup.KmsVerifierAddress),
		sdk.NewAttribute(rolluptypes.AttributeKeyTfheExecutorAddress, rollup.TfheExecutorAddress),
		sdk.NewAttribute(rolluptypes.AttributeKeyGatewayContractAddress, rollup.GatewayContractAddress),
		sdk.NewAttribute(rolluptypes.AttributeKeyAscContractAddress, rollup.AscContractAddress),
		sdk.NewAttribute(rolluptypes.AttributeKeyRelayerGAddress, rollup.RelayerGAddress),
	))

	return &rolluptypes.MsgInitRollupResponse{
		RollupId: rollupId,
		Status:   true,
	}, nil
}
