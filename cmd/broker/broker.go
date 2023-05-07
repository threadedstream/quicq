package main

import (
	"context"
	"github.com/threadedstream/quicthing/internal/broker"
	_ "github.com/threadedstream/quicthing/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	log.Println("starting broker...")
	ctx, cancel := context.WithCancel(context.Background())
	br := broker.New()

	go func() {
		for {
			select {
			case code := <-sigChan:
				log.Println("received term code " + code.String())
				cancel()
				return
				//case <-ticker.C:
				//	cancelCPU := prof.MustProfCPU("cpu_broker.prof")
				//	cancelMem := prof.MustProfMem("mem_broker.prof")
				//	func() { cancelCPU(); cancelMem() }()
				//	ticker.Reset(tickerDuration)
				//	continue outer
			}
		}
	}()

	if err := br.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
