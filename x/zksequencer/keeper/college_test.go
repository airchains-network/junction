package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/airchains-network/junction/testutil/keeper"
	"github.com/airchains-network/junction/testutil/nullify"
	"github.com/airchains-network/junction/x/zksequencer/keeper"
	"github.com/airchains-network/junction/x/zksequencer/types"

	"github.com/stretchr/testify/require"
)

func createNCollege(keeper keeper.Keeper, ctx context.Context, n int) []types.College {
	items := make([]types.College, n)
	for i := range items {
		items[i].Id = keeper.AppendCollege(ctx, items[i])
	}
	return items
}

func TestCollegeGet(t *testing.T) {
	keeper, ctx := keepertest.ZksequencerKeeper(t)
	items := createNCollege(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetCollege(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestCollegeRemove(t *testing.T) {
	keeper, ctx := keepertest.ZksequencerKeeper(t)
	items := createNCollege(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCollege(ctx, item.Id)
		_, found := keeper.GetCollege(ctx, item.Id)
		require.False(t, found)
	}
}

func TestCollegeGetAll(t *testing.T) {
	keeper, ctx := keepertest.ZksequencerKeeper(t)
	items := createNCollege(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCollege(ctx)),
	)
}

func TestCollegeCount(t *testing.T) {
	keeper, ctx := keepertest.ZksequencerKeeper(t)
	items := createNCollege(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetCollegeCount(ctx))
}
