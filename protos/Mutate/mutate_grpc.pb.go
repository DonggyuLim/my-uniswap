// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: protos/Mutate/mutate.proto

package mutate

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

// MutateClient is the client API for Mutate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MutateClient interface {
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	Approve(ctx context.Context, in *ApproveRequest, opts ...grpc.CallOption) (*ApproveResponse, error)
}

type mutateClient struct {
	cc grpc.ClientConnInterface
}

func NewMutateClient(cc grpc.ClientConnInterface) MutateClient {
	return &mutateClient{cc}
}

func (c *mutateClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, "/erc20.Mutate/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mutateClient) Approve(ctx context.Context, in *ApproveRequest, opts ...grpc.CallOption) (*ApproveResponse, error) {
	out := new(ApproveResponse)
	err := c.cc.Invoke(ctx, "/erc20.Mutate/Approve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MutateServer is the server API for Mutate service.
// All implementations must embed UnimplementedMutateServer
// for forward compatibility
type MutateServer interface {
	Transfer(context.Context, *TransferRequest) (*TransferResponse, error)
	Approve(context.Context, *ApproveRequest) (*ApproveResponse, error)
	mustEmbedUnimplementedMutateServer()
}

// UnimplementedMutateServer must be embedded to have forward compatible implementations.
type UnimplementedMutateServer struct {
}

func (UnimplementedMutateServer) Transfer(context.Context, *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedMutateServer) Approve(context.Context, *ApproveRequest) (*ApproveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Approve not implemented")
}
func (UnimplementedMutateServer) mustEmbedUnimplementedMutateServer() {}

// UnsafeMutateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MutateServer will
// result in compilation errors.
type UnsafeMutateServer interface {
	mustEmbedUnimplementedMutateServer()
}

func RegisterMutateServer(s grpc.ServiceRegistrar, srv MutateServer) {
	s.RegisterService(&Mutate_ServiceDesc, srv)
}

func _Mutate_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutateServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/erc20.Mutate/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutateServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mutate_Approve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutateServer).Approve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/erc20.Mutate/Approve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutateServer).Approve(ctx, req.(*ApproveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mutate_ServiceDesc is the grpc.ServiceDesc for Mutate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mutate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "erc20.Mutate",
	HandlerType: (*MutateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _Mutate_Transfer_Handler,
		},
		{
			MethodName: "Approve",
			Handler:    _Mutate_Approve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/Mutate/mutate.proto",
}
