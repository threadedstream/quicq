package broker

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/threadedstream/quicthing/internal/broker"
)

func Main() {
	br, ctx := setupBroker(context.Background())

	if err := br.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

func MainAsync() {
	br, ctx := setupBroker(context.Background())

	go func() {
		if err := br.Run(ctx); err != nil {
			log.Fatal(err)
		}
	}()
}

func setupBroker(ctx context.Context) (broker.Broker, context.Context) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	log.Println("starting broker...")
	ctx, cancel := context.WithCancel(ctx)
	br := broker.New()

	go func() {
		for {
			select {
			case code := <-sigChan:
				log.Println("received term code " + code.String())
				cancel()
				return
			}
		}
	}()
	return br, ctx
}
