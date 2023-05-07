package main

import (
	"context"
	"log"

	"github.com/threadedstream/quicthing/internal/broker"
	_ "github.com/threadedstream/quicthing/internal/config"
	"github.com/threadedstream/quicthing/internal/prof"
)

func main() {
	cancelCPU := prof.MustProfCPU("cpu_broker.prof")
	cancelMem := prof.MustProfMem("mem_broker.prof")

	defer func() { cancelCPU(); cancelMem() }()

	log.Println("starting broker...")
	ctx := context.Background()
	br := broker.New()
	if err := br.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
