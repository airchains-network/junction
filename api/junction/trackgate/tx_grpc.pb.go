// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: junction/trackgate/tx.proto

package trackgate

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
	Msg_UpdateParams_FullMethodName   = "/junction.trackgate.Msg/UpdateParams"
	Msg_InitStation_FullMethodName    = "/junction.trackgate.Msg/InitStation"
	Msg_SchemaCreation_FullMethodName = "/junction.trackgate.Msg/SchemaCreation"
	Msg_SchemaEngage_FullMethodName   = "/junction.trackgate.Msg/SchemaEngage"
	Msg_MigrateSchema_FullMethodName  = "/junction.trackgate.Msg/MigrateSchema"
	Msg_AuditSequencer_FullMethodName = "/junction.trackgate.Msg/AuditSequencer"
	Msg_LogBlobData_FullMethodName    = "/junction.trackgate.Msg/LogBlobData"
	Msg_IntegrityCheck_FullMethodName = "/junction.trackgate.Msg/IntegrityCheck"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	InitStation(ctx context.Context, in *MsgInitStation, opts ...grpc.CallOption) (*MsgInitStationResponse, error)
	SchemaCreation(ctx context.Context, in *MsgSchemaCreation, opts ...grpc.CallOption) (*MsgSchemaCreationResponse, error)
	SchemaEngage(ctx context.Context, in *MsgSchemaEngage, opts ...grpc.CallOption) (*MsgSchemaEngageResponse, error)
	MigrateSchema(ctx context.Context, in *MsgMigrateSchema, opts ...grpc.CallOption) (*MsgMigrateSchemaResponse, error)
	AuditSequencer(ctx context.Context, in *MsgAuditSequencer, opts ...grpc.CallOption) (*MsgAuditSequencerResponse, error)
	LogBlobData(ctx context.Context, in *MsgLogBlobData, opts ...grpc.CallOption) (*MsgLogBlobDataResponse, error)
	IntegrityCheck(ctx context.Context, in *MsgIntegrityCheck, opts ...grpc.CallOption) (*MsgIntegrityCheckResponse, error)
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

func (c *msgClient) InitStation(ctx context.Context, in *MsgInitStation, opts ...grpc.CallOption) (*MsgInitStationResponse, error) {
	out := new(MsgInitStationResponse)
	err := c.cc.Invoke(ctx, Msg_InitStation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SchemaCreation(ctx context.Context, in *MsgSchemaCreation, opts ...grpc.CallOption) (*MsgSchemaCreationResponse, error) {
	out := new(MsgSchemaCreationResponse)
	err := c.cc.Invoke(ctx, Msg_SchemaCreation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SchemaEngage(ctx context.Context, in *MsgSchemaEngage, opts ...grpc.CallOption) (*MsgSchemaEngageResponse, error) {
	out := new(MsgSchemaEngageResponse)
	err := c.cc.Invoke(ctx, Msg_SchemaEngage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MigrateSchema(ctx context.Context, in *MsgMigrateSchema, opts ...grpc.CallOption) (*MsgMigrateSchemaResponse, error) {
	out := new(MsgMigrateSchemaResponse)
	err := c.cc.Invoke(ctx, Msg_MigrateSchema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AuditSequencer(ctx context.Context, in *MsgAuditSequencer, opts ...grpc.CallOption) (*MsgAuditSequencerResponse, error) {
	out := new(MsgAuditSequencerResponse)
	err := c.cc.Invoke(ctx, Msg_AuditSequencer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) LogBlobData(ctx context.Context, in *MsgLogBlobData, opts ...grpc.CallOption) (*MsgLogBlobDataResponse, error) {
	out := new(MsgLogBlobDataResponse)
	err := c.cc.Invoke(ctx, Msg_LogBlobData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) IntegrityCheck(ctx context.Context, in *MsgIntegrityCheck, opts ...grpc.CallOption) (*MsgIntegrityCheckResponse, error) {
	out := new(MsgIntegrityCheckResponse)
	err := c.cc.Invoke(ctx, Msg_IntegrityCheck_FullMethodName, in, out, opts...)
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
	InitStation(context.Context, *MsgInitStation) (*MsgInitStationResponse, error)
	SchemaCreation(context.Context, *MsgSchemaCreation) (*MsgSchemaCreationResponse, error)
	SchemaEngage(context.Context, *MsgSchemaEngage) (*MsgSchemaEngageResponse, error)
	MigrateSchema(context.Context, *MsgMigrateSchema) (*MsgMigrateSchemaResponse, error)
	AuditSequencer(context.Context, *MsgAuditSequencer) (*MsgAuditSequencerResponse, error)
	LogBlobData(context.Context, *MsgLogBlobData) (*MsgLogBlobDataResponse, error)
	IntegrityCheck(context.Context, *MsgIntegrityCheck) (*MsgIntegrityCheckResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) InitStation(context.Context, *MsgInitStation) (*MsgInitStationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitStation not implemented")
}
func (UnimplementedMsgServer) SchemaCreation(context.Context, *MsgSchemaCreation) (*MsgSchemaCreationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchemaCreation not implemented")
}
func (UnimplementedMsgServer) SchemaEngage(context.Context, *MsgSchemaEngage) (*MsgSchemaEngageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchemaEngage not implemented")
}
func (UnimplementedMsgServer) MigrateSchema(context.Context, *MsgMigrateSchema) (*MsgMigrateSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MigrateSchema not implemented")
}
func (UnimplementedMsgServer) AuditSequencer(context.Context, *MsgAuditSequencer) (*MsgAuditSequencerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuditSequencer not implemented")
}
func (UnimplementedMsgServer) LogBlobData(context.Context, *MsgLogBlobData) (*MsgLogBlobDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogBlobData not implemented")
}
func (UnimplementedMsgServer) IntegrityCheck(context.Context, *MsgIntegrityCheck) (*MsgIntegrityCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IntegrityCheck not implemented")
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

func _Msg_InitStation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInitStation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InitStation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_InitStation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InitStation(ctx, req.(*MsgInitStation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SchemaCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSchemaCreation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SchemaCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SchemaCreation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SchemaCreation(ctx, req.(*MsgSchemaCreation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SchemaEngage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSchemaEngage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SchemaEngage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SchemaEngage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SchemaEngage(ctx, req.(*MsgSchemaEngage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MigrateSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMigrateSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MigrateSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MigrateSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MigrateSchema(ctx, req.(*MsgMigrateSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AuditSequencer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAuditSequencer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AuditSequencer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AuditSequencer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AuditSequencer(ctx, req.(*MsgAuditSequencer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_LogBlobData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLogBlobData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LogBlobData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_LogBlobData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LogBlobData(ctx, req.(*MsgLogBlobData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_IntegrityCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgIntegrityCheck)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).IntegrityCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_IntegrityCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).IntegrityCheck(ctx, req.(*MsgIntegrityCheck))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "junction.trackgate.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "InitStation",
			Handler:    _Msg_InitStation_Handler,
		},
		{
			MethodName: "SchemaCreation",
			Handler:    _Msg_SchemaCreation_Handler,
		},
		{
			MethodName: "SchemaEngage",
			Handler:    _Msg_SchemaEngage_Handler,
		},
		{
			MethodName: "MigrateSchema",
			Handler:    _Msg_MigrateSchema_Handler,
		},
		{
			MethodName: "AuditSequencer",
			Handler:    _Msg_AuditSequencer_Handler,
		},
		{
			MethodName: "LogBlobData",
			Handler:    _Msg_LogBlobData_Handler,
		},
		{
			MethodName: "IntegrityCheck",
			Handler:    _Msg_IntegrityCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "junction/trackgate/tx.proto",
}
