package trackgate

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/airchains-network/junction/testutil/sample"
	trackgatesimulation "github.com/airchains-network/junction/x/trackgate/simulation"
	"github.com/airchains-network/junction/x/trackgate/types"
)

// avoid unused import issue
var (
	_ = trackgatesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgInitStation = "op_weight_msg_init_station"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitStation int = 100

	opWeightMsgSchemaCreation = "op_weight_msg_schema_creation"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSchemaCreation int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	trackgateGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&trackgateGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgInitStation int
	simState.AppParams.GetOrGenerate(opWeightMsgInitStation, &weightMsgInitStation, nil,
		func(_ *rand.Rand) {
			weightMsgInitStation = defaultWeightMsgInitStation
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInitStation,
		trackgatesimulation.SimulateMsgInitStation(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSchemaCreation int
	simState.AppParams.GetOrGenerate(opWeightMsgSchemaCreation, &weightMsgSchemaCreation, nil,
		func(_ *rand.Rand) {
			weightMsgSchemaCreation = defaultWeightMsgSchemaCreation
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSchemaCreation,
		trackgatesimulation.SimulateMsgSchemaCreation(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgInitStation,
			defaultWeightMsgInitStation,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				trackgatesimulation.SimulateMsgInitStation(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSchemaCreation,
			defaultWeightMsgSchemaCreation,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				trackgatesimulation.SimulateMsgSchemaCreation(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
