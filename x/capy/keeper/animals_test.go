package keeper_test

import (
	"testing"

	keepertest "capy/testutil/keeper"
	"capy/testutil/nullify"
	"capy/x/capy/keeper"
	"capy/x/capy/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNAnimals(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Animals {
	items := make([]types.Animals, n)
	for i := range items {
		items[i].Id = keeper.AppendAnimals(ctx, items[i])
	}
	return items
}

func TestAnimalsGet(t *testing.T) {
	keeper, ctx := keepertest.CapyKeeper(t)
	items := createNAnimals(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAnimals(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAnimalsRemove(t *testing.T) {
	keeper, ctx := keepertest.CapyKeeper(t)
	items := createNAnimals(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAnimals(ctx, item.Id)
		_, found := keeper.GetAnimals(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAnimalsGetAll(t *testing.T) {
	keeper, ctx := keepertest.CapyKeeper(t)
	items := createNAnimals(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAnimals(ctx)),
	)
}

func TestAnimalsCount(t *testing.T) {
	keeper, ctx := keepertest.CapyKeeper(t)
	items := createNAnimals(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAnimalsCount(ctx))
}
