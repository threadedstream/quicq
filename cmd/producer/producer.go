package main

import (
	"context"
	"flag"
	"log"

	"github.com/threadedstream/quicthing/internal/publisher"
)

func main() {
	// parse flags
	flag.Parse()

	ctx := context.Background()
	p := publisher.New()
	if err := p.Connect(ctx); err != nil {
		log.Fatalf("failed to connect to server: %s", err.Error())
	}

	_, err := p.Post("news-feed", []byte("mykey"), []byte("my payload"))
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
}
