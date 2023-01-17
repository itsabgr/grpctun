package dialer

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/itsabgr/grpctun/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"time"
)

type packet struct {
	laddr  net.Addr
	stream proto.Proxy_BindUDPClient
	conn   *grpc.ClientConn
}

func (p_ *packet) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	inp, err := p_.stream.Recv()
	if err != nil {
		return 0, nil, err
	}
	addr, err = net.ResolveUDPAddr("udp", inp.Origin)
	if err != nil {
		return 0, nil, err
	}
	return copy(p, inp.Data), addr, nil
}

func (p_ *packet) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	return len(p), p_.stream.Send(&proto.PacketMessage{
		Data:   p,
		Origin: addr.String(),
	})
}

func (p_ *packet) Close() error {
	return p_.conn.Close()
}

func (p_ *packet) LocalAddr() net.Addr {
	return p_.laddr
}

func (p_ *packet) SetDeadline(t time.Time) error {
	p_.SetReadDeadline(t)
	return p_.SetWriteDeadline(t)
}

func (p_ *packet) SetReadDeadline(t time.Time) error {
	return nil
}

func (p_ *packet) SetWriteDeadline(t time.Time) error {
	return nil
}

func ListenPacketContext(ctx context.Context, insecureSkipVerify bool, proxy, network string) (net.PacketConn, error) {
	switch network {
	case "udp", "udp4", "udp6":
		laddr, err := net.ResolveTCPAddr("tcp", proxy)
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		conn, err := grpc.DialContext(ctx, proxy,
			grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
				InsecureSkipVerify: insecureSkipVerify,
			})),
		)
		if err != nil {
			return nil, err
		}
		stream, err := proto.NewProxyClient(conn).BindUDP(ctx)
		if err != nil {
			conn.Close()
			return nil, err
		}
		return &packet{laddr, stream, conn}, nil
	default:
		return nil, fmt.Errorf("unsupported network %q", network)
	}

}
