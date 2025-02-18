package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:       PortID,
		StudentsList: []Students{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated ID in students
	studentsIdMap := make(map[uint64]bool)
	studentsCount := gs.GetStudentsCount()
	for _, elem := range gs.StudentsList {
		if _, ok := studentsIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for students")
		}
		if elem.Id >= studentsCount {
			return fmt.Errorf("students id should be lower or equal than the last id")
		}
		studentsIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
