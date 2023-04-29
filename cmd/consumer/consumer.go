package main

import (
	"context"
	"flag"
	"github.com/threadedstream/quicthing/internal/consumer"
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
	"log"
	"sync"
	"time"
)

var (
	host = flag.String("host", "localhost", "host to dial into")
	port = flag.String("port", "3000", "server port")
)

func main() {
	// parse flags
	flag.Parse()

	ctx := context.Background()
	c := consumer.New()
	if err := c.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	resp, err := c.Subscribe("feed-topic")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got response with response type %s", quicq.ResponseType_name[int32(resp.GetResponseType())])

	wg := &sync.WaitGroup{}
	ticker := time.NewTicker(time.Minute * 1)

outer:
	for {
		select {
		case <-ticker.C:
			break outer
		case <-time.After(time.Millisecond * 100):
			resp, err = c.Poll()
			if err != nil {
				log.Printf("error: %s", err.Error())
				continue outer
			}
			if resp.GetPollResponse().GetRecords() != nil {
				wg.Add(1)
				// process messages only in case they're present
				go handleRecords(wg, resp.GetPollResponse().GetRecords())
			}
			continue outer
		}
	}

	wg.Wait()
}

func handleRecords(wg *sync.WaitGroup, records []*quicq.Record) {
	defer wg.Done()
	for _, rec := range records {
		log.Printf("[%s] => %s", rec.GetKey(), rec.GetPayload())
	}
}
