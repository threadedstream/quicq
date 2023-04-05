package conn

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"math/big"
)

type Connection interface {
	GenerateTLS() (*tls.Config, error)
	Setup()
	Send([]byte) error
	Rcv() ([]byte, error)
}

func DefaultGenerateTLSFunc() func() (*tls.Config, error) {
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
			Certificates: []tls.Certificate{tlsCert},
			NextProtos:   []string{"quicq"},
		}, nil
	}
}
