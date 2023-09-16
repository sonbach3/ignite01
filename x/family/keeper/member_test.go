package keeper_test

import (
	"testing"

	keepertest "capy/testutil/keeper"
	"capy/testutil/nullify"
	"capy/x/family/keeper"
	"capy/x/family/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNMember(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Member {
	items := make([]types.Member, n)
	for i := range items {
		items[i].Id = keeper.AppendMember(ctx, items[i])
	}
	return items
}

func TestMemberGet(t *testing.T) {
	keeper, ctx := keepertest.FamilyKeeper(t)
	items := createNMember(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMember(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMemberRemove(t *testing.T) {
	keeper, ctx := keepertest.FamilyKeeper(t)
	items := createNMember(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMember(ctx, item.Id)
		_, found := keeper.GetMember(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMemberGetAll(t *testing.T) {
	keeper, ctx := keepertest.FamilyKeeper(t)
	items := createNMember(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMember(ctx)),
	)
}

func TestMemberCount(t *testing.T) {
	keeper, ctx := keepertest.FamilyKeeper(t)
	items := createNMember(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMemberCount(ctx))
}
