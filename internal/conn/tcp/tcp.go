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
		ctx:    ctx,
		cancel: cancel,
		conn:   childConn,
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
	return nil
}

type QuicQTcpStream struct {
	conn   *net.TCPConn
	ctx    context.Context
	cancel context.CancelFunc
}

func (ts *QuicQTcpStream) Send(p []byte) (int, error) {
	return ts.conn.Write(p)
}

func (ts *QuicQTcpStream) Rcv(p []byte) (int, error) {
	return ts.conn.Read(p)
}

func (ts *QuicQTcpStream) Log(format string, args ...any) {
	format = "[StreamID xxx] => " + format
	log.Printf(format, args...)
}

func (ts *QuicQTcpStream) Context() context.Context {
	return ts.ctx
}

func (ts *QuicQTcpStream) Shutdown() error {
	_ = ts.conn.CloseRead()
	_ = ts.conn.CloseWrite()
	ts.cancel()
	return ts.conn.Close()
}
