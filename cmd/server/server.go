package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/quic-go/quic-go"
)

var (
	onShutdownCallbacks []func()
)

func getContents(path string) ([]byte, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(file)
}

// loadTLSConfig: borrowed from ngrok for a time being
func loadTLSConfig(rootCertPaths ...string) (*tls.Config, error) {
	pool := x509.NewCertPool()

	for _, certPath := range rootCertPaths {
		rootCrt, err := getContents(certPath)
		if err != nil {
			return nil, err
		}

		pemBlock, _ := pem.Decode(rootCrt)
		if pemBlock == nil {
			return nil, fmt.Errorf("Bad PEM data")
		}

		certs, err := x509.ParseCertificates(pemBlock.Bytes)
		if err != nil {
			return nil, err
		}

		pool.AddCert(certs[0])
	}

	return &tls.Config{RootCAs: pool}, nil
}

func handleConnection(wg *sync.WaitGroup, conn quic.Connection) {
	// no need to use wait group for now
	//defer wg.Done()
	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return
	}
	log.Printf("accepted a stream with id = %d\n", stream.StreamID().StreamNum())
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			stream.Close()
			return
		default:
			var p [512]byte
			_, err = stream.Read(p[:])
			if err != nil {
				log.Println("Failed to read a message: " + err.Error())
				continue
			}
			log.Printf("[%s] => %s\n", conn.RemoteAddr().String(), string(p[:]))
			// write the message back
			if _, err = stream.Write(p[:]); err != nil {
				log.Println("Failed to write a message: " + err.Error())
				continue
			}
		}
	}
}

func main() {
	tlsConf, err := generateTLS()
	if err != nil {
		log.Fatal(err)

	}
	quicConf := quic.Config{}
	listener, err := quic.ListenAddr(":3000", tlsConf, &quicConf)
	if err != nil {
		panic(err)
	}
	onShutdownCallbacks = append(onShutdownCallbacks, func() {
		listener.Close()
	})
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case t := <-sigChan:
				cancel()
				log.Printf("caught signal %d\n", t)
				for _, fn := range onShutdownCallbacks {
					fn()
				}
			}
		}
	}()

	log.Println("start accepting connections")
	for {
		select {
		case <-ctx.Done():
			listener.Close()
			return
		default:
			conn, err := listener.Accept(ctx)
			if err != nil {
				// it's highly discouraged, but we're good w/ that for the purpose of learning
				log.Println("failed to accept: " + err.Error())
				continue
			}
			log.Printf("got a new connection: %s\n", conn.RemoteAddr().String())
			go handleConnection(nil, conn)
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
