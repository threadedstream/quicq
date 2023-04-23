package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
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

func reusePort(network, address string, conn syscall.RawConn) error {
	return conn.Control(func(descriptor uintptr) {
		syscall.SetsockoptInt(int(descriptor), syscall.SOL_SOCKET, syscall.SO_REUSEPORT, 1)
	})
}
