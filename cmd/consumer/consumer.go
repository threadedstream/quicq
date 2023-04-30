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

	cancelCtx, cancel := context.WithCancel(ctx)
	dataChan, errChan := c.Notify(cancelCtx)

outer:
	for {
		select {
		case <-ticker.C:
			cancel()
			break outer
		case r := <-dataChan:
			if recs := r.GetPollResponse().GetRecords(); len(recs) > 0 {
				wg.Add(1)
				// process messages only in case they're present
				go handleRecords(wg, recs)
			}
			continue outer
		case e := <-errChan:
			log.Printf("error while polling: %s\n", e.Error())
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
