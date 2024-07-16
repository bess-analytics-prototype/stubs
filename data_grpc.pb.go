// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: data.proto

package mypackage

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	BessService_ProcessBess_FullMethodName = "/mypackage.BessService/ProcessBess"
)

// BessServiceClient is the client API for BessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BessServiceClient interface {
	ProcessBess(ctx context.Context, in *BessRequest, opts ...grpc.CallOption) (*BessResponse, error)
}

type bessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBessServiceClient(cc grpc.ClientConnInterface) BessServiceClient {
	return &bessServiceClient{cc}
}

func (c *bessServiceClient) ProcessBess(ctx context.Context, in *BessRequest, opts ...grpc.CallOption) (*BessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BessResponse)
	err := c.cc.Invoke(ctx, BessService_ProcessBess_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BessServiceServer is the server API for BessService service.
// All implementations must embed UnimplementedBessServiceServer
// for forward compatibility
type BessServiceServer interface {
	ProcessBess(context.Context, *BessRequest) (*BessResponse, error)
	mustEmbedUnimplementedBessServiceServer()
}

// UnimplementedBessServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBessServiceServer struct {
}

func (UnimplementedBessServiceServer) ProcessBess(context.Context, *BessRequest) (*BessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessBess not implemented")
}
func (UnimplementedBessServiceServer) mustEmbedUnimplementedBessServiceServer() {}

// UnsafeBessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BessServiceServer will
// result in compilation errors.
type UnsafeBessServiceServer interface {
	mustEmbedUnimplementedBessServiceServer()
}

func RegisterBessServiceServer(s grpc.ServiceRegistrar, srv BessServiceServer) {
	s.RegisterService(&BessService_ServiceDesc, srv)
}

func _BessService_ProcessBess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BessServiceServer).ProcessBess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BessService_ProcessBess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BessServiceServer).ProcessBess(ctx, req.(*BessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BessService_ServiceDesc is the grpc.ServiceDesc for BessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mypackage.BessService",
	HandlerType: (*BessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessBess",
			Handler:    _BessService_ProcessBess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}
