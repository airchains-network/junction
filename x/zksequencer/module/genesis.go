package zksequencer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/airchains-network/junction/x/zksequencer/keeper"
	"github.com/airchains-network/junction/x/zksequencer/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the college
	for _, elem := range genState.CollegeList {
		k.SetCollege(ctx, elem)
	}

	// Set college count
	k.SetCollegeCount(ctx, genState.CollegeCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if k.ShouldBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
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
