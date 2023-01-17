package server

import (
	"errors"
	"github.com/itsabgr/grpctun/pkg/proto"
	"google.golang.org/grpc/metadata"
	"net"
)

type ProxyServer struct {
	proto.UnimplementedProxyServer
}

func (p *ProxyServer) BindUDP(ctx proto.Proxy_BindUDPServer) error {
	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		return err
	}
	defer conn.Close()
	go func() {
		defer conn.Close()
		b := make([]byte, 8000)
		for {
			n, from, err := conn.ReadFrom(b)
			if err != nil {
				return
			}
			if nil != ctx.Send(&proto.PacketMessage{
				Data:   b[:n],
				Origin: from.String(),
			}) {
				return
			}
		}
	}()
	for {
		inp, err := ctx.Recv()
		if err != nil {
			return err
		}
		addr, err := net.ResolveUDPAddr("udp", inp.Origin)
		if err != nil {
			return err
		}
		_, err = conn.WriteTo(inp.Data, addr)
		if err != nil {
			return err
		}
	}
}

func (p *ProxyServer) ConnectTCP(ctx proto.Proxy_ConnectTCPServer) error {
	hdr, ok := metadata.FromIncomingContext(ctx.Context())
	if !ok || len(hdr.Get("x-target")) != 1 {
		return errors.New("bad request")
	}
	xTarget := hdr.Get("x-target")[0]
	conn, err := net.Dial("tcp", xTarget)
	if err != nil {
		return err
	}
	defer conn.Close()
	go func() {
		defer conn.Close()
		b := make([]byte, 8000)
		for {
			n, err := conn.Read(b)
			if err != nil {
				return
			}
			if nil != ctx.Send(&proto.StreamMessage{Data: b[:n]}) {
				return
			}
		}
	}()
	for {
		inp, err := ctx.Recv()
		if err != nil {
			return err
		}
		_, err = conn.Write(inp.Data)
		if err != nil {
			return err
		}
	}
}

func New() proto.ProxyServer {
	return &ProxyServer{}
}
