package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"math/big"
	"net"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

var tlsConf *tls.Config

func dialNoTLS() {
	dialer := net.Dialer{
		Control: reusePort,
	}
	conn, err := dialer.Dial("tcp4", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Close()
	if err != nil {
		log.Fatalf("close err: %s", err.Error())
	}
}

func dialTLS() {
	dialer := tls.Dialer{
		NetDialer: &net.Dialer{
			Control: reusePort,
		},
		Config: tlsConf,
	}
	conn, err := dialer.Dial("tcp4", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	if err = conn.Close(); err != nil {
		log.Fatal(err)
	}
}

func testConnect() {
	dialTLS()
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
	var err error
	tlsConf, err = getTLS()
	if err != nil {
		log.Fatal(err)
	}
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
		InsecureSkipVerify: true,
	}, nil
}

func reusePort(network, address string, conn syscall.RawConn) error {
	return conn.Control(func(descriptor uintptr) {
		syscall.SetsockoptInt(int(descriptor), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
	})
}
