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

	opWeightMsgSchemaEngage = "op_weight_msg_schema_engage"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSchemaEngage int = 100

	opWeightMsgMigrateSchema = "op_weight_msg_migrate_schema"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMigrateSchema int = 100

	opWeightMsgAuditSequencer = "op_weight_msg_audit_sequencer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAuditSequencer int = 100

	opWeightMsgLogBlobData = "op_weight_msg_log_blob_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLogBlobData int = 100

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

	var weightMsgSchemaEngage int
	simState.AppParams.GetOrGenerate(opWeightMsgSchemaEngage, &weightMsgSchemaEngage, nil,
		func(_ *rand.Rand) {
			weightMsgSchemaEngage = defaultWeightMsgSchemaEngage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSchemaEngage,
		trackgatesimulation.SimulateMsgSchemaEngage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMigrateSchema int
	simState.AppParams.GetOrGenerate(opWeightMsgMigrateSchema, &weightMsgMigrateSchema, nil,
		func(_ *rand.Rand) {
			weightMsgMigrateSchema = defaultWeightMsgMigrateSchema
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMigrateSchema,
		trackgatesimulation.SimulateMsgMigrateSchema(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAuditSequencer int
	simState.AppParams.GetOrGenerate(opWeightMsgAuditSequencer, &weightMsgAuditSequencer, nil,
		func(_ *rand.Rand) {
			weightMsgAuditSequencer = defaultWeightMsgAuditSequencer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAuditSequencer,
		trackgatesimulation.SimulateMsgAuditSequencer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLogBlobData int
	simState.AppParams.GetOrGenerate(opWeightMsgLogBlobData, &weightMsgLogBlobData, nil,
		func(_ *rand.Rand) {
			weightMsgLogBlobData = defaultWeightMsgLogBlobData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLogBlobData,
		trackgatesimulation.SimulateMsgLogBlobData(am.accountKeeper, am.bankKeeper, am.keeper),
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
		simulation.NewWeightedProposalMsg(
			opWeightMsgSchemaEngage,
			defaultWeightMsgSchemaEngage,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				trackgatesimulation.SimulateMsgSchemaEngage(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgMigrateSchema,
			defaultWeightMsgMigrateSchema,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				trackgatesimulation.SimulateMsgMigrateSchema(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAuditSequencer,
			defaultWeightMsgAuditSequencer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				trackgatesimulation.SimulateMsgAuditSequencer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgLogBlobData,
			defaultWeightMsgLogBlobData,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				trackgatesimulation.SimulateMsgLogBlobData(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
