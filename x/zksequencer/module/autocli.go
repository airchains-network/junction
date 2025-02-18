package zksequencer

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/airchains-network/junction/api/junction/zksequencer"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "CollegeAll",
					Use:       "list-college",
					Short:     "List all college",
				},
				{
					RpcMethod:      "College",
					Use:            "show-college [id]",
					Short:          "Shows a college by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "QueryVrfStudent",
					Use:            "query-vrf-student [id] [name] [details]",
					Short:          "Query query-vrf-student",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "name"}, {ProtoField: "details"}},
				},

				{
					RpcMethod:      "QueryVrfStudent2",
					Use:            "query-vrf-student-2 [id]",
					Short:          "Query query-vrf-student2",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateCollege",
					Use:            "create-college [students] [name] [details]",
					Short:          "Create college",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "students"}, {ProtoField: "name"}, {ProtoField: "details"}},
				},
				{
					RpcMethod:      "UpdateCollege",
					Use:            "update-college [id] [students] [name] [details]",
					Short:          "Update college",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "students"}, {ProtoField: "name"}, {ProtoField: "details"}},
				},
				{
					RpcMethod:      "DeleteCollege",
					Use:            "delete-college [id]",
					Short:          "Delete college",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "AddStudentViaZksequencer",
					Use:            "add-student-via-zksequencer [name] [details]",
					Short:          "Send a add-student-via-zksequencer tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "details"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
