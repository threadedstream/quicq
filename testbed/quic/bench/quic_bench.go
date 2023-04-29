package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"math/big"
	"time"

	"github.com/quic-go/quic-go"
)

var (
	tlsConf *tls.Config
)

func testConnect() {
	conn, err := quic.DialAddr(":5000", tlsConf, &quic.Config{})
	if err != nil {
		log.Fatal(err)
	}
	conn.CloseWithError(quic.ApplicationErrorCode(quic.NoError), "closed")
	if err != nil {
		log.Fatalf("close err: %s", err.Error())
	}
}

const (
	step1 int64 = 1e2
	step2 int64 = 1e3
	step3 int64 = 1e4
	step4 int64 = 1e5
)

func measure(num int64) {
	start := time.Now()
	for i := int64(0); i < num; i++ {
		testConnect()
	}
	end := time.Now().Sub(start).Seconds()
	log.Printf("%d conn: %f\n", num, end)
}

func main() {
	tlsConf, _ = getTLS()
	measure(step1)
	measure(step2)
	measure(step3)
	measure(step4)
}

func getTLS() (*tls.Config, error) {
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
