package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgInitStation{}, "junction/InitStation", nil)
	cdc.RegisterConcrete(&MsgSubmitPod{}, "junction/SubmitPod", nil)
	cdc.RegisterConcrete(&MsgVerifyPod{}, "junction/VerifyPod", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitStation{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitPod{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVerifyPod{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
