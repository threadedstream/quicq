package client

import (
	"context"
	"fmt"

	"github.com/quic-go/quic-go"
	"github.com/threadedstream/quicthing/internal/conn"
)

type Client interface {
	Dial(context.Context, string) (conn.Connection, error)
	RequestStream() (conn.Stream, error)
	Send(data []byte) error
	Rcv(data []byte) error
}

type QuicQClient struct {
	quic.Connection
	messageBus conn.Stream
}

// New initializes QuicQClient
func New() *QuicQClient {
	return &QuicQClient{}
}

// Dial dials a server host
func (qc *QuicQClient) Dial(ctx context.Context, addr string) error {
	tlsFn := conn.QuicQTLSFunc()
	tls, err := tlsFn()
	if err != nil {
		return err
	}
	conn, err := quic.DialAddrContext(ctx, addr, tls, &quic.Config{})
	if err != nil {
		return err
	}
	qc.Connection = conn
	return nil
}

// RequestStream attempts to open a stream with a peer
func (qc *QuicQClient) RequestStream() error {
	stream, err := qc.OpenStream()
	if err != nil {
		return err
	}
	qc.messageBus = &conn.QuicQStream{Stream: stream}
	return err
}

// Send sends data over a bus
func (qc *QuicQClient) Send(data []byte) error {
	n, err := qc.messageBus.Send(data)
	if err != nil {
		return err
	}
	if n != len(data) {
		return fmt.Errorf("partial send!")
	}
	return nil
}

// Rcv receives data from a bus
func (qc *QuicQClient) Rcv(data []byte) error {
	n, err := qc.messageBus.Rcv(data)
	if err != nil {
		return err
	}
	if n != len(data) {
		return fmt.Errorf("partial rcv!")
	}
	return nil
}
