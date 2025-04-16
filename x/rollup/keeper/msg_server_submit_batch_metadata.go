package keeper

import (
	"context"
	"strconv"

	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/rollup/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) SubmitBatchMetadata(goCtx context.Context, msg *types.MsgSubmitBatchMetadata) (*types.MsgSubmitBatchMetadataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	batchNo := msg.BatchNo
	rollupId := msg.RollupId
	daName := msg.DaName
	daCommitment := msg.DaCommitment
	daBlockHash := msg.DaHash
	daPointer := msg.DaPointer
	daNamespace := msg.DaNamespace

	batchStorePath, batchKeyByte := k.GetRollupBatchDbStoreKeys(rollupId, batchNo)
	batchDataStore := prefix.NewStore(storeAdapter, types.KeyPrefix(batchStorePath))

	// Check if the batch metadata already exists
	if batchDataStore.Has(batchKeyByte) {
		return nil, status.Error(codes.AlreadyExists, "batch metadata already exists")
	}

	var previousMerkleRootHash string

	if batchNo == 0 {
		return nil, status.Error(codes.InvalidArgument, "batch number must be greater than 0")
	}

	if batchNo < 2 {
		previousMerkleRootHash = "initial_batch_no_previous_hash" // This indicates that there is no previous Merkle root hash for the initial batch
	} else {
		var previousBatch types.Batch

		_, previousBatchNo := k.GetRollupBatchDbStoreKeys(rollupId, batchNo-1)
		previousBatchBytes := batchDataStore.Get(previousBatchNo)
		k.cdc.MustUnmarshal(previousBatchBytes, &previousBatch)
		previousMerkleRootHash = previousBatch.MerkleRootHash
	}

	// Create a new batch metadata entry
	var batchMetadata = types.Batch{
		Submitter:              msg.Creator,
		RollupId:               rollupId,
		BatchNo:                batchNo,
		MerkleRootHash:         "",
		PreviousMerkleRootHash: previousMerkleRootHash,
		ZkProof:                nil,
		PublicWitness:          nil,
		Timestamp:              ctx.BlockTime().String(),
		DaName:                 daName,
		DaCommitment:           daCommitment,
		DaHash:                 daBlockHash,
		DaPointer:              daPointer,
		DaNamespace:            daNamespace,
		IsFinalized:            false,
	}

	// Marshal and store the batch metadata
	batchDataBytes := k.cdc.MustMarshal(&batchMetadata)
	batchDataStore.Set(batchKeyByte, batchDataBytes)

	// updating rollup metadata
	rollupDataStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RollupDataKey))
	rollupBytes := rollupDataStore.Get([]byte(rollupId))
	var rollup types.RollupMetadata
	k.cdc.MustUnmarshal(rollupBytes, &rollup)

	rollup.RollupLatestBatchNo = batchNo
	rollup.DaLatestHash = daBlockHash

	rollupBytes = k.cdc.MustMarshal(&rollup)
	rollupDataStore.Set([]byte(rollupId), rollupBytes)

	// Emit an event for the new batch metadata
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		"batch-metadata-submitted",
		sdk.NewAttribute("submitter", batchMetadata.Submitter),
		sdk.NewAttribute("batch-no", strconv.FormatUint(batchMetadata.BatchNo, 10)),
		sdk.NewAttribute("timestamp", batchMetadata.Timestamp),
		sdk.NewAttribute("rollup-id", batchMetadata.RollupId),
		sdk.NewAttribute("da-name", batchMetadata.DaName),
		sdk.NewAttribute("da-commitment", batchMetadata.DaCommitment),
		sdk.NewAttribute("da-hash", batchMetadata.DaHash),
		sdk.NewAttribute("da-pointer", batchMetadata.DaPointer),
		sdk.NewAttribute("da-namespace", batchMetadata.DaNamespace),
		sdk.NewAttribute("previous-merkle-root-hash", batchMetadata.PreviousMerkleRootHash),
		sdk.NewAttribute("is-finalized", strconv.FormatBool(batchMetadata.IsFinalized)),
	))

	return &types.MsgSubmitBatchMetadataResponse{
		Status: true,
	}, nil
}
