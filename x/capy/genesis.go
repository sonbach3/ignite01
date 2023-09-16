package capy

import (
	"capy/x/capy/keeper"
	"capy/x/capy/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the animals
	for _, elem := range genState.AnimalsList {
		k.SetAnimals(ctx, elem)
	}

	// Set animals count
	k.SetAnimalsCount(ctx, genState.AnimalsCount)
	// Set all the animalSkill
	for _, elem := range genState.AnimalSkillList {
		k.SetAnimalSkill(ctx, elem)
	}

	// Set animalSkill count
	k.SetAnimalSkillCount(ctx, genState.AnimalSkillCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AnimalsList = k.GetAllAnimals(ctx)
	genesis.AnimalsCount = k.GetAnimalsCount(ctx)
	genesis.AnimalSkillList = k.GetAllAnimalSkill(ctx)
	genesis.AnimalSkillCount = k.GetAnimalSkillCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
