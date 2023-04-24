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

	"golang.org/x/sys/unix"
)

func listenTLS() {
	tlsConf, err := getTLS()
	if err != nil {
		log.Fatal(err)
	}
	server, err := tls.Listen("tcp4", ":3000", tlsConf)
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

	log.Println("listening on port 3000")
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Printf("err: %s", err.Error())
			return
		}
		conn.Write([]byte("dummy"))
	}
}

func listenNoTLS() {
	config := &net.ListenConfig{Control: reusePort}

	server, err := config.Listen(context.Background(), "tcp4", ":3000")
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

	log.Println("listening on port 3000")
	for {
		_, err := server.Accept()
		if err != nil {
			log.Printf("err: %s", err.Error())
			return
		}
	}
}

func main() {
	listenTLS()
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
