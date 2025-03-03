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
)

func (k msgServer) ProcessVrfDispute(goCtx context.Context, msg *types.MsgProcessVrfDispute) (*types.MsgProcessVrfDisputeResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	var votes = msg.Votes
	// Connect to collection store and VRF store.
	collectionStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey))
	vrfDisputeStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VrfDisputeStoreKey))
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

	// Check if the sender is a member of the collection.
	if !isCollectionMember(collection.Members, msg.Creator) {
		return nil, status.Errorf(codes.PermissionDenied, "sender '%s' is not a member of collection '%s'.", msg.Creator, msg.CollectionId)
	}

	// check if its grey listed
	if collection.GreyListed == false {
		return nil, status.Error(codes.Unavailable, "collection is not grey listed")
	}
	// Create the key for the VRF record.
	vrfKeyBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(vrfKeyBytes, msg.Key)
	// Ensure the VRF record exists.
	vrfRecordByte := vrfStore.Get(vrfKeyBytes)
	if vrfRecordByte == nil {
		return nil, status.Errorf(codes.NotFound, "VRF record for key %d does not exist", msg.Key)
	}
	// Unmarshal the VRF record.
	var vrfRecord types.VrfRecord
	k.cdc.MustUnmarshal(vrfRecordByte, &vrfRecord)
	if vrfRecord.KeyIndex == collection.Offset {
		return nil, status.Errorf(codes.InvalidArgument, "VRF record for key %d is not in dispute", msg.Key)
	}

	// check is all msg.members are valid collection members
	for _, msgMember := range msg.Members {
		isMember := false
		for _, collectionMember := range collection.Members {
			if msgMember == collectionMember {
				isMember = true
			}
		}
		if isMember == false {
			return nil, status.Error(codes.PermissionDenied, "invalid members passed")
		}
	}

	// Special Case : when there is only one member present in the collection
	if len(collection.Members) == 1 {
		// delete vrf record of that key
		vrfStore.Delete(vrfKeyBytes)

		collection.GreyListed = false
		collectionBytes = k.cdc.MustMarshal(&collection)
		collectionStore.Set(collectionKey, collectionBytes)

		return &types.MsgProcessVrfDisputeResponse{Result: true, Message: "Success"}, nil
	} else {

		// verify signatures
		verifiedVotes, signers, success, err := VerifyVrfDisputeSignatures(msg.Members, msg.Signatures, msg.PubKeys, votes, collection)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "signature verification failed: %v", err)
		}
		if !success {
			return nil, status.Error(codes.InvalidArgument, "signature verification was not successful")
		}

		// count votes
		voteResult := countDisputeVotes(verifiedVotes, signers)

		// store voting result
		byteResult := k.cdc.MustMarshal(&voteResult)
		vrfDisputeStore.Set(collectionKey, byteResult)

		collectionMembers := collection.Members
		var newCollectionMembers []string

		if voteResult.Result == true {
			// update collection members, remove record submitter
			RecordSubmitter := vrfRecord.RecordSubmitter
			for _, member := range collectionMembers {
				if member != RecordSubmitter {
					newCollectionMembers = append(newCollectionMembers, member)
				}
			}
			collection.Members = newCollectionMembers

			// delete vrf record of that key
			vrfStore.Delete(vrfKeyBytes)
			ctx.EventManager().EmitEvent(
				sdk.NewEvent("dispute-resolved",
					sdk.NewAttribute("collection-id", msg.CollectionId),
					sdk.NewAttribute("key", fmt.Sprintf("%d", msg.Key)),
					sdk.NewAttribute("culprit", RecordSubmitter),
					sdk.NewAttribute("result", "true"),
				),
			)
		} else {
			// update collection members, remove record verifier
			recordVerifier := vrfRecord.RecordVerifier
			for _, member := range collectionMembers {
				if member != recordVerifier {
					newCollectionMembers = append(newCollectionMembers, member)
				}
			}
			collection.Members = newCollectionMembers
			// update vrf record: set v
			vrfRecord.RecordVerifier = ""
			vrfRecordByte := k.cdc.MustMarshal(&vrfRecord)
			vrfStore.Set(vrfKeyBytes, vrfRecordByte)
			ctx.EventManager().EmitEvent(
				sdk.NewEvent("dispute-resolved",
					sdk.NewAttribute("collection-id", msg.CollectionId),
					sdk.NewAttribute("key", fmt.Sprintf("%d", msg.Key)),
					sdk.NewAttribute("culprit", recordVerifier),
					sdk.NewAttribute("result", "false"),
				),
			)
		}

		// remove collection from grey list
		collection.GreyListed = false
		collection.DisputeVoteYesHash = "none"
		collection.DisputeVoteNoHash = "none"
		collectionBytes = k.cdc.MustMarshal(&collection)
		collectionStore.Set(collectionKey, collectionBytes)

		return &types.MsgProcessVrfDisputeResponse{Result: true, Message: "Success"}, nil
	}
}
