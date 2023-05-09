package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientTCP "github.com/threadedstream/quicthing/internal/client/tcp"
	"github.com/threadedstream/quicthing/internal/conn"
	"github.com/threadedstream/quicthing/internal/server/tcp"
)

const (
	pollTimeout = 100 * time.Millisecond
	address     = "127.0.0.1:3000"
)

func main() {
	server := &tcp.QuicQTCPServer{}
	if err := server.Serve(address); err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	go client(ctx, cancel)
	for {
		select {
		case <-ctx.Done():
			println("shutting down")
			server.Shutdown()
			return
		case <-time.After(pollTimeout):
			c, err := server.AcceptClient(ctx)
			if err != nil {
				log.Fatal(err)
			}
			go handleConnection(ctx, c)
		}
	}
}

func handleConnection(ctx context.Context, conn conn.Connection) {
outer:
	for {
		select {
		case <-ctx.Done():
			log.Println("context done:" + ctx.Err().Error())
			return
		case <-time.After(pollTimeout):
			stream, err := conn.AcceptStream(ctx)
			if err != nil {
				log.Println("could not accept stream: " + err.Error())
				continue outer
			}
			go handleStream(ctx, stream)
			time.Sleep(1 * time.Second)
		}
	}
}

func handleStream(ctx context.Context, stream conn.Stream) {
	var p [512]byte
	_, err := stream.Rcv(p[:])
	if err != nil {
		return
	}
	log.Println("client message: " + string(p[:]))
	stream.Shutdown()
	return
}

func client(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	cli, err := clientTCP.Dial(ctx, address)
	if err != nil {
		log.Println("could not dial tcp address: %s" + err.Error())
		return
	}

	var x int
outer:
	for {
		select {
		case <-time.After(time.Minute * 2):
			return
		case <-time.After(pollTimeout):
			stream, e := cli.RequestStream()
			if e != nil {
				log.Println("RequestStream error: " + e.Error())
				continue outer
			}
			stream.Send([]byte(fmt.Sprintf("hey there %d", x)))
			x++
		}
	}
}
