package vrf

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/airchains-network/junction/testutil/sample"
	vrfsimulation "github.com/airchains-network/junction/x/vrf/simulation"
	"github.com/airchains-network/junction/x/vrf/types"
)

// avoid unused import issue
var (
	_ = vrfsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateStudents = "op_weight_msg_students"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateStudents int = 100

	opWeightMsgUpdateStudents = "op_weight_msg_students"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateStudents int = 100

	opWeightMsgDeleteStudents = "op_weight_msg_students"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteStudents int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	vrfGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		StudentsList: []types.Students{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		StudentsCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&vrfGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateStudents int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateStudents, &weightMsgCreateStudents, nil,
		func(_ *rand.Rand) {
			weightMsgCreateStudents = defaultWeightMsgCreateStudents
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateStudents,
		vrfsimulation.SimulateMsgCreateStudents(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateStudents int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateStudents, &weightMsgUpdateStudents, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateStudents = defaultWeightMsgUpdateStudents
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateStudents,
		vrfsimulation.SimulateMsgUpdateStudents(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteStudents int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteStudents, &weightMsgDeleteStudents, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteStudents = defaultWeightMsgDeleteStudents
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteStudents,
		vrfsimulation.SimulateMsgDeleteStudents(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateStudents,
			defaultWeightMsgCreateStudents,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vrfsimulation.SimulateMsgCreateStudents(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateStudents,
			defaultWeightMsgUpdateStudents,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vrfsimulation.SimulateMsgUpdateStudents(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteStudents,
			defaultWeightMsgDeleteStudents,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vrfsimulation.SimulateMsgDeleteStudents(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
