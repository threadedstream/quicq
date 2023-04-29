package server

import (
	"context"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/threadedstream/quicthing/internal/conn"
)

// Server is a common Server interface
type Server interface {
	Serve(addr string) error
	AcceptClient(context.Context) (conn.Connection, error)
	Handle(command any) (any, error)
	AddOnShutdownCallback(fn func())
	Shutdown() error
	Close() error
}

// QuicServer is a quic server
type QuicServer struct {
	quic.Connection
	listener            quic.Listener
	onShutdownCallbacks []func()
}

func (qs *QuicServer) Serve(addr string) error {
	tlsFn := conn.QuicQTLSFunc()
	tlsConf, err := tlsFn()
	if err != nil {
		return err
	}

	// try 0RTT thing
	conf := &quic.Config{
		EnableDatagrams: true,
		MaxIdleTimeout:  time.Second * 10,
	}
	listener, err := quic.ListenAddr(addr, tlsConf, conf)
	if err != nil {
		panic(err)
	}
	qs.listener = listener
	// close listener upon shutdown
	qs.onShutdownCallbacks = append(qs.onShutdownCallbacks, func() {
		qs.listener.Close()
	})
	return nil
}

// AcceptClient accepts remote client
func (qs *QuicServer) AcceptClient(ctx context.Context) (conn.Connection, error) {
	quicConn, err := qs.listener.Accept(ctx)
	if nil != err {
		return nil, err
	}
	quicQConn := &conn.QuicQConn{Connection: quicConn}
	return quicQConn, nil
}

func (qs *QuicServer) AddOnShutdownCallback(fn func()) {
	qs.onShutdownCallbacks = append(qs.onShutdownCallbacks, fn)
}

func (qs *QuicServer) Shutdown() error {
	// execute each on-shutdown callback
	for _, cb := range qs.onShutdownCallbacks {
		cb()
	}
	return qs.listener.Close()
}
