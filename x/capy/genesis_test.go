package capy_test

import (
	"testing"

	keepertest "capy/testutil/keeper"
	"capy/testutil/nullify"
	"capy/x/capy"
	"capy/x/capy/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AnimalsList: []types.Animals{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AnimalsCount: 2,
		AnimalSkillList: []types.AnimalSkill{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AnimalSkillCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CapyKeeper(t)
	capy.InitGenesis(ctx, *k, genesisState)
	got := capy.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AnimalsList, got.AnimalsList)
	require.Equal(t, genesisState.AnimalsCount, got.AnimalsCount)
	require.ElementsMatch(t, genesisState.AnimalSkillList, got.AnimalSkillList)
	require.Equal(t, genesisState.AnimalSkillCount, got.AnimalSkillCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
