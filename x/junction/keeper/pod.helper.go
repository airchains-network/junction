package keeper

import (
	"encoding/json"
	"github.com/ComputerKeeda/junction/x/junction/types"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
)

func GetPodKeyByte(stationId string, podNumber uint64) (string, []byte) {
	podStoreKey := "pods/" + stationId
	podNumberString := strconv.FormatUint(podNumber, 10)
	newPodStoreKey := podStoreKey + "/" + podNumberString
	podStoreKeyByte := []byte(newPodStoreKey)
	return podStoreKey, podStoreKeyByte
}

func (k Keeper) SubmitPodHelper(ctx sdk.Context, msg *types.MsgSubmitPod) *sdkerrors.Error {

	var stationId = msg.StationId
	var podNumber = msg.PodNumber
	var merkleRootHash = msg.MerkleRootHash
	var previousMerkleRootHash = msg.PreviousMerkleRootHash
	var publicWitness = msg.PublicWitness
	var timestamp = msg.Timestamp

	// check if witness format is correct
	var witness fr.Vector
	witnessCheck := json.Unmarshal(publicWitness, &witness)
	if witnessCheck != nil {
		return sdkerrors.ErrInvalidRequest
	}

	podStoreKey, podStoreKeyByte := GetPodKeyByte(stationId, podNumber) // "pods/{stationId}/{podNumber}
	podStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(podStoreKey))

	// check whether the previous merkle root hash and the incoming hash is same or not
	if podNumber > 1 {
		_, previousPodNumber := GetPodKeyByte(stationId, podNumber-1)
		previousPodDetailsByte := podStore.Get(previousPodNumber)
		var previousPodDetails types.Pods
		k.cdc.MustUnmarshal(previousPodDetailsByte, &previousPodDetails)

		if previousPodDetails.MerkleRootHash != previousMerkleRootHash {
			return sdkerrors.ErrInvalidRequest
		}
	}

	newPod := types.Pods{
		PodNumber:              podNumber,
		MerkleRootHash:         merkleRootHash,
		PreviousMerkleRootHash: previousMerkleRootHash,
		ZkProof:                []byte(""),
		Witness:                publicWitness,
		Timestamp:              timestamp,
		IsVerified:             false,
	}

	storingData := k.cdc.MustMarshal(&newPod)
	// "pods/{stationId}/{podNumberByte}
	podStore.Set(podStoreKeyByte, storingData)

	return nil
}
