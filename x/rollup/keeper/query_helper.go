package keeper

import (
	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/rollup/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetMoniker(ctx sdk.Context, moniker string) bool {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	rollupMonikerStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RollupMonikerKey))

	return rollupMonikerStore.Has([]byte(moniker))
}
