package rollup

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/airchains-network/junction/testutil/sample"
	rollupsimulation "github.com/airchains-network/junction/x/rollup/simulation"
	"github.com/airchains-network/junction/x/rollup/types"
)

// avoid unused import issue
var (
	_ = rollupsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgInitRollup = "op_weight_msg_init_rollup"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitRollup int = 100

	opWeightMsgInitProver = "op_weight_msg_init_prover"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitProver int = 100

	opWeightMsgSubmitBatchMetadata = "op_weight_msg_submit_batch_metadata"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitBatchMetadata int = 100

	opWeightMsgSubmitBatch = "op_weight_msg_submit_batch"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitBatch int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	rollupGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&rollupGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgInitRollup int
	simState.AppParams.GetOrGenerate(opWeightMsgInitRollup, &weightMsgInitRollup, nil,
		func(_ *rand.Rand) {
			weightMsgInitRollup = defaultWeightMsgInitRollup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInitRollup,
		rollupsimulation.SimulateMsgInitRollup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgInitProver int
	simState.AppParams.GetOrGenerate(opWeightMsgInitProver, &weightMsgInitProver, nil,
		func(_ *rand.Rand) {
			weightMsgInitProver = defaultWeightMsgInitProver
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInitProver,
		rollupsimulation.SimulateMsgInitProver(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitBatchMetadata int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitBatchMetadata, &weightMsgSubmitBatchMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitBatchMetadata = defaultWeightMsgSubmitBatchMetadata
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitBatchMetadata,
		rollupsimulation.SimulateMsgSubmitBatchMetadata(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitBatch int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitBatch, &weightMsgSubmitBatch, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitBatch = defaultWeightMsgSubmitBatch
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitBatch,
		rollupsimulation.SimulateMsgSubmitBatch(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgInitRollup,
			defaultWeightMsgInitRollup,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rollupsimulation.SimulateMsgInitRollup(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgInitProver,
			defaultWeightMsgInitProver,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rollupsimulation.SimulateMsgInitProver(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitBatchMetadata,
			defaultWeightMsgSubmitBatchMetadata,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rollupsimulation.SimulateMsgSubmitBatchMetadata(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitBatch,
			defaultWeightMsgSubmitBatch,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rollupsimulation.SimulateMsgSubmitBatch(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
