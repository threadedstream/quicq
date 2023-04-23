package main

import (
	"log"
	"net"
	"time"
)

func testConnect() {
	localAddr := &net.TCPAddr{
		Port: 45000,
	}
	remoteAddr := &net.TCPAddr{
		Port: 3000,
	}
	conn, err := net.DialTCP("tcp4", localAddr, remoteAddr)
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()
}

func main() {
	// 10 connections
	start := time.Now()
	for i := 0; i < 1e4; i++ {
		testConnect()
	}
	end := time.Now().Sub(start).Seconds()
	log.Printf("10 conn: %f\n", end)

	// 100 connections
	start = time.Now()
	for i := 0; i < 1e7; i++ {
		testConnect()
	}
	end = time.Now().Sub(start).Seconds()
	log.Printf("100 conn: %f\n", end)

	// 1000 connections
	for i := 0; i < 1e10; i++ {
		testConnect()
	}
	end = time.Now().Sub(start).Seconds()
	log.Printf("1000 conn: %f\n", end)
}
