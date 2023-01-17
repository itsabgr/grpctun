package dialer

import (
	"net"
	"time"
)

type packet2 struct {
	raddr  net.Addr
	packet net.PacketConn
}

func (p *packet2) Read(b []byte) (n int, err error) {
	return p.packet.WriteTo(b, p.raddr)
}

func (p *packet2) Write(b []byte) (n int, err error) {
	return p.packet.WriteTo(b, p.raddr)
}

func (p *packet2) Close() error {
	return p.packet.Close()
}

func (p *packet2) LocalAddr() net.Addr {
	return p.packet.LocalAddr()
}

func (p *packet2) RemoteAddr() net.Addr {
	return p.raddr
}

func (p *packet2) SetDeadline(t time.Time) error {
	return p.packet.SetDeadline(t)
}

func (p *packet2) SetReadDeadline(t time.Time) error {
	return p.packet.SetReadDeadline(t)

}

func (p *packet2) SetWriteDeadline(t time.Time) error {
	return p.packet.SetWriteDeadline(t)

}

func PacketToConn(packet net.PacketConn, host string) (net.Conn, error) {
	raddr, err := net.ResolveUDPAddr(packet.LocalAddr().Network(), host)
	if err != nil {
		return nil, err
	}
	return &packet2{raddr, packet}, nil

}
