// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: junction/wasm/v1/tx.proto

package wasmv1

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
	Msg_StoreCode_FullMethodName                       = "/junction.wasm.v1.Msg/StoreCode"
	Msg_InstantiateContract_FullMethodName             = "/junction.wasm.v1.Msg/InstantiateContract"
	Msg_InstantiateContract2_FullMethodName            = "/junction.wasm.v1.Msg/InstantiateContract2"
	Msg_ExecuteContract_FullMethodName                 = "/junction.wasm.v1.Msg/ExecuteContract"
	Msg_MigrateContract_FullMethodName                 = "/junction.wasm.v1.Msg/MigrateContract"
	Msg_UpdateAdmin_FullMethodName                     = "/junction.wasm.v1.Msg/UpdateAdmin"
	Msg_ClearAdmin_FullMethodName                      = "/junction.wasm.v1.Msg/ClearAdmin"
	Msg_UpdateInstantiateConfig_FullMethodName         = "/junction.wasm.v1.Msg/UpdateInstantiateConfig"
	Msg_UpdateParams_FullMethodName                    = "/junction.wasm.v1.Msg/UpdateParams"
	Msg_SudoContract_FullMethodName                    = "/junction.wasm.v1.Msg/SudoContract"
	Msg_PinCodes_FullMethodName                        = "/junction.wasm.v1.Msg/PinCodes"
	Msg_UnpinCodes_FullMethodName                      = "/junction.wasm.v1.Msg/UnpinCodes"
	Msg_StoreAndInstantiateContract_FullMethodName     = "/junction.wasm.v1.Msg/StoreAndInstantiateContract"
	Msg_RemoveCodeUploadParamsAddresses_FullMethodName = "/junction.wasm.v1.Msg/RemoveCodeUploadParamsAddresses"
	Msg_AddCodeUploadParamsAddresses_FullMethodName    = "/junction.wasm.v1.Msg/AddCodeUploadParamsAddresses"
	Msg_StoreAndMigrateContract_FullMethodName         = "/junction.wasm.v1.Msg/StoreAndMigrateContract"
	Msg_UpdateContractLabel_FullMethodName             = "/junction.wasm.v1.Msg/UpdateContractLabel"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// StoreCode to submit Wasm code to the system
	StoreCode(ctx context.Context, in *MsgStoreCode, opts ...grpc.CallOption) (*MsgStoreCodeResponse, error)
	// InstantiateContract creates a new smart contract instance for the given
	// code id.
	InstantiateContract(ctx context.Context, in *MsgInstantiateContract, opts ...grpc.CallOption) (*MsgInstantiateContractResponse, error)
	// InstantiateContract2 creates a new smart contract instance for the given
	// code id with a predictable address
	InstantiateContract2(ctx context.Context, in *MsgInstantiateContract2, opts ...grpc.CallOption) (*MsgInstantiateContract2Response, error)
	// Execute submits the given message data to a smart contract
	ExecuteContract(ctx context.Context, in *MsgExecuteContract, opts ...grpc.CallOption) (*MsgExecuteContractResponse, error)
	// Migrate runs a code upgrade/ downgrade for a smart contract
	MigrateContract(ctx context.Context, in *MsgMigrateContract, opts ...grpc.CallOption) (*MsgMigrateContractResponse, error)
	// UpdateAdmin sets a new admin for a smart contract
	UpdateAdmin(ctx context.Context, in *MsgUpdateAdmin, opts ...grpc.CallOption) (*MsgUpdateAdminResponse, error)
	// ClearAdmin removes any admin stored for a smart contract
	ClearAdmin(ctx context.Context, in *MsgClearAdmin, opts ...grpc.CallOption) (*MsgClearAdminResponse, error)
	// UpdateInstantiateConfig updates instantiate config for a smart contract
	UpdateInstantiateConfig(ctx context.Context, in *MsgUpdateInstantiateConfig, opts ...grpc.CallOption) (*MsgUpdateInstantiateConfigResponse, error)
	// UpdateParams defines a governance operation for updating the x/wasm
	// module parameters. The authority is defined in the keeper.
	//
	// Since: 0.40
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	// SudoContract defines a governance operation for calling sudo
	// on a contract. The authority is defined in the keeper.
	//
	// Since: 0.40
	SudoContract(ctx context.Context, in *MsgSudoContract, opts ...grpc.CallOption) (*MsgSudoContractResponse, error)
	// PinCodes defines a governance operation for pinning a set of
	// code ids in the wasmvm cache. The authority is defined in the keeper.
	//
	// Since: 0.40
	PinCodes(ctx context.Context, in *MsgPinCodes, opts ...grpc.CallOption) (*MsgPinCodesResponse, error)
	// UnpinCodes defines a governance operation for unpinning a set of
	// code ids in the wasmvm cache. The authority is defined in the keeper.
	//
	// Since: 0.40
	UnpinCodes(ctx context.Context, in *MsgUnpinCodes, opts ...grpc.CallOption) (*MsgUnpinCodesResponse, error)
	// StoreAndInstantiateContract defines a governance operation for storing
	// and instantiating the contract. The authority is defined in the keeper.
	//
	// Since: 0.40
	StoreAndInstantiateContract(ctx context.Context, in *MsgStoreAndInstantiateContract, opts ...grpc.CallOption) (*MsgStoreAndInstantiateContractResponse, error)
	// RemoveCodeUploadParamsAddresses defines a governance operation for
	// removing addresses from code upload params.
	// The authority is defined in the keeper.
	RemoveCodeUploadParamsAddresses(ctx context.Context, in *MsgRemoveCodeUploadParamsAddresses, opts ...grpc.CallOption) (*MsgRemoveCodeUploadParamsAddressesResponse, error)
	// AddCodeUploadParamsAddresses defines a governance operation for
	// adding addresses to code upload params.
	// The authority is defined in the keeper.
	AddCodeUploadParamsAddresses(ctx context.Context, in *MsgAddCodeUploadParamsAddresses, opts ...grpc.CallOption) (*MsgAddCodeUploadParamsAddressesResponse, error)
	// StoreAndMigrateContract defines a governance operation for storing
	// and migrating the contract. The authority is defined in the keeper.
	//
	// Since: 0.42
	StoreAndMigrateContract(ctx context.Context, in *MsgStoreAndMigrateContract, opts ...grpc.CallOption) (*MsgStoreAndMigrateContractResponse, error)
	// UpdateContractLabel sets a new label for a smart contract
	//
	// Since: 0.43
	UpdateContractLabel(ctx context.Context, in *MsgUpdateContractLabel, opts ...grpc.CallOption) (*MsgUpdateContractLabelResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) StoreCode(ctx context.Context, in *MsgStoreCode, opts ...grpc.CallOption) (*MsgStoreCodeResponse, error) {
	out := new(MsgStoreCodeResponse)
	err := c.cc.Invoke(ctx, Msg_StoreCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) InstantiateContract(ctx context.Context, in *MsgInstantiateContract, opts ...grpc.CallOption) (*MsgInstantiateContractResponse, error) {
	out := new(MsgInstantiateContractResponse)
	err := c.cc.Invoke(ctx, Msg_InstantiateContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) InstantiateContract2(ctx context.Context, in *MsgInstantiateContract2, opts ...grpc.CallOption) (*MsgInstantiateContract2Response, error) {
	out := new(MsgInstantiateContract2Response)
	err := c.cc.Invoke(ctx, Msg_InstantiateContract2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ExecuteContract(ctx context.Context, in *MsgExecuteContract, opts ...grpc.CallOption) (*MsgExecuteContractResponse, error) {
	out := new(MsgExecuteContractResponse)
	err := c.cc.Invoke(ctx, Msg_ExecuteContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MigrateContract(ctx context.Context, in *MsgMigrateContract, opts ...grpc.CallOption) (*MsgMigrateContractResponse, error) {
	out := new(MsgMigrateContractResponse)
	err := c.cc.Invoke(ctx, Msg_MigrateContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateAdmin(ctx context.Context, in *MsgUpdateAdmin, opts ...grpc.CallOption) (*MsgUpdateAdminResponse, error) {
	out := new(MsgUpdateAdminResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClearAdmin(ctx context.Context, in *MsgClearAdmin, opts ...grpc.CallOption) (*MsgClearAdminResponse, error) {
	out := new(MsgClearAdminResponse)
	err := c.cc.Invoke(ctx, Msg_ClearAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateInstantiateConfig(ctx context.Context, in *MsgUpdateInstantiateConfig, opts ...grpc.CallOption) (*MsgUpdateInstantiateConfigResponse, error) {
	out := new(MsgUpdateInstantiateConfigResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateInstantiateConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SudoContract(ctx context.Context, in *MsgSudoContract, opts ...grpc.CallOption) (*MsgSudoContractResponse, error) {
	out := new(MsgSudoContractResponse)
	err := c.cc.Invoke(ctx, Msg_SudoContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) PinCodes(ctx context.Context, in *MsgPinCodes, opts ...grpc.CallOption) (*MsgPinCodesResponse, error) {
	out := new(MsgPinCodesResponse)
	err := c.cc.Invoke(ctx, Msg_PinCodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UnpinCodes(ctx context.Context, in *MsgUnpinCodes, opts ...grpc.CallOption) (*MsgUnpinCodesResponse, error) {
	out := new(MsgUnpinCodesResponse)
	err := c.cc.Invoke(ctx, Msg_UnpinCodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StoreAndInstantiateContract(ctx context.Context, in *MsgStoreAndInstantiateContract, opts ...grpc.CallOption) (*MsgStoreAndInstantiateContractResponse, error) {
	out := new(MsgStoreAndInstantiateContractResponse)
	err := c.cc.Invoke(ctx, Msg_StoreAndInstantiateContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveCodeUploadParamsAddresses(ctx context.Context, in *MsgRemoveCodeUploadParamsAddresses, opts ...grpc.CallOption) (*MsgRemoveCodeUploadParamsAddressesResponse, error) {
	out := new(MsgRemoveCodeUploadParamsAddressesResponse)
	err := c.cc.Invoke(ctx, Msg_RemoveCodeUploadParamsAddresses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddCodeUploadParamsAddresses(ctx context.Context, in *MsgAddCodeUploadParamsAddresses, opts ...grpc.CallOption) (*MsgAddCodeUploadParamsAddressesResponse, error) {
	out := new(MsgAddCodeUploadParamsAddressesResponse)
	err := c.cc.Invoke(ctx, Msg_AddCodeUploadParamsAddresses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StoreAndMigrateContract(ctx context.Context, in *MsgStoreAndMigrateContract, opts ...grpc.CallOption) (*MsgStoreAndMigrateContractResponse, error) {
	out := new(MsgStoreAndMigrateContractResponse)
	err := c.cc.Invoke(ctx, Msg_StoreAndMigrateContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateContractLabel(ctx context.Context, in *MsgUpdateContractLabel, opts ...grpc.CallOption) (*MsgUpdateContractLabelResponse, error) {
	out := new(MsgUpdateContractLabelResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateContractLabel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// StoreCode to submit Wasm code to the system
	StoreCode(context.Context, *MsgStoreCode) (*MsgStoreCodeResponse, error)
	// InstantiateContract creates a new smart contract instance for the given
	// code id.
	InstantiateContract(context.Context, *MsgInstantiateContract) (*MsgInstantiateContractResponse, error)
	// InstantiateContract2 creates a new smart contract instance for the given
	// code id with a predictable address
	InstantiateContract2(context.Context, *MsgInstantiateContract2) (*MsgInstantiateContract2Response, error)
	// Execute submits the given message data to a smart contract
	ExecuteContract(context.Context, *MsgExecuteContract) (*MsgExecuteContractResponse, error)
	// Migrate runs a code upgrade/ downgrade for a smart contract
	MigrateContract(context.Context, *MsgMigrateContract) (*MsgMigrateContractResponse, error)
	// UpdateAdmin sets a new admin for a smart contract
	UpdateAdmin(context.Context, *MsgUpdateAdmin) (*MsgUpdateAdminResponse, error)
	// ClearAdmin removes any admin stored for a smart contract
	ClearAdmin(context.Context, *MsgClearAdmin) (*MsgClearAdminResponse, error)
	// UpdateInstantiateConfig updates instantiate config for a smart contract
	UpdateInstantiateConfig(context.Context, *MsgUpdateInstantiateConfig) (*MsgUpdateInstantiateConfigResponse, error)
	// UpdateParams defines a governance operation for updating the x/wasm
	// module parameters. The authority is defined in the keeper.
	//
	// Since: 0.40
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	// SudoContract defines a governance operation for calling sudo
	// on a contract. The authority is defined in the keeper.
	//
	// Since: 0.40
	SudoContract(context.Context, *MsgSudoContract) (*MsgSudoContractResponse, error)
	// PinCodes defines a governance operation for pinning a set of
	// code ids in the wasmvm cache. The authority is defined in the keeper.
	//
	// Since: 0.40
	PinCodes(context.Context, *MsgPinCodes) (*MsgPinCodesResponse, error)
	// UnpinCodes defines a governance operation for unpinning a set of
	// code ids in the wasmvm cache. The authority is defined in the keeper.
	//
	// Since: 0.40
	UnpinCodes(context.Context, *MsgUnpinCodes) (*MsgUnpinCodesResponse, error)
	// StoreAndInstantiateContract defines a governance operation for storing
	// and instantiating the contract. The authority is defined in the keeper.
	//
	// Since: 0.40
	StoreAndInstantiateContract(context.Context, *MsgStoreAndInstantiateContract) (*MsgStoreAndInstantiateContractResponse, error)
	// RemoveCodeUploadParamsAddresses defines a governance operation for
	// removing addresses from code upload params.
	// The authority is defined in the keeper.
	RemoveCodeUploadParamsAddresses(context.Context, *MsgRemoveCodeUploadParamsAddresses) (*MsgRemoveCodeUploadParamsAddressesResponse, error)
	// AddCodeUploadParamsAddresses defines a governance operation for
	// adding addresses to code upload params.
	// The authority is defined in the keeper.
	AddCodeUploadParamsAddresses(context.Context, *MsgAddCodeUploadParamsAddresses) (*MsgAddCodeUploadParamsAddressesResponse, error)
	// StoreAndMigrateContract defines a governance operation for storing
	// and migrating the contract. The authority is defined in the keeper.
	//
	// Since: 0.42
	StoreAndMigrateContract(context.Context, *MsgStoreAndMigrateContract) (*MsgStoreAndMigrateContractResponse, error)
	// UpdateContractLabel sets a new label for a smart contract
	//
	// Since: 0.43
	UpdateContractLabel(context.Context, *MsgUpdateContractLabel) (*MsgUpdateContractLabelResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) StoreCode(context.Context, *MsgStoreCode) (*MsgStoreCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreCode not implemented")
}
func (UnimplementedMsgServer) InstantiateContract(context.Context, *MsgInstantiateContract) (*MsgInstantiateContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstantiateContract not implemented")
}
func (UnimplementedMsgServer) InstantiateContract2(context.Context, *MsgInstantiateContract2) (*MsgInstantiateContract2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstantiateContract2 not implemented")
}
func (UnimplementedMsgServer) ExecuteContract(context.Context, *MsgExecuteContract) (*MsgExecuteContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteContract not implemented")
}
func (UnimplementedMsgServer) MigrateContract(context.Context, *MsgMigrateContract) (*MsgMigrateContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MigrateContract not implemented")
}
func (UnimplementedMsgServer) UpdateAdmin(context.Context, *MsgUpdateAdmin) (*MsgUpdateAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdmin not implemented")
}
func (UnimplementedMsgServer) ClearAdmin(context.Context, *MsgClearAdmin) (*MsgClearAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearAdmin not implemented")
}
func (UnimplementedMsgServer) UpdateInstantiateConfig(context.Context, *MsgUpdateInstantiateConfig) (*MsgUpdateInstantiateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInstantiateConfig not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) SudoContract(context.Context, *MsgSudoContract) (*MsgSudoContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SudoContract not implemented")
}
func (UnimplementedMsgServer) PinCodes(context.Context, *MsgPinCodes) (*MsgPinCodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PinCodes not implemented")
}
func (UnimplementedMsgServer) UnpinCodes(context.Context, *MsgUnpinCodes) (*MsgUnpinCodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnpinCodes not implemented")
}
func (UnimplementedMsgServer) StoreAndInstantiateContract(context.Context, *MsgStoreAndInstantiateContract) (*MsgStoreAndInstantiateContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreAndInstantiateContract not implemented")
}
func (UnimplementedMsgServer) RemoveCodeUploadParamsAddresses(context.Context, *MsgRemoveCodeUploadParamsAddresses) (*MsgRemoveCodeUploadParamsAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCodeUploadParamsAddresses not implemented")
}
func (UnimplementedMsgServer) AddCodeUploadParamsAddresses(context.Context, *MsgAddCodeUploadParamsAddresses) (*MsgAddCodeUploadParamsAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCodeUploadParamsAddresses not implemented")
}
func (UnimplementedMsgServer) StoreAndMigrateContract(context.Context, *MsgStoreAndMigrateContract) (*MsgStoreAndMigrateContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreAndMigrateContract not implemented")
}
func (UnimplementedMsgServer) UpdateContractLabel(context.Context, *MsgUpdateContractLabel) (*MsgUpdateContractLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContractLabel not implemented")
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

func _Msg_StoreCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStoreCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StoreCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_StoreCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StoreCode(ctx, req.(*MsgStoreCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_InstantiateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInstantiateContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InstantiateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_InstantiateContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InstantiateContract(ctx, req.(*MsgInstantiateContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_InstantiateContract2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInstantiateContract2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InstantiateContract2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_InstantiateContract2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InstantiateContract2(ctx, req.(*MsgInstantiateContract2))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ExecuteContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExecuteContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExecuteContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ExecuteContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExecuteContract(ctx, req.(*MsgExecuteContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MigrateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMigrateContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MigrateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MigrateContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MigrateContract(ctx, req.(*MsgMigrateContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateAdmin(ctx, req.(*MsgUpdateAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClearAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClearAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClearAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ClearAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClearAdmin(ctx, req.(*MsgClearAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateInstantiateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateInstantiateConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateInstantiateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateInstantiateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateInstantiateConfig(ctx, req.(*MsgUpdateInstantiateConfig))
	}
	return interceptor(ctx, in, info, handler)
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

func _Msg_SudoContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSudoContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SudoContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SudoContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SudoContract(ctx, req.(*MsgSudoContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_PinCodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPinCodes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PinCodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_PinCodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PinCodes(ctx, req.(*MsgPinCodes))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UnpinCodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnpinCodes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UnpinCodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UnpinCodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UnpinCodes(ctx, req.(*MsgUnpinCodes))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StoreAndInstantiateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStoreAndInstantiateContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StoreAndInstantiateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_StoreAndInstantiateContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StoreAndInstantiateContract(ctx, req.(*MsgStoreAndInstantiateContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveCodeUploadParamsAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveCodeUploadParamsAddresses)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveCodeUploadParamsAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveCodeUploadParamsAddresses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveCodeUploadParamsAddresses(ctx, req.(*MsgRemoveCodeUploadParamsAddresses))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddCodeUploadParamsAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddCodeUploadParamsAddresses)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddCodeUploadParamsAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddCodeUploadParamsAddresses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddCodeUploadParamsAddresses(ctx, req.(*MsgAddCodeUploadParamsAddresses))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StoreAndMigrateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStoreAndMigrateContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StoreAndMigrateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_StoreAndMigrateContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StoreAndMigrateContract(ctx, req.(*MsgStoreAndMigrateContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateContractLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateContractLabel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateContractLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateContractLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateContractLabel(ctx, req.(*MsgUpdateContractLabel))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "junction.wasm.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreCode",
			Handler:    _Msg_StoreCode_Handler,
		},
		{
			MethodName: "InstantiateContract",
			Handler:    _Msg_InstantiateContract_Handler,
		},
		{
			MethodName: "InstantiateContract2",
			Handler:    _Msg_InstantiateContract2_Handler,
		},
		{
			MethodName: "ExecuteContract",
			Handler:    _Msg_ExecuteContract_Handler,
		},
		{
			MethodName: "MigrateContract",
			Handler:    _Msg_MigrateContract_Handler,
		},
		{
			MethodName: "UpdateAdmin",
			Handler:    _Msg_UpdateAdmin_Handler,
		},
		{
			MethodName: "ClearAdmin",
			Handler:    _Msg_ClearAdmin_Handler,
		},
		{
			MethodName: "UpdateInstantiateConfig",
			Handler:    _Msg_UpdateInstantiateConfig_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "SudoContract",
			Handler:    _Msg_SudoContract_Handler,
		},
		{
			MethodName: "PinCodes",
			Handler:    _Msg_PinCodes_Handler,
		},
		{
			MethodName: "UnpinCodes",
			Handler:    _Msg_UnpinCodes_Handler,
		},
		{
			MethodName: "StoreAndInstantiateContract",
			Handler:    _Msg_StoreAndInstantiateContract_Handler,
		},
		{
			MethodName: "RemoveCodeUploadParamsAddresses",
			Handler:    _Msg_RemoveCodeUploadParamsAddresses_Handler,
		},
		{
			MethodName: "AddCodeUploadParamsAddresses",
			Handler:    _Msg_AddCodeUploadParamsAddresses_Handler,
		},
		{
			MethodName: "StoreAndMigrateContract",
			Handler:    _Msg_StoreAndMigrateContract_Handler,
		},
		{
			MethodName: "UpdateContractLabel",
			Handler:    _Msg_UpdateContractLabel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "junction/wasm/v1/tx.proto",
}
