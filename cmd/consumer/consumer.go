package main

import (
	"context"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/threadedstream/quicthing/internal/consumer"
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
)

const (
	messageGoal int64 = 5000
)

var (
	processedMessages atomic.Int64
)

func main() {
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
	ticker := time.NewTicker(time.Minute * 10)

	cancelCtx, cancel := context.WithCancel(ctx)
	dataChan, errChan := c.Notify(cancelCtx)

	go func() {
		for {
			time.Sleep(10 * time.Millisecond)
			if processedMessages.Load() >= messageGoal {
				cancel()
			}
			println(processedMessages.Load())
		}
	}()

outer:
	for {
		select {
		case <-cancelCtx.Done():
			println("processed all messages")
			break outer
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
	processedMessages.Add(int64(len(records)))
}
