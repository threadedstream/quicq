package producer

import (
	"context"
	"log"

	"github.com/threadedstream/quicthing/internal/publisher"
)

func Main() {
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
