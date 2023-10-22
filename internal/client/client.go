package client

import (
	"context"
	"time"

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
	c, err := quic.DialAddrContext(ctx, addr, tls, &quic.Config{
		MaxIdleTimeout: time.Minute * 5,
	})
	if err != nil {
		return err
	}
	qc.Connection = c
	return nil
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
func (qc *QuicQClient) Rcv() (bs []byte, err error) {
	bs, err = qc.messageBus.Rcv()
	if err != nil {
		return nil, err
	}
	return bs, nil
}
