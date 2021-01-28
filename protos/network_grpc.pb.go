// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

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

// NetworkClient is the client API for Network service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkClient interface {
	//Node->Master - Requesting to join the network
	RequestJoin(ctx context.Context, in *RequestJoinRequest, opts ...grpc.CallOption) (*RequestJoinResponse, error)
}

type networkClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkClient(cc grpc.ClientConnInterface) NetworkClient {
	return &networkClient{cc}
}

func (c *networkClient) RequestJoin(ctx context.Context, in *RequestJoinRequest, opts ...grpc.CallOption) (*RequestJoinResponse, error) {
	out := new(RequestJoinResponse)
	err := c.cc.Invoke(ctx, "/Network/RequestJoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServer is the server API for Network service.
// All implementations must embed UnimplementedNetworkServer
// for forward compatibility
type NetworkServer interface {
	//Node->Master - Requesting to join the network
	RequestJoin(context.Context, *RequestJoinRequest) (*RequestJoinResponse, error)
	mustEmbedUnimplementedNetworkServer()
}

// UnimplementedNetworkServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkServer struct {
}

func (UnimplementedNetworkServer) RequestJoin(context.Context, *RequestJoinRequest) (*RequestJoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestJoin not implemented")
}
func (UnimplementedNetworkServer) mustEmbedUnimplementedNetworkServer() {}

// UnsafeNetworkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkServer will
// result in compilation errors.
type UnsafeNetworkServer interface {
	mustEmbedUnimplementedNetworkServer()
}

func RegisterNetworkServer(s grpc.ServiceRegistrar, srv NetworkServer) {
	s.RegisterService(&Network_ServiceDesc, srv)
}

func _Network_RequestJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestJoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).RequestJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Network/RequestJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).RequestJoin(ctx, req.(*RequestJoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Network_ServiceDesc is the grpc.ServiceDesc for Network service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Network_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Network",
	HandlerType: (*NetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestJoin",
			Handler:    _Network_RequestJoin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/network.proto",
}
