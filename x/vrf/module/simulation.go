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
	opWeightMsgRegisterCollection = "op_weight_msg_register_collection"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterCollection int = 100

	opWeightMsgInitiateVrf = "op_weight_msg_initiate_vrf"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitiateVrf int = 100

	opWeightMsgValidateVrf = "op_weight_msg_validate_vrf"
	// TODO: Determine the simulation weight value
	defaultWeightMsgValidateVrf int = 100

	opWeightMsgProcessVrfDispute = "op_weight_msg_process_vrf_dispute"
	// TODO: Determine the simulation weight value
	defaultWeightMsgProcessVrfDispute int = 100

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
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&vrfGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterCollection int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterCollection, &weightMsgRegisterCollection, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterCollection = defaultWeightMsgRegisterCollection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterCollection,
		vrfsimulation.SimulateMsgRegisterCollection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgInitiateVrf int
	simState.AppParams.GetOrGenerate(opWeightMsgInitiateVrf, &weightMsgInitiateVrf, nil,
		func(_ *rand.Rand) {
			weightMsgInitiateVrf = defaultWeightMsgInitiateVrf
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInitiateVrf,
		vrfsimulation.SimulateMsgInitiateVrf(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgValidateVrf int
	simState.AppParams.GetOrGenerate(opWeightMsgValidateVrf, &weightMsgValidateVrf, nil,
		func(_ *rand.Rand) {
			weightMsgValidateVrf = defaultWeightMsgValidateVrf
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgValidateVrf,
		vrfsimulation.SimulateMsgValidateVrf(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgProcessVrfDispute int
	simState.AppParams.GetOrGenerate(opWeightMsgProcessVrfDispute, &weightMsgProcessVrfDispute, nil,
		func(_ *rand.Rand) {
			weightMsgProcessVrfDispute = defaultWeightMsgProcessVrfDispute
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProcessVrfDispute,
		vrfsimulation.SimulateMsgProcessVrfDispute(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterCollection,
			defaultWeightMsgRegisterCollection,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vrfsimulation.SimulateMsgRegisterCollection(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgInitiateVrf,
			defaultWeightMsgInitiateVrf,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vrfsimulation.SimulateMsgInitiateVrf(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgValidateVrf,
			defaultWeightMsgValidateVrf,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vrfsimulation.SimulateMsgValidateVrf(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgProcessVrfDispute,
			defaultWeightMsgProcessVrfDispute,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vrfsimulation.SimulateMsgProcessVrfDispute(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
