package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitStation{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSchemaCreation{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSchemaEngage{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMigrateSchema{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
