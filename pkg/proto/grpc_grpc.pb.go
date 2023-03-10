// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: grpc.proto

package proto

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

// ProxyClient is the client API for Proxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyClient interface {
	ConnectTCP(ctx context.Context, opts ...grpc.CallOption) (Proxy_ConnectTCPClient, error)
	BindUDP(ctx context.Context, opts ...grpc.CallOption) (Proxy_BindUDPClient, error)
}

type proxyClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyClient(cc grpc.ClientConnInterface) ProxyClient {
	return &proxyClient{cc}
}

func (c *proxyClient) ConnectTCP(ctx context.Context, opts ...grpc.CallOption) (Proxy_ConnectTCPClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[0], "/Proxy/ConnectTCP", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyConnectTCPClient{stream}
	return x, nil
}

type Proxy_ConnectTCPClient interface {
	Send(*StreamMessage) error
	Recv() (*StreamMessage, error)
	grpc.ClientStream
}

type proxyConnectTCPClient struct {
	grpc.ClientStream
}

func (x *proxyConnectTCPClient) Send(m *StreamMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *proxyConnectTCPClient) Recv() (*StreamMessage, error) {
	m := new(StreamMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *proxyClient) BindUDP(ctx context.Context, opts ...grpc.CallOption) (Proxy_BindUDPClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[1], "/Proxy/BindUDP", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyBindUDPClient{stream}
	return x, nil
}

type Proxy_BindUDPClient interface {
	Send(*PacketMessage) error
	Recv() (*PacketMessage, error)
	grpc.ClientStream
}

type proxyBindUDPClient struct {
	grpc.ClientStream
}

func (x *proxyBindUDPClient) Send(m *PacketMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *proxyBindUDPClient) Recv() (*PacketMessage, error) {
	m := new(PacketMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProxyServer is the server API for Proxy service.
// All implementations must embed UnimplementedProxyServer
// for forward compatibility
type ProxyServer interface {
	ConnectTCP(Proxy_ConnectTCPServer) error
	BindUDP(Proxy_BindUDPServer) error
	mustEmbedUnimplementedProxyServer()
}

// UnimplementedProxyServer must be embedded to have forward compatible implementations.
type UnimplementedProxyServer struct {
}

func (UnimplementedProxyServer) ConnectTCP(Proxy_ConnectTCPServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectTCP not implemented")
}
func (UnimplementedProxyServer) BindUDP(Proxy_BindUDPServer) error {
	return status.Errorf(codes.Unimplemented, "method BindUDP not implemented")
}
func (UnimplementedProxyServer) mustEmbedUnimplementedProxyServer() {}

// UnsafeProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyServer will
// result in compilation errors.
type UnsafeProxyServer interface {
	mustEmbedUnimplementedProxyServer()
}

func RegisterProxyServer(s grpc.ServiceRegistrar, srv ProxyServer) {
	s.RegisterService(&Proxy_ServiceDesc, srv)
}

func _Proxy_ConnectTCP_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProxyServer).ConnectTCP(&proxyConnectTCPServer{stream})
}

type Proxy_ConnectTCPServer interface {
	Send(*StreamMessage) error
	Recv() (*StreamMessage, error)
	grpc.ServerStream
}

type proxyConnectTCPServer struct {
	grpc.ServerStream
}

func (x *proxyConnectTCPServer) Send(m *StreamMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *proxyConnectTCPServer) Recv() (*StreamMessage, error) {
	m := new(StreamMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Proxy_BindUDP_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProxyServer).BindUDP(&proxyBindUDPServer{stream})
}

type Proxy_BindUDPServer interface {
	Send(*PacketMessage) error
	Recv() (*PacketMessage, error)
	grpc.ServerStream
}

type proxyBindUDPServer struct {
	grpc.ServerStream
}

func (x *proxyBindUDPServer) Send(m *PacketMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *proxyBindUDPServer) Recv() (*PacketMessage, error) {
	m := new(PacketMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Proxy_ServiceDesc is the grpc.ServiceDesc for Proxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Proxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Proxy",
	HandlerType: (*ProxyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectTCP",
			Handler:       _Proxy_ConnectTCP_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "BindUDP",
			Handler:       _Proxy_BindUDP_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc.proto",
}
