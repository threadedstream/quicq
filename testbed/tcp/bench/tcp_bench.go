package main

import (
	"log"
	"net"
	"syscall"
	"time"
)

// #include <stdlib.h>
import "C"

func testConnect() {
	dialer := net.Dialer{
		Control: reusePort,
		LocalAddr: &net.TCPAddr{
			Port: 45000,
		},
	}
	conn, err := dialer.Dial("tcp4", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Close()
	if err != nil {
		log.Fatalf("close err: %s", err.Error())
	}
}

func main() {
	// 100 connections
	start := time.Now()
	for i := 0; i < 1e2; i++ {
		testConnect()
	}
	end := time.Now().Sub(start).Seconds()
	log.Printf("10 conn: %f\n", end)

	// 1000 connections
	start = time.Now()
	for i := 0; i < 1e3; i++ {
		testConnect()
	}
	end = time.Now().Sub(start).Seconds()
	log.Printf("100 conn: %f\n", end)

	// 10000 connections
	for i := 0; i < 1e4; i++ {
		testConnect()
	}
	end = time.Now().Sub(start).Seconds()
	log.Printf("1000 conn: %f\n", end)
}

func reusePort(network, address string, conn syscall.RawConn) error {
	return conn.Control(func(descriptor uintptr) {
		syscall.SetsockoptInt(int(descriptor), syscall.SOL_SOCKET, syscall.SO_REUSEPORT, 1)
	})
}
