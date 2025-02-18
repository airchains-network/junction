package zksequencer_test

import (
	"testing"

	keepertest "github.com/airchains-network/junction/testutil/keeper"
	"github.com/airchains-network/junction/testutil/nullify"
	zksequencer "github.com/airchains-network/junction/x/zksequencer/module"
	"github.com/airchains-network/junction/x/zksequencer/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		CollegeList: []types.College{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CollegeCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ZksequencerKeeper(t)
	zksequencer.InitGenesis(ctx, k, genesisState)
	got := zksequencer.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	portID := k.GetPort(ctx)
	require.Equal(t, genesisState.PortId, portID)

	colleges := k.GetAllCollege(ctx)
	require.ElementsMatch(t, genesisState.CollegeList, colleges)

	count := k.GetCollegeCount(ctx)
	require.Equal(t, genesisState.CollegeCount, count)
	// this line is used by starport scaffolding # genesis/test/assert
}
