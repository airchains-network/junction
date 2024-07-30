package app

import (
	"cosmossdk.io/store/snapshots"
	snapshottypes "cosmossdk.io/store/snapshots/types"
	"encoding/json"
	"fmt"
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/server"
	"math/rand"
	"path/filepath"
	"testing"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	cmttypes "github.com/cometbft/cometbft/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	wasmkeeper "github.com/airchains-network/junction/x/wasm/keeper"
)

// SetupOptions defines arguments that are passed into `WasmApp` constructor.
type SetupOptions struct {
	Logger   log.Logger
	DB       *dbm.MemDB
	AppOpts  servertypes.AppOptions
	WasmOpts []wasmkeeper.Option
}

func setup(t testing.TB, chainID string, withGenesis bool, invCheckPeriod uint, opts ...wasmkeeper.Option) (*App, GenesisState) {
	db := dbm.NewMemDB()
	nodeHome := t.TempDir()
	snapshotDir := filepath.Join(nodeHome, "data", "snapshots")

	snapshotDB, err := dbm.NewDB("metadata", dbm.GoLevelDBBackend, snapshotDir)
	require.NoError(t, err)
	t.Cleanup(func() { snapshotDB.Close() })
	snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	require.NoError(t, err)

	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = nodeHome // ensure unique folder
	appOptions[server.FlagInvCheckPeriod] = invCheckPeriod
	app, err := New(log.NewNopLogger(), db, nil, true, appOptions, opts, bam.SetChainID(chainID), bam.SetSnapshot(snapshotStore, snapshottypes.SnapshotOptions{KeepRecent: 2}))
	require.NoError(t, err)
	if withGenesis {
		return app, app.DefaultGenesis()
	}
	return app, GenesisState{}
}

// SetupWithGenesisValSet initializes a new WasmApp with a validator set and genesis accounts
// that also act as delegators. For simplicity, each validator is bonded with a delegation
// of one consensus engine unit in the default token of the WasmApp from first genesis
// account. A Nop logger is set in WasmApp.
func SetupWithGenesisValSet(
	t *testing.T,
	valSet *cmttypes.ValidatorSet,
	genAccs []authtypes.GenesisAccount,
	chainID string,
	opts []wasmkeeper.Option,
	balances ...banktypes.Balance,
) *App {
	t.Helper()

	app, genesisState := setup(t, chainID, true, 5, opts...)
	genesisState, err := GenesisStateWithValSet(app.AppCodec(), genesisState, valSet, genAccs, balances...)
	require.NoError(t, err)

	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	require.NoError(t, err)

	// init chain will set the validator set and initialize the genesis accounts
	consensusParams := simtestutil.DefaultConsensusParams
	consensusParams.Block.MaxGas = 100 * simtestutil.DefaultGenTxGas
	_, err = app.InitChain(&abci.RequestInitChain{
		ChainId:         chainID,
		Time:            time.Now().UTC(),
		Validators:      []abci.ValidatorUpdate{},
		ConsensusParams: consensusParams,
		InitialHeight:   app.LastBlockHeight() + 1,
		AppStateBytes:   stateBytes,
	})
	require.NoError(t, err)

	_, err = app.FinalizeBlock(&abci.RequestFinalizeBlock{
		Height:             app.LastBlockHeight() + 1,
		Hash:               app.LastCommitID().Hash,
		NextValidatorsHash: valSet.Hash(),
	})
	require.NoError(t, err)

	return app
}

// SetupWithEmptyStore set up a wasmd app instance with empty DB
func SetupWithEmptyStore(t testing.TB) *App {
	app, _ := setup(t, "testing", false, 0)
	return app
}

// SignAndDeliverWithoutCommit signs and delivers a transaction. No commit
func SignAndDeliverWithoutCommit(t *testing.T, txCfg client.TxConfig, app *bam.BaseApp, msgs []sdk.Msg, fees sdk.Coins, chainID string, accNums, accSeqs []uint64, blockTime time.Time, priv ...cryptotypes.PrivKey) (*abci.ResponseFinalizeBlock, error) {
	tx, err := simtestutil.GenSignedMockTx(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		txCfg,
		msgs,
		fees,
		simtestutil.DefaultGenTxGas,
		chainID,
		accNums,
		accSeqs,
		priv...,
	)
	require.NoError(t, err)

	bz, err := txCfg.TxEncoder()(tx)
	require.NoError(t, err)

	return app.FinalizeBlock(&abci.RequestFinalizeBlock{
		Height: app.LastBlockHeight() + 1,
		Time:   blockTime,
		Txs:    [][]byte{bz},
	})
}

// GenesisStateWithValSet returns a new genesis state with the validator set
// copied from simtestutil with delegation not added to supply
func GenesisStateWithValSet(
	codec codec.Codec,
	genesisState map[string]json.RawMessage,
	valSet *cmttypes.ValidatorSet,
	genAccs []authtypes.GenesisAccount,
	balances ...banktypes.Balance,
) (map[string]json.RawMessage, error) {
	// set genesis accounts
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	genesisState[authtypes.ModuleName] = codec.MustMarshalJSON(authGenesis)

	validators := make([]stakingtypes.Validator, 0, len(valSet.Validators))
	delegations := make([]stakingtypes.Delegation, 0, len(valSet.Validators))

	bondAmt := sdk.DefaultPowerReduction

	for _, val := range valSet.Validators {
		pk, err := cryptocodec.FromCmtPubKeyInterface(val.PubKey)
		if err != nil {
			return nil, fmt.Errorf("failed to convert pubkey: %w", err)
		}

		pkAny, err := codectypes.NewAnyWithValue(pk)
		if err != nil {
			return nil, fmt.Errorf("failed to create new any: %w", err)
		}

		validator := stakingtypes.Validator{
			OperatorAddress:   sdk.ValAddress(val.Address).String(),
			ConsensusPubkey:   pkAny,
			Jailed:            false,
			Status:            stakingtypes.Bonded,
			Tokens:            bondAmt,
			DelegatorShares:   sdkmath.LegacyOneDec(),
			Description:       stakingtypes.Description{},
			UnbondingHeight:   int64(0),
			UnbondingTime:     time.Unix(0, 0).UTC(),
			Commission:        stakingtypes.NewCommission(sdkmath.LegacyZeroDec(), sdkmath.LegacyZeroDec(), sdkmath.LegacyZeroDec()),
			MinSelfDelegation: sdkmath.ZeroInt(),
		}
		validators = append(validators, validator)
		delegations = append(delegations, stakingtypes.NewDelegation(genAccs[0].GetAddress().String(), sdk.ValAddress(val.Address).String(), sdkmath.LegacyOneDec()))

	}

	// set validators and delegations
	stakingGenesis := stakingtypes.NewGenesisState(stakingtypes.DefaultParams(), validators, delegations)
	genesisState[stakingtypes.ModuleName] = codec.MustMarshalJSON(stakingGenesis)

	signingInfos := make([]slashingtypes.SigningInfo, len(valSet.Validators))
	for i, val := range valSet.Validators {
		signingInfos[i] = slashingtypes.SigningInfo{
			Address:              sdk.ConsAddress(val.Address).String(),
			ValidatorSigningInfo: slashingtypes.ValidatorSigningInfo{},
		}
	}
	slashingGenesis := slashingtypes.NewGenesisState(slashingtypes.DefaultParams(), signingInfos, nil)
	genesisState[slashingtypes.ModuleName] = codec.MustMarshalJSON(slashingGenesis)

	// add bonded amount to bonded pool module account
	balances = append(balances, banktypes.Balance{
		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
		Coins:   sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, bondAmt.MulRaw(int64(len(valSet.Validators))))},
	})

	totalSupply := sdk.NewCoins()
	for _, b := range balances {
		// add genesis acc tokens to total supply
		totalSupply = totalSupply.Add(b.Coins...)
	}

	// update total supply
	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, totalSupply, []banktypes.Metadata{}, []banktypes.SendEnabled{})
	genesisState[banktypes.ModuleName] = codec.MustMarshalJSON(bankGenesis)
	println(string(genesisState[banktypes.ModuleName]))
	return genesisState, nil
}
