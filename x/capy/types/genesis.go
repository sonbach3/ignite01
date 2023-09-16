package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AnimalsList:     []Animals{},
		AnimalSkillList: []AnimalSkill{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in animals
	animalsIdMap := make(map[uint64]bool)
	animalsCount := gs.GetAnimalsCount()
	for _, elem := range gs.AnimalsList {
		if _, ok := animalsIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for animals")
		}
		if elem.Id >= animalsCount {
			return fmt.Errorf("animals id should be lower or equal than the last id")
		}
		animalsIdMap[elem.Id] = true
	}
	// Check for duplicated ID in animalSkill
	animalSkillIdMap := make(map[uint64]bool)
	animalSkillCount := gs.GetAnimalSkillCount()
	for _, elem := range gs.AnimalSkillList {
		if _, ok := animalSkillIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for animalSkill")
		}
		if elem.Id >= animalSkillCount {
			return fmt.Errorf("animalSkill id should be lower or equal than the last id")
		}
		animalSkillIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
