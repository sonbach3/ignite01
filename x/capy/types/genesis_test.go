package types_test

import (
	"testing"

	"capy/x/capy/types"
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated animals",
			genState: &types.GenesisState{
				AnimalsList: []types.Animals{
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
			desc: "invalid animals count",
			genState: &types.GenesisState{
				AnimalsList: []types.Animals{
					{
						Id: 1,
					},
				},
				AnimalsCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated animalSkill",
			genState: &types.GenesisState{
				AnimalSkillList: []types.AnimalSkill{
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
			desc: "invalid animalSkill count",
			genState: &types.GenesisState{
				AnimalSkillList: []types.AnimalSkill{
					{
						Id: 1,
					},
				},
				AnimalSkillCount: 0,
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
