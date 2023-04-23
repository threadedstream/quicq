package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	addr := &net.TCPAddr{
		Port: 3000,
	}
	server, err := net.ListenTCP("tcp4", addr)
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
		_, err := server.AcceptTCP()
		if err != nil {
			log.Println("err: %s", err.Error())
			return
		}
	}
}
