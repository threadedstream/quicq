package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/threadedstream/quicthing/internal/consumer"
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
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

	for {
		time.Sleep(2 * time.Second)
		resp, err = c.Poll()
		if err != nil {
			log.Printf("error: %s", err.Error())
			continue
		}
		go handleRecords(resp.GetPollResponse().GetRecords())
	}
}

func handleRecords(records []*quicq.Record) {
	for _, rec := range records {
		log.Printf("> [%s] => %s", rec.GetKey(), rec.GetPayload())
	}
}
