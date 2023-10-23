package consumer

import (
	"context"
	"log"
	"sync"

	"github.com/threadedstream/quicthing/internal/consumer"
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
)

func Main() {
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

	for {
		resp, err := c.Poll()
		if err != nil {
			log.Printf("poll error: %s", err)
			continue
		}
		wg.Add(1)
		go handleRecords(wg, resp.GetPollResponse().GetRecords())
	}
}

func handleRecords(wg *sync.WaitGroup, records []*quicq.Record) {
	defer wg.Done()
	for _, rec := range records {
		log.Printf("[%s] => %s", rec.GetKey(), rec.GetPayload())
	}
}
