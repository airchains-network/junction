// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: junction/vrf/tx.proto

package vrf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_UpdateParams_FullMethodName       = "/junction.vrf.Msg/UpdateParams"
	Msg_RegisterCollection_FullMethodName = "/junction.vrf.Msg/RegisterCollection"
	Msg_InitiateVrf_FullMethodName        = "/junction.vrf.Msg/InitiateVrf"
	Msg_ValidateVrf_FullMethodName        = "/junction.vrf.Msg/ValidateVrf"
	Msg_ProcessVrfDispute_FullMethodName  = "/junction.vrf.Msg/ProcessVrfDispute"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	RegisterCollection(ctx context.Context, in *MsgRegisterCollection, opts ...grpc.CallOption) (*MsgRegisterCollectionResponse, error)
	InitiateVrf(ctx context.Context, in *MsgInitiateVrf, opts ...grpc.CallOption) (*MsgInitiateVrfResponse, error)
	ValidateVrf(ctx context.Context, in *MsgValidateVrf, opts ...grpc.CallOption) (*MsgValidateVrfResponse, error)
	ProcessVrfDispute(ctx context.Context, in *MsgProcessVrfDispute, opts ...grpc.CallOption) (*MsgProcessVrfDisputeResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RegisterCollection(ctx context.Context, in *MsgRegisterCollection, opts ...grpc.CallOption) (*MsgRegisterCollectionResponse, error) {
	out := new(MsgRegisterCollectionResponse)
	err := c.cc.Invoke(ctx, Msg_RegisterCollection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) InitiateVrf(ctx context.Context, in *MsgInitiateVrf, opts ...grpc.CallOption) (*MsgInitiateVrfResponse, error) {
	out := new(MsgInitiateVrfResponse)
	err := c.cc.Invoke(ctx, Msg_InitiateVrf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ValidateVrf(ctx context.Context, in *MsgValidateVrf, opts ...grpc.CallOption) (*MsgValidateVrfResponse, error) {
	out := new(MsgValidateVrfResponse)
	err := c.cc.Invoke(ctx, Msg_ValidateVrf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ProcessVrfDispute(ctx context.Context, in *MsgProcessVrfDispute, opts ...grpc.CallOption) (*MsgProcessVrfDisputeResponse, error) {
	out := new(MsgProcessVrfDisputeResponse)
	err := c.cc.Invoke(ctx, Msg_ProcessVrfDispute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	RegisterCollection(context.Context, *MsgRegisterCollection) (*MsgRegisterCollectionResponse, error)
	InitiateVrf(context.Context, *MsgInitiateVrf) (*MsgInitiateVrfResponse, error)
	ValidateVrf(context.Context, *MsgValidateVrf) (*MsgValidateVrfResponse, error)
	ProcessVrfDispute(context.Context, *MsgProcessVrfDispute) (*MsgProcessVrfDisputeResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) RegisterCollection(context.Context, *MsgRegisterCollection) (*MsgRegisterCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCollection not implemented")
}
func (UnimplementedMsgServer) InitiateVrf(context.Context, *MsgInitiateVrf) (*MsgInitiateVrfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitiateVrf not implemented")
}
func (UnimplementedMsgServer) ValidateVrf(context.Context, *MsgValidateVrf) (*MsgValidateVrfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateVrf not implemented")
}
func (UnimplementedMsgServer) ProcessVrfDispute(context.Context, *MsgProcessVrfDispute) (*MsgProcessVrfDisputeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessVrfDispute not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RegisterCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RegisterCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterCollection(ctx, req.(*MsgRegisterCollection))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_InitiateVrf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInitiateVrf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InitiateVrf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_InitiateVrf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InitiateVrf(ctx, req.(*MsgInitiateVrf))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ValidateVrf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgValidateVrf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ValidateVrf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ValidateVrf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ValidateVrf(ctx, req.(*MsgValidateVrf))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ProcessVrfDispute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgProcessVrfDispute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ProcessVrfDispute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ProcessVrfDispute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ProcessVrfDispute(ctx, req.(*MsgProcessVrfDispute))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "junction.vrf.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "RegisterCollection",
			Handler:    _Msg_RegisterCollection_Handler,
		},
		{
			MethodName: "InitiateVrf",
			Handler:    _Msg_InitiateVrf_Handler,
		},
		{
			MethodName: "ValidateVrf",
			Handler:    _Msg_ValidateVrf_Handler,
		},
		{
			MethodName: "ProcessVrfDispute",
			Handler:    _Msg_ProcessVrfDispute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "junction/vrf/tx.proto",
}
