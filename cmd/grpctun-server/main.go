package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/itsabgr/grpctun/internal/selfsign"
	"github.com/itsabgr/grpctun/pkg/proto"
	server "github.com/itsabgr/grpctun/pkg/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
)

var flagPort = flag.Int("port", 443, "server port")
var flagCert = flag.String("cert", "", "certificate path")
var flagKey = flag.String("key", "", "key path")

func init() {
	flag.Parse()
}

func main() {
	var err error
	var ln net.Listener
	ln, err = net.Listen("tcp", fmt.Sprintf(":%d", *flagPort))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var creds credentials.TransportCredentials
	if len(*flagKey) == 0 {
		creds = credentials.NewTLS(&tls.Config{Certificates: []tls.Certificate{selfsign.Generate()}})
	} else {
		creds, err = credentials.NewServerTLSFromFile(*flagCert, *flagKey)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	defer grpcServer.Stop()
	proto.RegisterProxyServer(grpcServer, server.New())
	err = grpcServer.Serve(ln)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
