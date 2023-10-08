package tests

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/threadedstream/quicthing/internal/cmd/broker"
	"github.com/threadedstream/quicthing/internal/consumer"
	"github.com/threadedstream/quicthing/internal/publisher"
)

func init() {
	// run broker asynchronously
	broker.MainAsync()
}

func BenchmarkPerformanceTest(b *testing.B) {

	ctx := context.Background()
	const topicName = "feed-topic"
	p := setupPublisher(ctx)
	c := setupConsumer(ctx, topicName)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := p.Post(topicName, []byte("key"), []byte("value"))
		require.NoError(b, err)
		resp, err := c.Poll()
		require.NoError(b, err)
		require.Greater(b, len(resp.GetPollResponse().GetRecords()), 0)
	}
}

func setupPublisher(ctx context.Context) publisher.Publisher {
	p := publisher.New()
	if err := p.Connect(ctx); err != nil {
		log.Fatalf("failed to connect to server: %s", err.Error())
	}
	return p
}

func setupConsumer(ctx context.Context, topicName string) consumer.Consumer {
	c := consumer.New()
	if err := c.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	_, err := c.Subscribe(topicName)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
