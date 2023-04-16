package consumer

import (
	"bytes"
	"context"
	"math/rand"

	"github.com/threadedstream/quicthing/internal/client"
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
	"google.golang.org/protobuf/proto"
)

// Consumer is a consumer interface
type Consumer interface {
	Connect(ctx context.Context) error
	Subscribe(topic string) (*quicq.Response, error)
	Unsubscribe(topic string) (*quicq.Response, error)
	FetchTopicMetadata() (*quicq.Response, error)
}

// QuicQConsumer is a Consumer implementation
type QuicQConsumer struct {
	id     int64
	client *client.QuicQClient
}

// New initializes a QuicQConsumer object
func New() *QuicQConsumer {
	return &QuicQConsumer{
		id: rand.Int63(),
	}
}

// Connect connects to broker
func (qc *QuicQConsumer) Connect(ctx context.Context) error {
	return qc.connect(ctx)
}

func (qc *QuicQConsumer) connect(ctx context.Context) error {
	const brokerAddr = "0.0.0.0:9999"
	qc.client = client.New()
	if err := qc.client.Dial(ctx, brokerAddr); err != nil {
		return err
	}
	if err := qc.client.RequestStream(); err != nil {
		return err
	}
	return nil
}

// Subscribe subscribes to particular topic
func (qc *QuicQConsumer) Subscribe(topic string) (*quicq.Response, error) {
	req := &quicq.Request{
		RequestType: quicq.RequestType_REQUEST_SUBSCRIBE,
		Request: &quicq.Request_SubscribeRequest{
			SubscribeRequest: &quicq.SubscribeRequest{
				Topic:      topic,
				ConsumerID: qc.id,
			},
		},
	}
	return qc.do(req)
}

// Unsubscribe unsubscribes from particular topic
func (qc *QuicQConsumer) Unsubscribe(topic string) (*quicq.Response, error) {
	req := &quicq.Request{
		RequestType: quicq.RequestType_REQUEST_UNSUBSCRIBE,
		Request: &quicq.Request_UnsubscribeRequest{
			UnsubscribeRequest: &quicq.UnsubscribeRequest{
				Topic:      topic,
				ConsumerID: qc.id,
			},
		},
	}
	return qc.do(req)
}

func (qc *QuicQConsumer) FetchTopicMetadata() (*quicq.Response, error) {
	req := &quicq.Request{
		RequestType: quicq.RequestType_REQUEST_FETCH_TOPIC_METADATA,
		Request: &quicq.Request_FetchTopicMetadataRequest{
			FetchTopicMetadataRequest: &quicq.FetchTopicMetadataRequest{
				ConsumerID: qc.id,
			},
		},
	}
	return qc.do(req)
}

func (qc *QuicQConsumer) do(req *quicq.Request) (*quicq.Response, error) {
	bs, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	if err = qc.client.Send(bs); err != nil {
		return nil, err
	}

	var responseBytes [1024]byte
	if err = qc.client.Rcv(responseBytes[:]); err != nil {
		return nil, err
	}

	rb := bytes.Trim(responseBytes[:], "\x00")
	resp := new(quicq.Response)
	if err = proto.Unmarshal(rb, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
