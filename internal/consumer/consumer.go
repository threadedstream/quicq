package consumer

import (
	"bytes"
	"context"

	"github.com/threadedstream/quicthing/internal/client"
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
	"google.golang.org/protobuf/proto"
)

// Consumer is a consumer interface
type Consumer interface {
	Run(ctx context.Context) error
	Subscribe(topic string) error
	Unsubscribe(topic string) error
}

// QuicQConsumer is a Consumer implementation
type QuicQConsumer struct {
	id     int64
	client *client.QuicQClient
}

// New initializes a QuicQConsumer object
func New() *QuicQConsumer {
	return &QuicQConsumer{}
}

// Run runs consumer
func (qc *QuicQConsumer) Run(ctx context.Context) error {
	return qc.run(ctx)
}

func (qc *QuicQConsumer) run(ctx context.Context) error {
	const brokerAddr = "0.0.0.0:9999"
	qc.client = client.New()
	if err := qc.client.Dial(ctx, brokerAddr); err != nil {
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
