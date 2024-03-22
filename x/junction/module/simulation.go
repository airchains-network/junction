package junction

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/airchains-network/junction/testutil/sample"
	junctionsimulation "github.com/airchains-network/junction/x/junction/simulation"
	"github.com/airchains-network/junction/x/junction/types"
)

// avoid unused import issue
var (
	_ = junctionsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgInitStation = "op_weight_msg_init_station"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitStation int = 100

	opWeightMsgSubmitPod = "op_weight_msg_submit_pod"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitPod int = 100

	opWeightMsgVerifyPod = "op_weight_msg_verify_pod"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVerifyPod int = 100

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
	junctionGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&junctionGenesis)
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
		junctionsimulation.SimulateMsgInitStation(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitPod int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitPod, &weightMsgSubmitPod, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitPod = defaultWeightMsgSubmitPod
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitPod,
		junctionsimulation.SimulateMsgSubmitPod(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVerifyPod int
	simState.AppParams.GetOrGenerate(opWeightMsgVerifyPod, &weightMsgVerifyPod, nil,
		func(_ *rand.Rand) {
			weightMsgVerifyPod = defaultWeightMsgVerifyPod
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVerifyPod,
		junctionsimulation.SimulateMsgVerifyPod(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgInitiateVrf int
	simState.AppParams.GetOrGenerate(opWeightMsgInitiateVrf, &weightMsgInitiateVrf, nil,
		func(_ *rand.Rand) {
			weightMsgInitiateVrf = defaultWeightMsgInitiateVrf
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInitiateVrf,
		junctionsimulation.SimulateMsgInitiateVrf(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgValidateVrf int
	simState.AppParams.GetOrGenerate(opWeightMsgValidateVrf, &weightMsgValidateVrf, nil,
		func(_ *rand.Rand) {
			weightMsgValidateVrf = defaultWeightMsgValidateVrf
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgValidateVrf,
		junctionsimulation.SimulateMsgValidateVrf(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgProcessVrfDispute int
	simState.AppParams.GetOrGenerate(opWeightMsgProcessVrfDispute, &weightMsgProcessVrfDispute, nil,
		func(_ *rand.Rand) {
			weightMsgProcessVrfDispute = defaultWeightMsgProcessVrfDispute
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProcessVrfDispute,
		junctionsimulation.SimulateMsgProcessVrfDispute(am.accountKeeper, am.bankKeeper, am.keeper),
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
				junctionsimulation.SimulateMsgInitStation(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitPod,
			defaultWeightMsgSubmitPod,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				junctionsimulation.SimulateMsgSubmitPod(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVerifyPod,
			defaultWeightMsgVerifyPod,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				junctionsimulation.SimulateMsgVerifyPod(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgInitiateVrf,
			defaultWeightMsgInitiateVrf,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				junctionsimulation.SimulateMsgInitiateVrf(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgValidateVrf,
			defaultWeightMsgValidateVrf,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				junctionsimulation.SimulateMsgValidateVrf(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgProcessVrfDispute,
			defaultWeightMsgProcessVrfDispute,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				junctionsimulation.SimulateMsgProcessVrfDispute(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
