package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"github.com/airchains-network/junction/x/vrf/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidateVrf verifies the VRF proof and ensures that only the latest 100 records remain.
func (k msgServer) ValidateVrf(goCtx context.Context, msg *types.MsgValidateVrf) (*types.MsgValidateVrfResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	verifierAddress := msg.Creator

	// Connect to collection store and VRF store.
	collectionStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey))
	vrfStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey+"/"+msg.CollectionId))

	collectionKey := []byte(msg.CollectionId)

	// Check if the collection exists.
	collectionBytes := collectionStore.Get(collectionKey)
	if collectionBytes == nil {
		return nil, status.Errorf(codes.NotFound, "collection '%s' does not exist", msg.CollectionId)
	}

	// Unmarshal the collection details.
	var collection types.Collection
	k.cdc.MustUnmarshal(collectionBytes, &collection)

	// check if its grey listed
	if collection.GreyListed == true {
		return nil, status.Errorf(codes.InvalidArgument, "collection '%s' is grey listed", msg.CollectionId)
	}

	// Check if the sender is a member of the collection.
	if !isCollectionMember(collection.Members, msg.Creator) {
		return nil, status.Errorf(codes.PermissionDenied, "sender '%s' is not a member of collection '%s'.", msg.Creator, msg.CollectionId)
	}

	// Create the key for the VRF record.
	vrfKeyBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(vrfKeyBytes, msg.Key)

	// Ensure the VRF record exists.
	vrfRecord := vrfStore.Get(vrfKeyBytes)
	if vrfRecord == nil {
		return nil, status.Errorf(codes.NotFound, "key %d does not exist", msg.Key)
	}

	// Unmarshal the VRF record.
	var record types.VrfRecord
	k.cdc.MustUnmarshal(vrfRecord, &record)
	if record.VerificationStatus {
		return nil, status.Errorf(codes.AlreadyExists, "record for key %d is already verified", msg.Key)
	}

	if len(collection.Members) > 1 {
		// If there are multiple members, the initiator cannot verify.
		vrfRequesterAddr := record.RecordSubmitter
		if vrfRequesterAddr == verifierAddress {
			return nil, status.Errorf(codes.PermissionDenied, "initiator cannot verify")
		}
	}

	// Verify VRF proof.
	isValid, err := VerifyVRFProof(string(record.VrfPubkey), msg.SerializedRc, record.Proof, record.OriginDigest)
	if err != nil || isValid == false {

		// saving accused submitter and verifier
		record.RecordVerifier = msg.Creator
		vrfRecordBytes := k.cdc.MustMarshal(&record)
		vrfStore.Set(vrfKeyBytes, vrfRecordBytes)

		// save grey listed collection data
		collection.GreyListed = true
		blockHeight := ctx.BlockHeight()
		disputeYesHash := GetDisputeVoteHash(true, blockHeight, msg.Key, record.RecordSubmitter, msg.Creator)
		disputeNoHash := GetDisputeVoteHash(false, blockHeight, msg.Key, record.RecordSubmitter, msg.Creator)
		collection.DisputeVoteYesHash = disputeYesHash
		collection.DisputeVoteNoHash = disputeNoHash
		ctx.EventManager().EmitEvent(
			sdk.NewEvent("dispute-found",
				sdk.NewAttribute("collection-id", msg.CollectionId),
				sdk.NewAttribute("key", fmt.Sprintf("%d", msg.Key)),
				sdk.NewAttribute("submitter", record.RecordSubmitter),
				sdk.NewAttribute("verifier", msg.Creator),
				sdk.NewAttribute("dispute-yes-hash", disputeYesHash),
				sdk.NewAttribute("dispute-no-hash", disputeNoHash),
			),
		)
		collectionBytes = k.cdc.MustMarshal(&collection)
		collectionStore.Set(collectionKey, collectionBytes)

		return &types.MsgValidateVrfResponse{Result: false}, nil
	}

	// Mark the record as verified and save it.
	record.VerificationStatus = true
	record.RecordVerifier = verifierAddress
	updatedRecordBytes := k.cdc.MustMarshal(&record)
	vrfStore.Set(vrfKeyBytes, updatedRecordBytes)

	// Clean up old records, ensuring only the latest 100 remain.
	k.cleanupOldVrfRecords(vrfStore, &collection, collectionKey, collectionStore, msg.Key)

	// update collection
	collection.Index += 1
	collectionBytes = k.cdc.MustMarshal(&collection)
	collectionStore.Set(collectionKey, collectionBytes)

	return &types.MsgValidateVrfResponse{Result: true}, nil

}

func GetDisputeVoteHash(voteOption bool, blockHeight int64, key uint32, recordSubmitter string, recordVerifier string) string {
	hash := sha256.Sum256([]byte(fmt.Sprintf("%t%d%d%s%s", voteOption, blockHeight, key, recordSubmitter, recordVerifier)))
	return fmt.Sprintf("%x", hash)
}
