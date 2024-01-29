// Setup gRPC protobuf meta

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: provider.proto

package provider

import (
	context "context"
	common "github.com/cble-platform/cble-provider-grpc/pkg/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Provider_Handshake_FullMethodName  = "/Provider/Handshake"
	Provider_Configure_FullMethodName  = "/Provider/Configure"
	Provider_Deploy_FullMethodName     = "/Provider/Deploy"
	Provider_Destroy_FullMethodName    = "/Provider/Destroy"
	Provider_GetConsole_FullMethodName = "/Provider/GetConsole"
)

// ProviderClient is the client API for Provider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderClient interface {
	Handshake(ctx context.Context, in *common.HandshakeRequest, opts ...grpc.CallOption) (*common.HandshakeReply, error)
	Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*common.RPCStatusReply, error)
	Deploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*DeployReply, error)
	Destroy(ctx context.Context, in *DestroyRequest, opts ...grpc.CallOption) (*DestroyReply, error)
	GetConsole(ctx context.Context, in *GetConsoleRequest, opts ...grpc.CallOption) (*GetConsoleReply, error)
}

type providerClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderClient(cc grpc.ClientConnInterface) ProviderClient {
	return &providerClient{cc}
}

func (c *providerClient) Handshake(ctx context.Context, in *common.HandshakeRequest, opts ...grpc.CallOption) (*common.HandshakeReply, error) {
	out := new(common.HandshakeReply)
	err := c.cc.Invoke(ctx, Provider_Handshake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*common.RPCStatusReply, error) {
	out := new(common.RPCStatusReply)
	err := c.cc.Invoke(ctx, Provider_Configure_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) Deploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*DeployReply, error) {
	out := new(DeployReply)
	err := c.cc.Invoke(ctx, Provider_Deploy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) Destroy(ctx context.Context, in *DestroyRequest, opts ...grpc.CallOption) (*DestroyReply, error) {
	out := new(DestroyReply)
	err := c.cc.Invoke(ctx, Provider_Destroy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) GetConsole(ctx context.Context, in *GetConsoleRequest, opts ...grpc.CallOption) (*GetConsoleReply, error) {
	out := new(GetConsoleReply)
	err := c.cc.Invoke(ctx, Provider_GetConsole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderServer is the server API for Provider service.
// All implementations must embed UnimplementedProviderServer
// for forward compatibility
type ProviderServer interface {
	Handshake(context.Context, *common.HandshakeRequest) (*common.HandshakeReply, error)
	Configure(context.Context, *ConfigureRequest) (*common.RPCStatusReply, error)
	Deploy(context.Context, *DeployRequest) (*DeployReply, error)
	Destroy(context.Context, *DestroyRequest) (*DestroyReply, error)
	GetConsole(context.Context, *GetConsoleRequest) (*GetConsoleReply, error)
	mustEmbedUnimplementedProviderServer()
}

// UnimplementedProviderServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServer struct {
}

func (UnimplementedProviderServer) Handshake(context.Context, *common.HandshakeRequest) (*common.HandshakeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handshake not implemented")
}
func (UnimplementedProviderServer) Configure(context.Context, *ConfigureRequest) (*common.RPCStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedProviderServer) Deploy(context.Context, *DeployRequest) (*DeployReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deploy not implemented")
}
func (UnimplementedProviderServer) Destroy(context.Context, *DestroyRequest) (*DestroyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedProviderServer) GetConsole(context.Context, *GetConsoleRequest) (*GetConsoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsole not implemented")
}
func (UnimplementedProviderServer) mustEmbedUnimplementedProviderServer() {}

// UnsafeProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServer will
// result in compilation errors.
type UnsafeProviderServer interface {
	mustEmbedUnimplementedProviderServer()
}

func RegisterProviderServer(s grpc.ServiceRegistrar, srv ProviderServer) {
	s.RegisterService(&Provider_ServiceDesc, srv)
}

func _Provider_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.HandshakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provider_Handshake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).Handshake(ctx, req.(*common.HandshakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provider_Configure_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).Configure(ctx, req.(*ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provider_Deploy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).Deploy(ctx, req.(*DeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provider_Destroy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).Destroy(ctx, req.(*DestroyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_GetConsole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConsoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).GetConsole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provider_GetConsole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).GetConsole(ctx, req.(*GetConsoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Provider_ServiceDesc is the grpc.ServiceDesc for Provider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Provider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Provider",
	HandlerType: (*ProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handshake",
			Handler:    _Provider_Handshake_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Provider_Configure_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _Provider_Deploy_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _Provider_Destroy_Handler,
		},
		{
			MethodName: "GetConsole",
			Handler:    _Provider_GetConsole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider.proto",
}
