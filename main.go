package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/itsabgr/grpctun/pkg/dialer"
	"github.com/things-go/go-socks5"
	"net"
)

var flagPort = flag.Int("port", 1080, "local socks5 interface port")
var flagServer = flag.String("server", "localhost:443", "proxy server addr")
var flagInsecure = flag.Bool("insecure", false, "insecure skip ssl certificates")

func init() {
	flag.Parse()
}

type resolver struct{}

func (r resolver) Resolve(ctx context.Context, _ string) (context.Context, net.IP, error) {
	return ctx, nil, nil
}

func main() {
	server := socks5.NewServer(
		socks5.WithResolver(resolver{}),
		socks5.WithDial(func(ctx context.Context, network, addr string) (net.Conn, error) {
			switch network {
			case "tcp", "tcp4", "tcp6":
				return dialer.DialContext(ctx, *flagInsecure, *flagServer, network, addr)
			case "udp", "udp4", "udp6":
				pconn, err := dialer.ListenPacketContext(ctx, *flagInsecure, *flagServer, network)
				if err != nil {
					return nil, err
				}
				conn, err := dialer.PacketToConn(pconn, addr)
				if err != nil {
					pconn.Close()
					return nil, err
				}
				return conn, nil
			}
			return net.Dial(network, addr)
		}))
	err := server.ListenAndServe("tcp", fmt.Sprintf("localhost:%d", *flagPort))
	if err != nil {
		fmt.Println(err)
	}
}
