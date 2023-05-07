package main

import (
	"context"
	"log"

	"github.com/threadedstream/quicthing/internal/prof"
	"github.com/threadedstream/quicthing/internal/publisher"
)

func main() {
	cancelCPU := prof.MustProfCPU("cpu_producer.prof")
	cancelMem := prof.MustProfMem("mem_producer.prof")

	defer func() { cancelCPU(); cancelMem() }()

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
