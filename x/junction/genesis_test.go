package junction_test

import (
	"testing"

	keepertest "github.com/ComputerKeeda/junction/testutil/keeper"
	"github.com/ComputerKeeda/junction/testutil/nullify"
	"github.com/ComputerKeeda/junction/x/junction"
	"github.com/ComputerKeeda/junction/x/junction/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.JunctionKeeper(t)
	junction.InitGenesis(ctx, *k, genesisState)
	got := junction.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
