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

func createNAnimalSkill(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AnimalSkill {
	items := make([]types.AnimalSkill, n)
	for i := range items {
		items[i].Id = keeper.AppendAnimalSkill(ctx, items[i])
	}
	return items
}

func TestAnimalSkillGet(t *testing.T) {
	keeper, ctx := keepertest.CapyKeeper(t)
	items := createNAnimalSkill(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAnimalSkill(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAnimalSkillRemove(t *testing.T) {
	keeper, ctx := keepertest.CapyKeeper(t)
	items := createNAnimalSkill(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAnimalSkill(ctx, item.Id)
		_, found := keeper.GetAnimalSkill(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAnimalSkillGetAll(t *testing.T) {
	keeper, ctx := keepertest.CapyKeeper(t)
	items := createNAnimalSkill(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAnimalSkill(ctx)),
	)
}

func TestAnimalSkillCount(t *testing.T) {
	keeper, ctx := keepertest.CapyKeeper(t)
	items := createNAnimalSkill(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAnimalSkillCount(ctx))
}
