package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MemberList: []Member{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in member
	memberIdMap := make(map[uint64]bool)
	memberCount := gs.GetMemberCount()
	for _, elem := range gs.MemberList {
		if _, ok := memberIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for member")
		}
		if elem.Id >= memberCount {
			return fmt.Errorf("member id should be lower or equal than the last id")
		}
		memberIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
