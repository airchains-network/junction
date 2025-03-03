package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"encoding/binary"
	"fmt"
	"github.com/airchains-network/junction/x/vrf/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/big"
)

func (k msgServer) InitiateVrf(goCtx context.Context, msg *types.MsgInitiateVrf) (*types.MsgInitiateVrfResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	upperBond := msg.UpperBound
	// Connect to collection store
	collectionStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey))
	// Connect to the VRF store for this collection.
	vrfStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey+"/"+msg.CollectionId))

	collectionKey := []byte(msg.CollectionId)

	// Check if the collection exists.
	collectionBytes := collectionStore.Get(collectionKey)
	if collectionBytes == nil {
		errMsg := fmt.Sprintf("Collection '%s' does not exist", msg.CollectionId)
		return &types.MsgInitiateVrfResponse{Result: false}, status.Error(codes.NotFound, errMsg)
	}

	// Unmarshal the collection details.
	var collection types.Collection
	k.cdc.MustUnmarshal(collectionBytes, &collection)

	// check if its grey listed. e.g. voting dispute initiated
	if collection.GreyListed == true {
		return &types.MsgInitiateVrfResponse{Result: false}, status.Error(codes.Unavailable, "collection is grey listed")
	}

	// Check if the sender is a member of the collection.
	isMember := false
	for _, member := range collection.Members {
		if member == msg.Creator {
			isMember = true
			break
		}
	}
	if !isMember {
		errMsg := fmt.Sprintf("Sender '%s' is not a member of collection '%s'.", msg.Creator, msg.CollectionId)
		return &types.MsgInitiateVrfResponse{Result: false}, status.Error(codes.PermissionDenied, errMsg)
	}

	// Convert the key to a 4-byte representation.
	vrfKey := msg.Key
	if vrfKey > 4294967294 {
		errMsg := fmt.Sprintf("VRF key %d exceeds the maximum allowed value %d; please create a new collection", vrfKey, 4294967295)
		return &types.MsgInitiateVrfResponse{Result: false}, status.Error(codes.ResourceExhausted, errMsg)
	}
	vrfKeyBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(vrfKeyBytes, vrfKey)

	// Check that the key is in sequence.
	expectedKey := collection.Index + 1
	if vrfKey < expectedKey {
		errMsg := fmt.Sprintf("Invalid VRF key id: passed %d, expected %d. This key has already been initiated.", vrfKey, expectedKey)
		return &types.MsgInitiateVrfResponse{Result: false}, status.Error(codes.InvalidArgument, errMsg)
	}
	if vrfKey > expectedKey {
		errMsg := fmt.Sprintf("Invalid VRF key id: passed %d, expected %d. Initiate preceding keys first.", vrfKey, expectedKey)
		return &types.MsgInitiateVrfResponse{Result: false}, status.Error(codes.InvalidArgument, errMsg)
	}

	// Ensure the VRF record doesn't already exist.
	if existingRecord := vrfStore.Get(vrfKeyBytes); existingRecord != nil {
		errMsg := fmt.Sprintf("VRF record for key %d already exists", vrfKey)
		return &types.MsgInitiateVrfResponse{Result: false}, status.Error(codes.AlreadyExists, errMsg)
	}

	// Generate deterministic random number from the provided proof.
	vrn, vrnErr := GenerateDeterministicRandomNumber(msg.Proof)
	if vrnErr != nil {
		return &types.MsgInitiateVrfResponse{Result: false}, status.Error(codes.InvalidArgument, "error generating deterministic random number")
	}

	vrnBigInt := new(big.Int).SetBytes(vrn)

	upperBondBigInt := new(big.Int).SetUint64(upperBond)
	wrappedIndex := new(big.Int).Mod(vrnBigInt, upperBondBigInt).Uint64()

	// Construct the VRF record.
	vrfRecord := types.VrfRecord{
		RecordSubmitter:    msg.Creator,
		RecordVerifier:     "",
		KeyIndex:           msg.Key,
		CollectionId:       msg.CollectionId,
		UpperBound:         msg.UpperBound,
		VrfPubkey:          msg.VrfPubkey,
		Proof:              msg.Proof,
		OriginDigest:       msg.OriginDigest,
		VerificationStatus: false,
		Vrn:                vrn,
		WrappedIndex:       wrappedIndex,
		SerializedRc:       msg.SerializedRc,
	}

	vrfRecordBytes := k.cdc.MustMarshal(&vrfRecord)
	vrfStore.Set(vrfKeyBytes, vrfRecordBytes)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent("vrf-initiated",
			sdk.NewAttribute("collection-id", msg.CollectionId),
			sdk.NewAttribute("key", fmt.Sprintf("%d", msg.Key)),
			sdk.NewAttribute("submitter", msg.Creator),
			sdk.NewAttribute("wrapped-index", fmt.Sprintf("%d", wrappedIndex)),
		),
	)

	return &types.MsgInitiateVrfResponse{Result: true}, nil

}
