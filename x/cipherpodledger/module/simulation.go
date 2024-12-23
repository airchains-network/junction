package cipherpodledger

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/airchains-network/junction/testutil/sample"
	cipherpodledgersimulation "github.com/airchains-network/junction/x/cipherpodledger/simulation"
	"github.com/airchains-network/junction/x/cipherpodledger/types"
)

// avoid unused import issue
var (
	_ = cipherpodledgersimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgRegisterFhvm = "op_weight_msg_register_fhvm"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterFhvm int = 100

	opWeightMsgSubmitPod = "op_weight_msg_submit_pod"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitPod int = 100

	opWeightMsgVerifyPod = "op_weight_msg_verify_pod"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVerifyPod int = 100

	opWeightMsgLogBlobData = "op_weight_msg_log_blob_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLogBlobData int = 100

	opWeightMsgIntegrityCheck = "op_weight_msg_integrity_check"
	// TODO: Determine the simulation weight value
	defaultWeightMsgIntegrityCheck int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cipherpodledgerGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cipherpodledgerGenesis)
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

	var weightMsgRegisterFhvm int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterFhvm, &weightMsgRegisterFhvm, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterFhvm = defaultWeightMsgRegisterFhvm
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterFhvm,
		cipherpodledgersimulation.SimulateMsgRegisterFhvm(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitPod int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitPod, &weightMsgSubmitPod, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitPod = defaultWeightMsgSubmitPod
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitPod,
		cipherpodledgersimulation.SimulateMsgSubmitPod(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVerifyPod int
	simState.AppParams.GetOrGenerate(opWeightMsgVerifyPod, &weightMsgVerifyPod, nil,
		func(_ *rand.Rand) {
			weightMsgVerifyPod = defaultWeightMsgVerifyPod
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVerifyPod,
		cipherpodledgersimulation.SimulateMsgVerifyPod(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLogBlobData int
	simState.AppParams.GetOrGenerate(opWeightMsgLogBlobData, &weightMsgLogBlobData, nil,
		func(_ *rand.Rand) {
			weightMsgLogBlobData = defaultWeightMsgLogBlobData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLogBlobData,
		cipherpodledgersimulation.SimulateMsgLogBlobData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgIntegrityCheck int
	simState.AppParams.GetOrGenerate(opWeightMsgIntegrityCheck, &weightMsgIntegrityCheck, nil,
		func(_ *rand.Rand) {
			weightMsgIntegrityCheck = defaultWeightMsgIntegrityCheck
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgIntegrityCheck,
		cipherpodledgersimulation.SimulateMsgIntegrityCheck(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterFhvm,
			defaultWeightMsgRegisterFhvm,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cipherpodledgersimulation.SimulateMsgRegisterFhvm(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitPod,
			defaultWeightMsgSubmitPod,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cipherpodledgersimulation.SimulateMsgSubmitPod(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVerifyPod,
			defaultWeightMsgVerifyPod,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cipherpodledgersimulation.SimulateMsgVerifyPod(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgLogBlobData,
			defaultWeightMsgLogBlobData,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cipherpodledgersimulation.SimulateMsgLogBlobData(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgIntegrityCheck,
			defaultWeightMsgIntegrityCheck,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cipherpodledgersimulation.SimulateMsgIntegrityCheck(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
