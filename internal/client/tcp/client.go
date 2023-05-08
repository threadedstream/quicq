package tcp

import (
	"context"
	"errors"
	"github.com/threadedstream/quicthing/internal/conn"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/threadedstream/quicthing/internal/conn/tcp"
)

type QuicQTCPClient struct {
	raddr *net.TCPAddr
	conn  *tcp.QuicQTcpConn
}

func Dial(ctx context.Context, addr string) (*QuicQTCPClient, error) {
	if parts := strings.Split(addr, ":"); len(parts) == 2 {
		host, port := parts[0], func() int { val, _ := strconv.Atoi(parts[1]); return val }()
		raddr := &net.TCPAddr{
			IP:   net.ParseIP(host),
			Port: port,
		}
		c, err := net.DialTCP("tcp4", nil, raddr)
		if err != nil {
			log.Println("could not dial tcp address: %s" + err.Error())
			return nil, err
		}
		return &QuicQTCPClient{conn: &tcp.QuicQTcpConn{TCPConn: c}, raddr: raddr}, nil
	}
	return nil, errors.New("address should adhere to the format host:port")
}

func (tc *QuicQTCPClient) RequestStream() (conn.Stream, error) {
	_, err := net.DialTCP("tcp4", nil, tc.raddr)
	return nil, err
}
