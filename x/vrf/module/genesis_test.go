package vrf_test

import (
	"testing"

	keepertest "github.com/airchains-network/junction/testutil/keeper"
	"github.com/airchains-network/junction/testutil/nullify"
	vrf "github.com/airchains-network/junction/x/vrf/module"
	"github.com/airchains-network/junction/x/vrf/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		StudentsList: []types.Students{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		StudentsCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VrfKeeper(t)
	vrf.InitGenesis(ctx, k, genesisState)
	got := vrf.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	portID := k.GetPort(ctx)

	require.Equal(t, genesisState.PortId, portID)

	students := k.GetAllStudents(ctx)
	require.ElementsMatch(t, genesisState.StudentsList, students)

	count := k.GetStudentsCount(ctx)
	require.Equal(t, genesisState.StudentsCount, count)
	// this line is used by starport scaffolding # genesis/test/assert
}
