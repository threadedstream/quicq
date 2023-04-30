package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"math/big"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/quic-go/quic-go"
)

func main() {
	tlsConf, err := getTLS()
	if err != nil {
		log.Fatalf("err: %s", err.Error())
	}
	server, err := quic.ListenAddr(":5000", tlsConf, &quic.Config{
		Allow0RTT: func(net.Addr) bool { return true },
	})
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGABRT, syscall.SIGINT)

	go func() {
		for {
			select {
			case _ = <-ch:
				server.Close()
				return
			default:
				continue
			}
		}
	}()

	log.Println("listening on port 5000")
	for {
		_, err := server.Accept(context.Background())
		if err != nil {
			log.Printf("err: %s", err.Error())
			return
		}
	}
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
		NextProtos:         []string{"quicq"},
	}, nil
}
