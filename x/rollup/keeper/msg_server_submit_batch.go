package keeper

import (
	"context"
	"encoding/json"

	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/rollup/types"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) SubmitBatch(goCtx context.Context, msg *types.MsgSubmitBatch) (*types.MsgSubmitBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	var rollupId = msg.RollupId
	var batchNo = msg.BatchNo
	var merkleRootHash = msg.MerkleRootHash
	var previousMerkleRootHash = msg.PreviousMerkleRootHash
	var zkProof = msg.ZkProof
	var publicWitness = msg.PublicWitness
	var timestamp = ctx.BlockTime().String()

	var witness fr.Vector

	rollup, err := k.getRollupById(ctx, rollupId)
	if err != nil {
		return &types.MsgSubmitBatchResponse{
			Status: false,
		}, err
	}

	vk := rollup.ProverVerificationKey

	if rollup.RollupLatestFinalizedBatchNo+1 != msg.BatchNo {
		return &types.MsgSubmitBatchResponse{
			Status: false,
		}, status.Error(codes.FailedPrecondition, "batch number is incorrect")
	}

	witnessCheck := json.Unmarshal(publicWitness, &witness)
	if witnessCheck != nil {
		return &types.MsgSubmitBatchResponse{
			Status: false,
		}, status.Error(codes.FailedPrecondition, "invalid witness")
	}

	batchStorePath, batchKeyByte := k.GetRollupBatchDbStoreKeys(rollupId, batchNo)
	batchDataStore := prefix.NewStore(storeAdapter, types.KeyPrefix(batchStorePath))

	if batchNo > 1 {
		_, previousBatchKey := k.GetRollupBatchDbStoreKeys(rollupId, batchNo-1)
		previousBatchBytes := batchDataStore.Get(previousBatchKey)
		var previousBatch types.Batch
		k.cdc.MustUnmarshal(previousBatchBytes, &previousBatch)

		if previousBatch.MerkleRootHash != previousMerkleRootHash {
			return &types.MsgSubmitBatchResponse{
				Status: false,
			}, status.Error(codes.FailedPrecondition, "invalid merkleRootHash")
		}
	}

	var batchData types.Batch
	batchByte := batchDataStore.Get(batchKeyByte)
	k.cdc.MustUnmarshal(batchByte, &batchData)

	// Verify Batch
	verified, err := k.VerifyBatch(zkProof, witness, vk)
	if err != nil {
		return &types.MsgSubmitBatchResponse{
			Status: false,
		}, err
	}

	if !verified {
		return &types.MsgSubmitBatchResponse{
			Status: false,
		}, status.Error(codes.Aborted, "verification failed")
	}

	// Update Rollup Metadata and Update Batch Data as well
	// we have rollup and batchData

	batchData.MerkleRootHash = merkleRootHash
	batchData.ZkProof = zkProof
	batchData.PublicWitness = publicWitness
	batchData.Timestamp = timestamp
	batchData.IsFinalized = true

	rollup.RollupLatestFinalizedBatchNo = batchNo

	rollupDataStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RollupDataKey))
	byteRollup := k.cdc.MustMarshal(&rollup)
	byteRollupId := []byte(rollup.RollupId)
	rollupDataStore.Set(byteRollupId, byteRollup)

	batchDataBytes := k.cdc.MustMarshal(&batchData)
	batchDataStore.Set(batchKeyByte, batchDataBytes)

	return &types.MsgSubmitBatchResponse{
		Status: true,
	}, nil
}
