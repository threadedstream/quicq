package server

import (
	"context"

	"github.com/quic-go/quic-go"
	"github.com/threadedstream/quicthing/internal/conn"
)

type Connection interface{}

type Stream interface {
	Send([]byte) error
	Rcv() ([]byte, error)
}

type Server interface {
	Serve(addr string) error
	AcceptStream() (Stream, error)
	AcceptClient(context.Context) (Connection, error)
	Close() error
}

type QuicServer struct {
	quic.Connection
	listener            quic.Listener
	onShutdownCallbacks []func()
}

func (qs *QuicServer) Serve(addr string) error {
	tlsFn := conn.DefaultGenerateTLSFunc()
	tlsConf, err := tlsFn()
	if err != nil {
		return err
	}

	listener, err := quic.ListenAddr(addr, tlsConf, &quic.Config{})
	if err != nil {
		panic(err)
	}
	qs.listener = listener
	return nil
}

func (qs *QuicServer) AcceptStream() (Stream, error) {
	return nil, nil
}

func (qs *QuicServer) AcceptClient(ctx context.Context) (Connection, error) {
	return qs.listener.Accept(ctx)
}
