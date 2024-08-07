package keeper_test

import (
	"github.com/airchains-network/junction/app"
	"os"
	"testing"
	"time"

	wasmvm "github.com/CosmWasm/wasmvm/v2"
	wasmvmtypes "github.com/CosmWasm/wasmvm/v2/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/airchains-network/junction/x/wasm/keeper"
	"github.com/airchains-network/junction/x/wasm/types"
)

func TestSnapshotter(t *testing.T) {
	specs := map[string]struct {
		wasmFiles []string
	}{
		"single contract": {
			wasmFiles: []string{"./testdata/reflect_1_5.wasm"},
		},
		"multiple contract": {
			wasmFiles: []string{"./testdata/reflect_1_5.wasm", "./testdata/burner.wasm", "./testdata/reflect_1_5.wasm"},
		},
		"duplicate contracts": {
			wasmFiles: []string{"./testdata/reflect_1_5.wasm", "./testdata/reflect_1_5.wasm"},
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			// setup source app
			srcjunctionApp, genesisAddr := exampleApp(t)

			// store wasm codes on chain
			ctx := srcjunctionApp.NewUncachedContext(false, tmproto.Header{
				ChainID: "foo",
				Height:  srcjunctionApp.LastBlockHeight() + 1,
				Time:    time.Now(),
			})
			wasmKeeper := srcjunctionApp.WasmKeeper
			contractKeeper := keeper.NewDefaultPermissionKeeper(&wasmKeeper)

			srcCodeIDToChecksum := make(map[uint64][]byte, len(spec.wasmFiles))
			for i, v := range spec.wasmFiles {
				wasmCode, err := os.ReadFile(v)
				require.NoError(t, err)
				codeID, checksum, err := contractKeeper.Create(ctx, genesisAddr, wasmCode, nil)
				require.NoError(t, err)
				require.Equal(t, uint64(i+1), codeID)
				srcCodeIDToChecksum[codeID] = checksum
			}
			// create snapshot
			_, err := srcjunctionApp.Commit()
			require.NoError(t, err)

			snapshotHeight := uint64(srcjunctionApp.LastBlockHeight())
			snapshot, err := srcjunctionApp.SnapshotManager().Create(snapshotHeight)
			require.NoError(t, err)
			assert.NotNil(t, snapshot)

			originalMaxWasmSize := types.MaxWasmSize
			types.MaxWasmSize = 1
			t.Cleanup(func() {
				types.MaxWasmSize = originalMaxWasmSize
			})

			// when snapshot imported into dest app instance
			destjunctionApp := app.SetupWithEmptyStore(t)
			require.NoError(t, destjunctionApp.SnapshotManager().Restore(*snapshot))
			for i := uint32(0); i < snapshot.Chunks; i++ {
				chunkBz, err := srcjunctionApp.SnapshotManager().LoadChunk(snapshot.Height, snapshot.Format, i)
				require.NoError(t, err)
				end, err := destjunctionApp.SnapshotManager().RestoreChunk(chunkBz)
				require.NoError(t, err)
				if end {
					break
				}
			}

			// then all wasm contracts are imported
			wasmKeeper = destjunctionApp.WasmKeeper
			ctx = destjunctionApp.NewUncachedContext(false, tmproto.Header{
				ChainID: "foo",
				Height:  destjunctionApp.LastBlockHeight() + 1,
				Time:    time.Now(),
			})

			destCodeIDToChecksum := make(map[uint64][]byte, len(spec.wasmFiles))
			wasmKeeper.IterateCodeInfos(ctx, func(id uint64, info types.CodeInfo) bool {
				bz, err := wasmKeeper.GetByteCode(ctx, id)
				require.NoError(t, err)

				hash, err := wasmvm.CreateChecksum(bz)
				require.NoError(t, err)
				destCodeIDToChecksum[id] = hash[:]
				assert.Equal(t, hash[:], wasmvmtypes.Checksum(info.CodeHash))
				return false
			})
			assert.Equal(t, srcCodeIDToChecksum, destCodeIDToChecksum)
		})
	}
}

func exampleApp(t *testing.T) (*app.App, sdk.AccAddress) {
	senderPrivKey := ed25519.GenPrivKey()
	pubKey, err := cryptocodec.ToCmtPubKeyInterface(senderPrivKey.PubKey())
	require.NoError(t, err)

	senderAddr := senderPrivKey.PubKey().Address().Bytes()
	acc := authtypes.NewBaseAccount(senderAddr, senderPrivKey.PubKey(), 0, 0)
	amount, ok := sdkmath.NewIntFromString("10000000000000000000")
	require.True(t, ok)

	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, amount)),
	}
	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})
	junctionApp := app.SetupWithGenesisValSet(t, valSet, []authtypes.GenesisAccount{acc}, "testing", nil, balance)

	return junctionApp, senderAddr
}
