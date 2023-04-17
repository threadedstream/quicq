package main

import (
	"context"
	"flag"
	"github.com/threadedstream/quicthing/internal/publisher"
	"log"
)

var (
	host = flag.String("host", "localhost", "host to dial into")
	port = flag.String("port", "3000", "server port")
)

func main() {
	// parse flags
	flag.Parse()

	ctx := context.Background()
	p := publisher.New()
	if err := p.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	_, err := p.Post("news-feed", []byte("mykey"), []byte("my payload"))
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
}
