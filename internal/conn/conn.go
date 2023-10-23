package conn

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net"

	"github.com/quic-go/quic-go"
)

// Stream is a general stream interface used throughout the app
type Stream interface {
	Send([]byte) (int, error)
	Rcv() ([]byte, error)
	Log(string, ...any)
	Context() context.Context
	Shutdown() error
}

// Connection is an interface to quic connection
type Connection interface {
	OpenStream() (Stream, error)
	AcceptStream(context.Context) (Stream, error)
	ReceiveDatagram() ([]byte, error)
	SendDatagram([]byte) error
	Log(format string, args ...any)
	RemoteAddr() net.Addr
}

// QuicQConn is an implementation of Connection
type QuicQConn struct {
	quic.Connection
}

func (qc *QuicQConn) Log(format string, args ...any) {
	format = "[" + qc.RemoteAddr().String() + "] => " + format
	log.Printf(format, args...)
}

// ReceiveDatagram receives datagram from a peer
func (qc *QuicQConn) ReceiveDatagram() ([]byte, error) {
	return qc.Connection.ReceiveMessage()
}

// SendDatagram sends datagram to a peer
func (qc *QuicQConn) SendDatagram(data []byte) error {
	return qc.Connection.SendMessage(data)
}

// AcceptStream accepts remote stream
func (qc *QuicQConn) AcceptStream(ctx context.Context) (Stream, error) {
	// TODO(threadedstream): reuse streams using memory pool?
	stream, err := qc.Connection.AcceptStream(ctx)
	if err != nil {
		return nil, err
	}
	return &QuicQStream{Stream: stream}, nil
}

// OpenStream requests to initiate a quic stream
func (qc *QuicQConn) OpenStream() (Stream, error) {
	// TODO(threadedstream): reuse streams using memory pool?
	stream, err := qc.Connection.OpenStream()
	if err != nil {
		return nil, err
	}
	return &QuicQStream{Stream: stream}, nil
}

// RemoteAddr returns remote address of a connected host
func (qc *QuicQConn) RemoteAddr() net.Addr {
	return qc.Connection.RemoteAddr()
}

// QuicQStream is a Stream implementation
type QuicQStream struct {
	quic.Stream
}

// Send sends data over a quic stream
func (qs *QuicQStream) Send(p []byte) (int, error) {
	return qs.Stream.Write(p)
}

// Rcv receives data from a quic stream
func (qs *QuicQStream) Rcv() ([]byte, error) {
	bs := make([]byte, 8192)
	n, err := qs.Stream.Read(bs)
	if n <= len(bs) {
		bs = bs[:n]
	} else {
		return nil, errors.New("n > len(bs)")
	}
	return bs, err
}

// Context returns stream's context
func (qs *QuicQStream) Context() context.Context {
	return qs.Stream.Context()
}

// Shutdown completely disposes of stream canceling context as well
func (qs *QuicQStream) Shutdown() error {
	qs.Stream.CancelRead(quic.StreamErrorCode(quic.NoError))
	qs.Stream.CancelWrite(quic.StreamErrorCode(quic.NoError))
	return qs.Stream.Close()
}

func (qs *QuicQStream) Log(format string, args ...any) {
	format = "[StreamID " + fmt.Sprint(qs.Stream.StreamID()) + "] => " + format
	log.Printf(format, args...)
}

// QuicQTLSFunc generates a default in-memory tls config with no io access to external certificates
func QuicQTLSFunc() func() (*tls.Config, error) {
	return func() (*tls.Config, error) {
		key, err := rsa.GenerateKey(rand.Reader, 1024)
		if err != nil {
			return nil, err
		}
		template := x509.Certificate{SerialNumber: big.NewInt(1)}
		certDer, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
		if err != nil {
			return nil, err
		}
		keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
		certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDer})

		tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
		if err != nil {
			return nil, err
		}

		return &tls.Config{
			Certificates:       []tls.Certificate{tlsCert},
			NextProtos:         []string{"quicq"},
			InsecureSkipVerify: true,
		}, nil
	}
}
