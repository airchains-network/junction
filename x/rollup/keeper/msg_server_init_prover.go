package keeper

import (
	"context"
	"encoding/json"

	"cosmossdk.io/store/prefix"
	bls12381 "github.com/airchains-network/gnark/backend/groth16/bls12-381"
	"github.com/airchains-network/junction/x/rollup/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// InitProver initializes a prover for a specific rollup. It checks if the rollup exists, validates the provided verification key,
// and ensures that no prover details already exist before storing the new prover information in the database.

func (k msgServer) InitProver(goCtx context.Context, msg *types.MsgInitProver) (*types.MsgInitProverResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	rollupDataStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RollupDataKey))

	var rollupId = msg.RollupId
	var proverVerificationKey = msg.ProverVerificationKey
	var proverType = msg.ProverType
	var proverEndpoint = msg.ProverEndpoint

	// Check if the rollup exists in the data store.
	if !rollupDataStore.Has([]byte(rollupId)) {
		return nil, status.Error(codes.NotFound, "rollup not found")
	}

	// Check if the verification key is valid. This is crucial because an invalid verification key could lead to security vulnerabilities
	// or incorrect behavior in the prover's operations. We unmarshal the provided verification key into a VerifyingKey structure.
	var vk bls12381.VerifyingKey

	err := json.Unmarshal(proverVerificationKey, &vk)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid verification key")
	}

	var rollup types.RollupMetadata
	rollupBytes := rollupDataStore.Get([]byte(rollupId))
	k.cdc.MustUnmarshal(rollupBytes, &rollup)
	// Prevent overwriting existing prover information by checking if any prover details are already set.
	if rollup.ProverType != "none" || rollup.ProverVerificationKey != nil || rollup.ProverEndpoint != "none" {
		return nil, status.Error(codes.AlreadyExists, "prover details already exist")
	}

	// Store the new prover details in the rollup metadata.
	rollup.ProverType = proverType
	rollup.ProverVerificationKey = proverVerificationKey
	rollup.ProverEndpoint = proverEndpoint

	rollupBytes = k.cdc.MustMarshal(&rollup)
	rollupDataStore.Set([]byte(rollupId), rollupBytes)

	// Emit an event indicating that the prover has been successfully initialized.
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		"rollup-prover-initialized",
		sdk.NewAttribute("rollup-id", rollupId),
		sdk.NewAttribute("rollup-moniker", rollup.Moniker),
		sdk.NewAttribute("prover-type", proverType),
		sdk.NewAttribute("prover-endpoint", proverEndpoint),
	))

	return &types.MsgInitProverResponse{
		Status: true,
	}, nil
}
