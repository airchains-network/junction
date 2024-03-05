package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ComputerKeeda/junction/testutil/keeper"
	"github.com/ComputerKeeda/junction/x/junction/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.JunctionKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
