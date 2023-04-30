package main

import (
	"context"
	"log"

	"github.com/threadedstream/quicthing/internal/broker"
	_ "github.com/threadedstream/quicthing/internal/config"
)

func main() {
	log.Println("starting broker...")
	ctx := context.Background()
	br := broker.New()
	if err := br.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
