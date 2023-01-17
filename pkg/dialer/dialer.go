package dialer

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/itsabgr/grpctun/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"net"
	"time"
)

type dialer struct {
	network      string
	laddr, raddr net.Addr
	stream       proto.Proxy_ConnectTCPClient
	conn         *grpc.ClientConn
}

func (d *dialer) Read(b []byte) (n int, err error) {
	inp, err := d.stream.Recv()
	if err != nil {
		return 0, err
	}
	return copy(b, inp.Data), err
}

func (d *dialer) Write(b []byte) (n int, err error) {
	return len(b), d.stream.Send(&proto.StreamMessage{Data: b})
}

func (d *dialer) Close() error {
	return d.conn.Close()
}

func (d *dialer) LocalAddr() net.Addr {
	return d.laddr
}

func (d *dialer) RemoteAddr() net.Addr {
	return d.raddr
}

func (d *dialer) SetDeadline(t time.Time) error {
	d.SetReadDeadline(t)
	return d.SetWriteDeadline(t)
}

func (d *dialer) SetReadDeadline(t time.Time) error {
	return nil
}

func (d *dialer) SetWriteDeadline(t time.Time) error {
	return nil
}

func DialContext(ctx context.Context, insecureSkipVerify bool, proxy, network, host string) (net.Conn, error) {
	switch network {
	case "tcp", "tcp4", "tcp6":
		laddr, err := net.ResolveTCPAddr("tcp", proxy)
		if err != nil {
			return nil, err
		}
		raddr, err := net.ResolveTCPAddr(network, host)
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
		ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"x-target": host}))
		stream, err := proto.NewProxyClient(conn).ConnectTCP(ctx)
		if err != nil {
			conn.Close()
			return nil, err
		}
		return &dialer{network, laddr, raddr, stream, conn}, nil
	default:
		return nil, fmt.Errorf("unsupported network %q", network)
	}

}
