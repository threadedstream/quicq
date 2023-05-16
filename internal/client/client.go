package client

import (
	"context"

	"github.com/quic-go/quic-go"
	"github.com/threadedstream/quicthing/internal/conn"
)

type Client interface {
	RequestStream() (conn.Stream, error)
	Send(data []byte) error
	Rcv(data []byte) error
}

type QuicQClient struct {
	quic.Connection
	messageBus conn.Stream
}

// Dial dials a server host
func Dial(ctx context.Context, addr string) (*QuicQClient, error) {
	tlsFn := conn.QuicQTLSFunc()
	tls, err := tlsFn()
	if err != nil {
		return nil, err
	}
	c, err := quic.DialAddrContext(ctx, addr, tls, &quic.Config{})
	if err != nil {
		return nil, err
	}
	return &QuicQClient{Connection: c}, nil
}

// RequestStream attempts to open a stream with a peer
func (qc *QuicQClient) RequestStream() (conn.Stream, error) {
	stream, err := qc.OpenStream()
	if err != nil {
		return nil, err
	}
	s := &conn.QuicQStream{Stream: stream}
	qc.messageBus = s
	return s, nil
}

// Send sends data over a bus
func (qc *QuicQClient) Send(data []byte) error {
	_, err := qc.messageBus.Send(data)
	if err != nil {
		return err
	}
	return nil
}

// Rcv receives data from a bus
func (qc *QuicQClient) Rcv(data []byte) error {
	_, err := qc.messageBus.Rcv(data)
	if err != nil {
		return err
	}
	return nil
}
