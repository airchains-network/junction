package app

import (
	storetypes "cosmossdk.io/store/types"
	"fmt"
	"github.com/airchains-network/junction/x/wasm"
	wasmkeeper "github.com/airchains-network/junction/x/wasm/keeper"
	wasmtypes "github.com/airchains-network/junction/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/runtime"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/spf13/cast"
	"path/filepath"
	"strings"
)

// AllCapabilities returns all capabilities available with the current wasmvm
// See https://github.com/CosmWasm/cosmwasm/blob/main/docs/CAPABILITIES-BUILT-IN.md
// This functionality is going to be moved upstream: https://github.com/CosmWasm/wasmvm/issues/425
func AllCapabilities() []string {
	return []string{
		"iterator",
		"staking",
		"stargate",
		"cosmwasm_1_1",
		"cosmwasm_1_2",
		"cosmwasm_1_3",
		"cosmwasm_1_4",
	}
}

// registerWasmModules registers the wasm modules with the given app options and wasm options.
func (app *App) registerWasmModules(appOpts servertypes.AppOptions, wasmOpts []wasmkeeper.Option) {
	// set up non depinject support modules store keys
	app.ParamsKeeper.Subspace(wasmtypes.ModuleName)

	// Create and register KVStoreKey
	wasmStoreKey := storetypes.NewKVStoreKey(wasmtypes.StoreKey)
	app.MountStore(wasmStoreKey, storetypes.StoreTypeIAVL)

	scopedWasmKeeper := app.CapabilityKeeper.ScopeToModule(wasmtypes.ModuleName)

	// Create Wasm Keeper
	homePath := cast.ToString(appOpts.Get(flags.FlagHome))

	wasmDir := filepath.Join(homePath, "wasm")
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic(fmt.Sprintf("error while reading wasm config: %s", err))
	}

	availableCapabilities := strings.Join(AllCapabilities(), ",")
	app.WasmKeeper = wasmkeeper.NewKeeper(
		app.appCodec,
		runtime.NewKVStoreService(wasmStoreKey),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		distrkeeper.NewQuerier(app.DistrKeeper),
		app.IBCFeeKeeper, // ISC4 Wrapper: fee IBC middleware
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.PortKeeper,
		scopedWasmKeeper,
		app.TransferKeeper,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		availableCapabilities,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		wasmOpts...,
	)

	app.ScopedWasmKeeper = scopedWasmKeeper
}
