package keeper

import (
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	portkeeper "github.com/cosmos/ibc-go/v8/modules/core/05-port/keeper"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
	"github.com/stretchr/testify/require"

	"github.com/airchains-network/junction/x/cipherpodledger/keeper"
	"github.com/airchains-network/junction/x/cipherpodledger/types"
)

func CipherpodledgerKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	appCodec := codec.NewProtoCodec(registry)
	capabilityKeeper := capabilitykeeper.NewKeeper(appCodec, storeKey, memStoreKey)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	scopedKeeper := capabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	portKeeper := portkeeper.NewKeeper(scopedKeeper)
	scopeModule := capabilityKeeper.ScopeToModule(types.ModuleName)

	k := keeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authority.String(),
		func() *ibckeeper.Keeper {
			return &ibckeeper.Keeper{
				PortKeeper: &portKeeper,
			}
		},
		func(string) capabilitykeeper.ScopedKeeper {
			return scopeModule
		},
		nil,
		nil,
		nil,
		nil,
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
