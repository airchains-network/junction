package keeper

import (
	"fmt"

	"github.com/airchains-network/junction/x/cipherpodledger/types"
)

func (k Keeper) GetPodKeyByte(stationId string, podNumber uint64) (string, []byte) {
	return fmt.Sprintf("%s/%s", types.PodDataStoreKey, stationId), []byte(fmt.Sprintf("%d", podNumber))
}
