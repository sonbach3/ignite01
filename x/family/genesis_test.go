package family_test

import (
	"testing"

	keepertest "capy/testutil/keeper"
	"capy/testutil/nullify"
	"capy/x/family"
	"capy/x/family/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MemberList: []types.Member{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		MemberCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FamilyKeeper(t)
	family.InitGenesis(ctx, *k, genesisState)
	got := family.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.MemberList, got.MemberList)
	require.Equal(t, genesisState.MemberCount, got.MemberCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
