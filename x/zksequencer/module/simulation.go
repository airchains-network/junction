package zksequencer

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/airchains-network/junction/testutil/sample"
	zksequencersimulation "github.com/airchains-network/junction/x/zksequencer/simulation"
	"github.com/airchains-network/junction/x/zksequencer/types"
)

// avoid unused import issue
var (
	_ = zksequencersimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateCollege = "op_weight_msg_college"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCollege int = 100

	opWeightMsgUpdateCollege = "op_weight_msg_college"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCollege int = 100

	opWeightMsgDeleteCollege = "op_weight_msg_college"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteCollege int = 100

	opWeightMsgAddStudentViaZksequencer = "op_weight_msg_add_student_via_zksequencer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddStudentViaZksequencer int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	zksequencerGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		CollegeList: []types.College{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		CollegeCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&zksequencerGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateCollege int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateCollege, &weightMsgCreateCollege, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCollege = defaultWeightMsgCreateCollege
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCollege,
		zksequencersimulation.SimulateMsgCreateCollege(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateCollege int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateCollege, &weightMsgUpdateCollege, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCollege = defaultWeightMsgUpdateCollege
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCollege,
		zksequencersimulation.SimulateMsgUpdateCollege(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteCollege int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteCollege, &weightMsgDeleteCollege, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteCollege = defaultWeightMsgDeleteCollege
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteCollege,
		zksequencersimulation.SimulateMsgDeleteCollege(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddStudentViaZksequencer int
	simState.AppParams.GetOrGenerate(opWeightMsgAddStudentViaZksequencer, &weightMsgAddStudentViaZksequencer, nil,
		func(_ *rand.Rand) {
			weightMsgAddStudentViaZksequencer = defaultWeightMsgAddStudentViaZksequencer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddStudentViaZksequencer,
		zksequencersimulation.SimulateMsgAddStudentViaZksequencer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateCollege,
			defaultWeightMsgCreateCollege,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				zksequencersimulation.SimulateMsgCreateCollege(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateCollege,
			defaultWeightMsgUpdateCollege,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				zksequencersimulation.SimulateMsgUpdateCollege(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteCollege,
			defaultWeightMsgDeleteCollege,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				zksequencersimulation.SimulateMsgDeleteCollege(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddStudentViaZksequencer,
			defaultWeightMsgAddStudentViaZksequencer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				zksequencersimulation.SimulateMsgAddStudentViaZksequencer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
