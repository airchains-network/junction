package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/airchains-network/junction/testutil/keeper"
	"github.com/airchains-network/junction/testutil/nullify"
	"github.com/airchains-network/junction/x/vrf/keeper"
	"github.com/airchains-network/junction/x/vrf/types"

	"github.com/stretchr/testify/require"
)

func createNStudents(keeper keeper.Keeper, ctx context.Context, n int) []types.Students {
	items := make([]types.Students, n)
	for i := range items {
		items[i].Id = keeper.AppendStudents(ctx, items[i])
	}
	return items
}

func TestStudentsGet(t *testing.T) {
	keeper, ctx := keepertest.VrfKeeper(t)
	items := createNStudents(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetStudents(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestStudentsRemove(t *testing.T) {
	keeper, ctx := keepertest.VrfKeeper(t)
	items := createNStudents(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStudents(ctx, item.Id)
		_, found := keeper.GetStudents(ctx, item.Id)
		require.False(t, found)
	}
}

func TestStudentsGetAll(t *testing.T) {
	keeper, ctx := keepertest.VrfKeeper(t)
	items := createNStudents(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStudents(ctx)),
	)
}

func TestStudentsCount(t *testing.T) {
	keeper, ctx := keepertest.VrfKeeper(t)
	items := createNStudents(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetStudentsCount(ctx))
}
