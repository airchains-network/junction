package types_test

import (
	"testing"

	"github.com/airchains-network/junction/x/vrf/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated students",
			genState: &types.GenesisState{
				StudentsList: []types.Students{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid students count",
			genState: &types.GenesisState{
				StudentsList: []types.Students{
					{
						Id: 1,
					},
				},
				StudentsCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
