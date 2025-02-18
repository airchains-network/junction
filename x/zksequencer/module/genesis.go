package zksequencer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/airchains-network/junction/x/zksequencer/keeper"
	"github.com/airchains-network/junction/x/zksequencer/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set params first as this is typically required for other operations
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	// Set port before college operations
	k.SetPort(ctx, genState.PortId)

	// Only try to bind to port if it is not already bound
	if k.ShouldBound(ctx, genState.PortId) {
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}

	// Set college count before setting individual colleges
	k.SetCollegeCount(ctx, genState.CollegeCount)

	// Set all the colleges
	for _, elem := range genState.CollegeList {
		k.SetCollege(ctx, elem)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.CollegeList = k.GetAllCollege(ctx)
	genesis.CollegeCount = k.GetCollegeCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
