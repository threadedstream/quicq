package main

import (
	"context"
	"github.com/threadedstream/quicthing/internal/server/tcp"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
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
			_, err := server.AcceptClient(ctx)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func client(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	if parts := strings.Split(address, ":"); len(parts) == 2 {
		host, port := parts[0], func() int { val, _ := strconv.Atoi(parts[1]); return val }()
		raddr := &net.TCPAddr{
			IP:   net.ParseIP(host),
			Port: port,
		}
		_, err := net.DialTCP("tcp4", nil, raddr)
		if err != nil {
			log.Println("could not dial tcp address: %s" + err.Error())
			return
		}
	}
	//outer:
	//	for {
	//		select {
	//		case <-time.After(time.Minute * 2):
	//			return
	//		case <-time.After(pollTimeout):
	//		}
	//	}
	//}
}
