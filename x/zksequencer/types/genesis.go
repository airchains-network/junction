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
		PortId:      PortID,
		CollegeList: []College{},
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
	// Check for duplicated ID in college
	collegeIdMap := make(map[uint64]bool)
	collegeCount := gs.GetCollegeCount()
	for _, elem := range gs.CollegeList {
		if _, ok := collegeIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for college")
		}
		if elem.Id >= collegeCount {
			return fmt.Errorf("college id should be lower or equal than the last id")
		}
		collegeIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
