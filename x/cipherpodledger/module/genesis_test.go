package cipherpodledger_test

import (
	"testing"

	keepertest "github.com/airchains-network/junction/testutil/keeper"
	"github.com/airchains-network/junction/testutil/nullify"
	cipherpodledger "github.com/airchains-network/junction/x/cipherpodledger/module"
	"github.com/airchains-network/junction/x/cipherpodledger/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CipherpodledgerKeeper(t)
	cipherpodledger.InitGenesis(ctx, k, genesisState)
	got := cipherpodledger.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
