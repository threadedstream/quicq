package producer

import (
	"context"
	"fmt"
	"log"

	"github.com/threadedstream/quicthing/internal/common"
	"github.com/threadedstream/quicthing/internal/publisher"
)

func Main() {
	ctx := context.Background()
	p := publisher.New()
	if err := p.Connect(ctx); err != nil {
		log.Fatalf("failed to connect to server: %s", err.Error())
	}

	var messages []common.Message
	for i := 1; i <= 15; i++ {
		messages = append(messages, common.Message{
			Key:     []byte(fmt.Sprintf("key%d", i)),
			Payload: []byte(fmt.Sprintf("payload%d", i)),
		})
	}

	for {
		_, err := p.PostBulk("feed-topic", messages)
		if err != nil {
			log.Printf("error: %s", err.Error())
		}
	}
}
