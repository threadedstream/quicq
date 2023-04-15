package main

import (
	"context"
	"flag"
	"log"

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

	resp, err = c.Subscribe("news-feed")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got response with response type %s", quicq.ResponseType_name[int32(resp.GetResponseType())])

	resp, err = c.Unsubscribe("news-feed")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got response with response type %s", quicq.ResponseType_name[int32(resp.GetResponseType())])

	if resp, err = c.FetchTopicMetadata(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Got response with response type %s", quicq.ResponseType_name[int32(resp.GetResponseType())])

	fetchMdResp := resp.GetFetchTopicMetadataResponse()
	log.Printf("Topics: %v", fetchMdResp.GetTopics())
}
