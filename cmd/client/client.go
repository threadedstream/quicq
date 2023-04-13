package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/threadedstream/quicthing/internal/conn"
)

var (
	host = flag.String("host", "localhost", "host to dial into")
	port = flag.String("port", "3000", "server port")
)

func main() {
	// parse flags
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	defaultGenerateTLSFn := conn.DefaultGenerateTLSFunc()
	tlsConf, err := defaultGenerateTLSFn()
	if err != nil {
		log.Fatal(err)
	}
	tlsConf.InsecureSkipVerify = true

	addr := *host + ":" + *port
	log.Println("Dialing host " + addr)
	conn, err := quic.DialAddrContext(ctx, addr, tlsConf, &quic.Config{})
	if err != nil {
		log.Fatal(err)
	}

	stream, err := conn.OpenStream()
	if err != nil {
		conn.CloseWithError(quic.ApplicationErrorCode(quic.StreamStateError), "failed to open a stream")
		log.Fatal(err)
	}
	peerCtx := stream.Context()
commloop:
	for {
		select {
		case <-peerCtx.Done():
			log.Println("Gracefully shutting down an app")
			break commloop
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

	// gracefully terminate connection
	conn.CloseWithError(quic.ApplicationErrorCode(quic.NoError), "closed")
}
