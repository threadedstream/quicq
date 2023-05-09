package main

import (
	"context"
	"github.com/threadedstream/quicthing/internal/publisher"
	"log"
)

func main() {
	ctx := context.Background()
	p := publisher.New()
	if err := p.Connect(ctx); err != nil {
		log.Fatalf("failed to connect to server: %s", err.Error())
	}

	for {
		_, err := p.Post("feed-topic", []byte("mykey"), []byte("my payload"))
		if err != nil {
			log.Printf("error: %s", err.Error())
		}
	}
}
