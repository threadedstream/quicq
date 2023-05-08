package tcp

import (
	"context"
	"errors"
	"github.com/threadedstream/quicthing/internal/conn"
	"github.com/threadedstream/quicthing/internal/conn/tcp"
	"net"
	"strconv"
	"strings"
)

type QuicQTCPServer struct {
	conn                *tcp.QuicQTcpConn
	onShutdownCallbacks []func()
}

func (ts *QuicQTCPServer) Serve(addr string) error {
	if parts := strings.Split(addr, ":"); len(parts) == 2 {
		localAddr := &net.TCPAddr{
			IP:   net.ParseIP(parts[0]),
			Port: func() int { val, _ := strconv.Atoi(parts[1]); return val }(),
		}
		listener, err := net.ListenTCP("tcp4", localAddr)
		if err != nil {
			return err
		}
		ts.conn = &tcp.QuicQTcpConn{
			TCPListener: listener,
		}
	}
	return errors.New("could not parse address")
}

func (ts *QuicQTCPServer) AcceptClient(ctx context.Context) (conn.Connection, error) {
	c, err := ts.conn.AcceptTCP()
	if err != nil {
		return nil, err
	}
	return &tcp.QuicQTcpConn{TCPConn: c}, nil
}

func (ts *QuicQTCPServer) AddOnShutdownCallback(fn func()) {
	ts.onShutdownCallbacks = append(ts.onShutdownCallbacks, fn)
}

func (ts *QuicQTCPServer) Shutdown() error {
	// execute each on-shutdown callback
	for _, cb := range ts.onShutdownCallbacks {
		cb()
	}
	return ts.conn.Close()
}
