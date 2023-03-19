// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: factory.proto

package factory

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
	FactoryService_Replenish_FullMethodName = "/grpc_gen.factory.FactoryService/Replenish"
)

// FactoryServiceClient is the client API for FactoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FactoryServiceClient interface {
	Replenish(ctx context.Context, in *FactoryRequest, opts ...grpc.CallOption) (*FactoryResponse, error)
}

type factoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFactoryServiceClient(cc grpc.ClientConnInterface) FactoryServiceClient {
	return &factoryServiceClient{cc}
}

func (c *factoryServiceClient) Replenish(ctx context.Context, in *FactoryRequest, opts ...grpc.CallOption) (*FactoryResponse, error) {
	out := new(FactoryResponse)
	err := c.cc.Invoke(ctx, FactoryService_Replenish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FactoryServiceServer is the server API for FactoryService service.
// All implementations must embed UnimplementedFactoryServiceServer
// for forward compatibility
type FactoryServiceServer interface {
	Replenish(context.Context, *FactoryRequest) (*FactoryResponse, error)
	mustEmbedUnimplementedFactoryServiceServer()
}

// UnimplementedFactoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFactoryServiceServer struct {
}

func (UnimplementedFactoryServiceServer) Replenish(context.Context, *FactoryRequest) (*FactoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Replenish not implemented")
}
func (UnimplementedFactoryServiceServer) mustEmbedUnimplementedFactoryServiceServer() {}

// UnsafeFactoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FactoryServiceServer will
// result in compilation errors.
type UnsafeFactoryServiceServer interface {
	mustEmbedUnimplementedFactoryServiceServer()
}

func RegisterFactoryServiceServer(s grpc.ServiceRegistrar, srv FactoryServiceServer) {
	s.RegisterService(&FactoryService_ServiceDesc, srv)
}

func _FactoryService_Replenish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FactoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FactoryServiceServer).Replenish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FactoryService_Replenish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FactoryServiceServer).Replenish(ctx, req.(*FactoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FactoryService_ServiceDesc is the grpc.ServiceDesc for FactoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FactoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_gen.factory.FactoryService",
	HandlerType: (*FactoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Replenish",
			Handler:    _FactoryService_Replenish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "factory.proto",
}
