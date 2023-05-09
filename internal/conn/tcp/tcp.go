package tcp

import (
	"context"
	"github.com/threadedstream/quicthing/internal/conn"
	"log"
	"net"
)

type QuicQTcpConn struct {
	*net.TCPListener
	*net.TCPConn
}

func (tc *QuicQTcpConn) OpenStream() (conn.Stream, error) {
	return nil, nil
}

func (tc *QuicQTcpConn) AcceptStream(ctx context.Context) (conn.Stream, error) {
	childConn, err := tc.TCPListener.AcceptTCP()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(ctx)
	return &QuicQTcpStream{
		Ctx:    ctx,
		Cancel: cancel,
		Conn:   childConn,
	}, nil
}

func (tc *QuicQTcpConn) ReceiveDatagram() ([]byte, error) {
	panic("unimplemented!")
}

func (tc *QuicQTcpConn) SendDatagram([]byte) error {
	panic("unimplemented!")
}

func (tc *QuicQTcpConn) Log(format string, args ...any) {
	return
}

func (tc *QuicQTcpConn) RemoteAddr() net.Addr {
	return &net.TCPAddr{}
}

type QuicQTcpStream struct {
	Conn   *net.TCPConn
	Ctx    context.Context
	Cancel context.CancelFunc
}

func (ts *QuicQTcpStream) Send(p []byte) (int, error) {
	return ts.Conn.Write(p)
}

func (ts *QuicQTcpStream) Rcv(p []byte) (int, error) {
	return ts.Conn.Read(p)
}

func (ts *QuicQTcpStream) Log(format string, args ...any) {
	format = "[StreamID xxx] => " + format
	log.Printf(format, args...)
}

func (ts *QuicQTcpStream) Context() context.Context {
	return ts.Ctx
}

func (ts *QuicQTcpStream) Shutdown() error {
	_ = ts.Conn.CloseRead()
	_ = ts.Conn.CloseWrite()
	ts.Cancel()
	return ts.Conn.Close()
}
