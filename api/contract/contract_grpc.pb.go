// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: contract/contract.proto

package contract

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ContractService_CreateContract_FullMethodName            = "/api.contract.ContractService/CreateContract"
	ContractService_UpdateContract_FullMethodName            = "/api.contract.ContractService/UpdateContract"
	ContractService_SignContract_FullMethodName              = "/api.contract.ContractService/SignContract"
	ContractService_DeleteContract_FullMethodName            = "/api.contract.ContractService/DeleteContract"
	ContractService_GetContract_FullMethodName               = "/api.contract.ContractService/GetContract"
	ContractService_ListContract_FullMethodName              = "/api.contract.ContractService/ListContract"
	ContractService_ListTotalContractEachRoom_FullMethodName = "/api.contract.ContractService/ListTotalContractEachRoom"
)

// ContractServiceClient is the client API for ContractService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContractServiceClient interface {
	CreateContract(ctx context.Context, in *CreateContractRequest, opts ...grpc.CallOption) (*CreateContractReply, error)
	UpdateContract(ctx context.Context, in *UpdateContractRequest, opts ...grpc.CallOption) (*UpdateContractReply, error)
	SignContract(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignReply, error)
	DeleteContract(ctx context.Context, in *DeleteContractRequest, opts ...grpc.CallOption) (*DeleteContractReply, error)
	GetContract(ctx context.Context, in *GetContractRequest, opts ...grpc.CallOption) (*GetContractReply, error)
	ListContract(ctx context.Context, in *ListContractRequest, opts ...grpc.CallOption) (*ListContractReply, error)
	ListTotalContractEachRoom(ctx context.Context, in *ListTotalContractEachRoomRequest, opts ...grpc.CallOption) (*ListTotalContractEachRoomReply, error)
}

type contractServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContractServiceClient(cc grpc.ClientConnInterface) ContractServiceClient {
	return &contractServiceClient{cc}
}

func (c *contractServiceClient) CreateContract(ctx context.Context, in *CreateContractRequest, opts ...grpc.CallOption) (*CreateContractReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateContractReply)
	err := c.cc.Invoke(ctx, ContractService_CreateContract_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) UpdateContract(ctx context.Context, in *UpdateContractRequest, opts ...grpc.CallOption) (*UpdateContractReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateContractReply)
	err := c.cc.Invoke(ctx, ContractService_UpdateContract_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) SignContract(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignReply)
	err := c.cc.Invoke(ctx, ContractService_SignContract_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) DeleteContract(ctx context.Context, in *DeleteContractRequest, opts ...grpc.CallOption) (*DeleteContractReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteContractReply)
	err := c.cc.Invoke(ctx, ContractService_DeleteContract_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) GetContract(ctx context.Context, in *GetContractRequest, opts ...grpc.CallOption) (*GetContractReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetContractReply)
	err := c.cc.Invoke(ctx, ContractService_GetContract_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) ListContract(ctx context.Context, in *ListContractRequest, opts ...grpc.CallOption) (*ListContractReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListContractReply)
	err := c.cc.Invoke(ctx, ContractService_ListContract_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) ListTotalContractEachRoom(ctx context.Context, in *ListTotalContractEachRoomRequest, opts ...grpc.CallOption) (*ListTotalContractEachRoomReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTotalContractEachRoomReply)
	err := c.cc.Invoke(ctx, ContractService_ListTotalContractEachRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContractServiceServer is the server API for ContractService service.
// All implementations must embed UnimplementedContractServiceServer
// for forward compatibility.
type ContractServiceServer interface {
	CreateContract(context.Context, *CreateContractRequest) (*CreateContractReply, error)
	UpdateContract(context.Context, *UpdateContractRequest) (*UpdateContractReply, error)
	SignContract(context.Context, *SignRequest) (*SignReply, error)
	DeleteContract(context.Context, *DeleteContractRequest) (*DeleteContractReply, error)
	GetContract(context.Context, *GetContractRequest) (*GetContractReply, error)
	ListContract(context.Context, *ListContractRequest) (*ListContractReply, error)
	ListTotalContractEachRoom(context.Context, *ListTotalContractEachRoomRequest) (*ListTotalContractEachRoomReply, error)
	mustEmbedUnimplementedContractServiceServer()
}

// UnimplementedContractServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedContractServiceServer struct{}

func (UnimplementedContractServiceServer) CreateContract(context.Context, *CreateContractRequest) (*CreateContractReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContract not implemented")
}
func (UnimplementedContractServiceServer) UpdateContract(context.Context, *UpdateContractRequest) (*UpdateContractReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContract not implemented")
}
func (UnimplementedContractServiceServer) SignContract(context.Context, *SignRequest) (*SignReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignContract not implemented")
}
func (UnimplementedContractServiceServer) DeleteContract(context.Context, *DeleteContractRequest) (*DeleteContractReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContract not implemented")
}
func (UnimplementedContractServiceServer) GetContract(context.Context, *GetContractRequest) (*GetContractReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContract not implemented")
}
func (UnimplementedContractServiceServer) ListContract(context.Context, *ListContractRequest) (*ListContractReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContract not implemented")
}
func (UnimplementedContractServiceServer) ListTotalContractEachRoom(context.Context, *ListTotalContractEachRoomRequest) (*ListTotalContractEachRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTotalContractEachRoom not implemented")
}
func (UnimplementedContractServiceServer) mustEmbedUnimplementedContractServiceServer() {}
func (UnimplementedContractServiceServer) testEmbeddedByValue()                         {}

// UnsafeContractServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContractServiceServer will
// result in compilation errors.
type UnsafeContractServiceServer interface {
	mustEmbedUnimplementedContractServiceServer()
}

func RegisterContractServiceServer(s grpc.ServiceRegistrar, srv ContractServiceServer) {
	// If the following call pancis, it indicates UnimplementedContractServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ContractService_ServiceDesc, srv)
}

func _ContractService_CreateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).CreateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractService_CreateContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).CreateContract(ctx, req.(*CreateContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_UpdateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).UpdateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractService_UpdateContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).UpdateContract(ctx, req.(*UpdateContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_SignContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).SignContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractService_SignContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).SignContract(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_DeleteContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).DeleteContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractService_DeleteContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).DeleteContract(ctx, req.(*DeleteContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_GetContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).GetContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractService_GetContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).GetContract(ctx, req.(*GetContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_ListContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).ListContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractService_ListContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).ListContract(ctx, req.(*ListContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_ListTotalContractEachRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTotalContractEachRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).ListTotalContractEachRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractService_ListTotalContractEachRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).ListTotalContractEachRoom(ctx, req.(*ListTotalContractEachRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContractService_ServiceDesc is the grpc.ServiceDesc for ContractService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContractService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.contract.ContractService",
	HandlerType: (*ContractServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContract",
			Handler:    _ContractService_CreateContract_Handler,
		},
		{
			MethodName: "UpdateContract",
			Handler:    _ContractService_UpdateContract_Handler,
		},
		{
			MethodName: "SignContract",
			Handler:    _ContractService_SignContract_Handler,
		},
		{
			MethodName: "DeleteContract",
			Handler:    _ContractService_DeleteContract_Handler,
		},
		{
			MethodName: "GetContract",
			Handler:    _ContractService_GetContract_Handler,
		},
		{
			MethodName: "ListContract",
			Handler:    _ContractService_ListContract_Handler,
		},
		{
			MethodName: "ListTotalContractEachRoom",
			Handler:    _ContractService_ListTotalContractEachRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract/contract.proto",
}
