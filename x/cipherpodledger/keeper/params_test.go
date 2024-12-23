package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/airchains-network/junction/testutil/keeper"
	"github.com/airchains-network/junction/x/cipherpodledger/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CipherpodledgerKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
