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
	"time"

	"github.com/quic-go/quic-go"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	tlsConf, err := generateTLS()
	if err != nil {
		log.Fatal(err)
	}
	tlsConf.InsecureSkipVerify = true
	conn, err := quic.DialAddrContext(ctx, ":3000", tlsConf, &quic.Config{})
	if err != nil {
		panic(err)
	}
	stream, err := conn.OpenStream()
	if err != nil {
		panic(err)
	}
	peerCtx := stream.Context()
	for {
		select {
		case <-peerCtx.Done():
			log.Println("Gracefully shutting down an app")
			return
		default:
			n, err := stream.Write([]byte("hello world!"))
			if err != nil {
				log.Println("Failed to write a message: " + err.Error())
				continue
			}
			log.Printf("%d bytes were written\n", n)
			var p [512]byte
			if _, err = stream.Read(p[:]); err != nil {
				log.Println("Failed to read: " + err.Error())
				continue
			}
			time.Sleep(time.Second * 2)
		}
	}
}

func generateTLS() (*tls.Config, error) {
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
