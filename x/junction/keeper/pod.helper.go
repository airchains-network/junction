package keeper

import (
	"cosmossdk.io/store/prefix"
	"strconv"
	"github.com/ComputerKeeda/junction/x/junction/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func GetPodKeyByte(stationId string, podNumber uint64) (string, []byte) {
	podStoreKey := "pods/" + stationId
	podNumberString := strconv.FormatUint(podNumber, 10)
	podStoreKeyByte := []byte(podNumberString)
	return podStoreKey, podStoreKeyByte
}

func (k Keeper) GetPodHelper(ctx sdk.Context, stationId string, podNumber uint64) (pods types.Pods, sdkErr error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	podStoreKey, podStoreKeyByte := GetPodKeyByte(stationId, podNumber) // "pods/{stationId}/{podNumber}
	podStore := prefix.NewStore(storeAdapter, types.KeyPrefix(podStoreKey))
	podDetailsByte := podStore.Get(podStoreKeyByte)

	if podDetailsByte == nil {
		return pods, sdkerrors.ErrKeyNotFound
	}

	var podDetails types.Pods
	k.cdc.MustUnmarshal(podDetailsByte, &podDetails)

	return podDetails, nil
}
