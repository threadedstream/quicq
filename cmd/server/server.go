package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/threadedstream/quicthing/internal/conn"
	"github.com/threadedstream/quicthing/internal/server"
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

func handleConnection(wg *sync.WaitGroup, conn conn.Connection) {
	// no need to use wait group for now
	//defer wg.Done()
	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return
	}
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			stream.Close()
			return
		default:
			var p [512]byte
			_, err = stream.Rcv(p[:])
			if err != nil {
				log.Printf("[%s] Failed to read a message: %s\n", conn.RemoteAddr().String(), err.Error())
				continue
			}
			log.Printf("[%s] => %s\n", conn.RemoteAddr().String(), string(p[:]))
			// write the message back
			if _, err = stream.Send(p[:]); err != nil {
				log.Printf("[%s] Failed to write a message: %s\n", conn.RemoteAddr().String(), err.Error())
				continue
			}
		}
	}
}

func main() {
	addr := "0.0.0.0:3000"
	server := &server.QuicQServer{}
	if err := server.Serve(addr); nil != err {
		log.Fatal(err)
	}
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
			}
		}
	}()

	log.Println("start accepting connections on " + addr)
	for {
		select {
		case <-ctx.Done():
			server.Shutdown()
			return
		default:
			conn, err := server.AcceptClient(ctx)
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
