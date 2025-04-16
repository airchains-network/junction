package keeper

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"cosmossdk.io/store/prefix"
	bls12381 "github.com/airchains-network/gnark/backend/groth16/bls12-381"
	"github.com/airchains-network/junction/x/rollup/types"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRollupBatchDbStoreKeys(rollupId string, batchNo uint64) (string, []byte) {
	batchStorePath := "batches/" + rollupId
	batchNoString := strconv.FormatUint(batchNo, 10)
	batchKeyByte := []byte(batchNoString)
	return batchStorePath, batchKeyByte
}

func (k Keeper) getRollupByMoniker(ctx sdk.Context, moniker string) (types.RollupMetadata, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	rollupRegistryStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RollupMonikerKey))
	rollupIdBytes := rollupRegistryStore.Get([]byte(moniker))
	rollupDataStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RollupDataKey))
	rollupBytes := rollupDataStore.Get(rollupIdBytes)
	var rollup types.RollupMetadata
	if rollupBytes == nil {
		return types.RollupMetadata{}, status.Error(codes.InvalidArgument, "rollup does not exist")
	}
	k.cdc.MustUnmarshal(rollupBytes, &rollup)
	return rollup, nil
}

func (k Keeper) getRollupById(ctx sdk.Context, rollupId string) (types.RollupMetadata, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	rollupDataStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RollupDataKey))
	rollupBytes := rollupDataStore.Get([]byte(rollupId))
	var rollup types.RollupMetadata
	if rollupBytes == nil {
		return types.RollupMetadata{}, status.Error(codes.InvalidArgument, "rollup does not exist")
	}
	k.cdc.MustUnmarshal(rollupBytes, &rollup)
	return rollup, nil
}

func (k Keeper) getRollupBatchInfoByBatchNo(ctx sdk.Context, rollupId string, batchNo uint64) (types.Batch, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	storePath, batchKey := k.GetRollupBatchDbStoreKeys(rollupId, batchNo)
	batchDataStore := prefix.NewStore(storeAdapter, types.KeyPrefix(storePath))
	batchBytes := batchDataStore.Get(batchKey)
	var batchInfo types.Batch
	if batchBytes == nil {
		return types.Batch{}, status.Error(codes.InvalidArgument, "batch does not exist")
	}
	k.cdc.MustUnmarshal(batchBytes, &batchInfo)
	return batchInfo, nil
}

func (k Keeper) VerifyBatch(zkProof []byte, witness fr.Vector, vkBytes []byte) (bool, error) {
	// Verification Variables requirement and unmarshal codes below
	var proof *bls12381.Proof
	var vk bls12381.VerifyingKey

	proofErr := json.Unmarshal(zkProof, &proof)
	if proofErr != nil {
		return false, status.Error(codes.InvalidArgument, "invalid proof provided in argument")
	}

	unmarshalVkError := json.Unmarshal(vkBytes, &vk)
	if unmarshalVkError != nil {
		return false, status.Error(codes.Unavailable, "error in unmarshalling verification key")
	}

	verifyErr := bls12381.Verify(proof, &vk, witness)
	if verifyErr != nil {
		return false, status.Error(codes.Aborted, "verification failed"+verifyErr.Error())
	}

	return true, nil
}

func uint64SliceToStringSlice(uint64s []uint64) []string {
	stringSlice := make([]string, len(uint64s))
	for i, v := range uint64s {
		stringSlice[i] = strconv.FormatUint(v, 10)
	}
	return stringSlice
}

// IsValidEthereumAddress checks whether the passed wallet address is a valid ethereum address.
// A valid Ethereum address must:
// - Begin with "0x"
// - Be 42 characters in length (including the "0x" prefix)
// - Only contain hexadecimal characters (0-9, a-f, A-F) after the prefix
func IsValidEthereumAddress(address string) bool {
	// Check if the address starts with "0x" and has the correct length
	if len(address) != 42 || !strings.HasPrefix(address, "0x") {
		return false
	}

	// Check if all remaining characters are hex characters
	for _, c := range address[2:] {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}

	return true
}

// IsValidCosmosWalletAndContractAddress checks whether the passed address is a valid Cosmos wallet or contract address.
// A valid Cosmos address must have a human-readable part (HRP) of "air".
func IsValidCosmosWalletAndContractAddress(address string) bool {
	hrp, _, err := bech32.DecodeAndConvert(address)
	if err != nil {
		return false
	}
	return hrp == "air"
}

// ************** Ledger Entry Helper Functions **************

type TotalStakeAmountResponse struct {
	Creators []struct {
		CreatorAddress string `json:"creator_address"`
		Rollups        []struct {
			RollupID     string `json:"rollup_id"`
			AmountStaked int    `json:"amount_staked"`
			Denom        string `json:"denom"`
		} `json:"rollups"`
	} `json:"creators"`
	TotalStakedAmount int `json:"total_staked_amount"`
}

func (k Keeper) Inner_GetTotalStakedAmount(ctx sdk.Context) (TotalStakeAmountResponse, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	ledgerEntryStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LedgerEntryRollupCreatorKey))
	rollupStakingInfoStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RollupStakingInfoKey)) // for later use
	iterator := ledgerEntryStore.Iterator(nil, nil)
	defer iterator.Close()

	var totalStakedAmount int
	var creators []struct {
		CreatorAddress string `json:"creator_address"`
		Rollups        []struct {
			RollupID     string `json:"rollup_id"`
			AmountStaked int    `json:"amount_staked"`
			Denom        string `json:"denom"`
		} `json:"rollups"`
	}

	for ; iterator.Valid(); iterator.Next() {
		creatorAddress := string(iterator.Key())
		rollupIdsBytes := iterator.Value()
		var rollupIds []string
		json.Unmarshal(rollupIdsBytes, &rollupIds)

		// Create a new creator entry
		creator := struct {
			CreatorAddress string `json:"creator_address"`
			Rollups        []struct {
				RollupID     string `json:"rollup_id"`
				AmountStaked int    `json:"amount_staked"`
				Denom        string `json:"denom"`
			} `json:"rollups"`
		}{
			CreatorAddress: creatorAddress,
			Rollups: []struct {
				RollupID     string `json:"rollup_id"`
				AmountStaked int    `json:"amount_staked"`
				Denom        string `json:"denom"`
			}{},
		}

		// Process each rollup for this creator
		for _, rollupId := range rollupIds {
			rollupStakingInfoKey := []byte(rollupId)
			rollupStakingInfoBytes := rollupStakingInfoStore.Get(rollupStakingInfoKey)
			if rollupStakingInfoBytes == nil {
				continue
			}

			var rollupStakingInfo types.LedgerEntry
			k.cdc.MustUnmarshal(rollupStakingInfoBytes, &rollupStakingInfo)

			// Add to total staked amount
			totalStakedAmount += int(rollupStakingInfo.AmountStaked)

			// Add rollup info to creator's rollups
			rollupInfo := struct {
				RollupID     string `json:"rollup_id"`
				AmountStaked int    `json:"amount_staked"`
				Denom        string `json:"denom"`
			}{
				RollupID:     rollupId,
				AmountStaked: int(rollupStakingInfo.AmountStaked),
				Denom:        rollupStakingInfo.Denom,
			}

			creator.Rollups = append(creator.Rollups, rollupInfo)
		}

		// Add creator to the list if they have rollups
		if len(creator.Rollups) > 0 {
			creators = append(creators, creator)
		}
	}

	// Create and return the response
	response := TotalStakeAmountResponse{
		Creators:          creators,
		TotalStakedAmount: totalStakedAmount,
	}

	return response, nil
}

func (k Keeper) Inner_GetStakeDetailsByUser(ctx sdk.Context, address string) (types.QueryGetStakeDetailsByUserResponse, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	ledgerEntryStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LedgerEntryRollupCreatorKey))
	rollupStakingInfoStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RollupStakingInfoKey)) // for later use

	addressBytes := []byte(address)
	rollupIdsBytes := ledgerEntryStore.Get(addressBytes)
	if rollupIdsBytes == nil {
		return types.QueryGetStakeDetailsByUserResponse{}, status.Error(codes.NotFound, "no rollups found for address")
	}

	var rollupIds []string
	json.Unmarshal(rollupIdsBytes, &rollupIds)

	var totalStakedAmount int64
	var rollups []types.RollupStake
	for _, rollupId := range rollupIds {
		rollupStakingInfoKey := []byte(rollupId)
		rollupStakingInfoBytes := rollupStakingInfoStore.Get(rollupStakingInfoKey)
		if rollupStakingInfoBytes == nil {
			return types.QueryGetStakeDetailsByUserResponse{}, status.Error(codes.NotFound, fmt.Sprintf("rollup staking info not found for rollup id: %s", rollupId))
		}

		var rollupStakingInfo types.LedgerEntry
		k.cdc.MustUnmarshal(rollupStakingInfoBytes, &rollupStakingInfo)
		totalStakedAmount += int64(rollupStakingInfo.AmountStaked)

		rollups = append(rollups, types.RollupStake{
			RollupId:     rollupId,
			AmountStaked: int64(rollupStakingInfo.AmountStaked),
			Denom:        rollupStakingInfo.Denom,
		})
	}

	rollupsPtrs := make([]*types.RollupStake, len(rollups))
	for i := range rollups {
		rollupsPtrs[i] = &rollups[i]
	}

	return types.QueryGetStakeDetailsByUserResponse{
		TotalStakedAmount: totalStakedAmount,
		Rollups:           rollupsPtrs,
	}, nil
}

func (k Keeper) Inner_GetRollupStakedAmount(ctx sdk.Context, rollupId string) (types.QueryGetRollupStakedAmountResponse, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	rollupStakingInfoStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RollupStakingInfoKey))

	rollupStakingInfoKey := []byte(rollupId)
	rollupStakingInfoBytes := rollupStakingInfoStore.Get(rollupStakingInfoKey)

	if rollupStakingInfoBytes == nil {
		return types.QueryGetRollupStakedAmountResponse{}, status.Error(codes.NotFound, fmt.Sprintf("rollup staking info not found for rollup id: %s", rollupId))
	}

	var rollupStakingInfo types.LedgerEntry
	k.cdc.MustUnmarshal(rollupStakingInfoBytes, &rollupStakingInfo)

	return types.QueryGetRollupStakedAmountResponse{
		AmountStaked: uint64(rollupStakingInfo.AmountStaked),
		Denom:        rollupStakingInfo.Denom,
	}, nil
	
}
