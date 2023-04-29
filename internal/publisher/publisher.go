package publisher

import (
	"bytes"
	"context"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/threadedstream/quicthing/internal/client"
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Publisher interface {
	Post(topic string, key, payload []byte) (*quicq.Response, error)
}

// QuicQProducer is a Consumer implementation
type QuicQProducer struct {
	client *client.QuicQClient
}

// New initializes a QuicQProducer object
func New() *QuicQProducer {
	return &QuicQProducer{}
}

// Connect connects to broker
func (qp *QuicQProducer) Connect(ctx context.Context) error {
	return qp.connect(ctx)
}

func (qp *QuicQProducer) connect(ctx context.Context) error {
	const brokerAddr = "127.0.0.1:9999"
	qp.client = client.New()
	if err := qp.client.Dial(ctx, brokerAddr); err != nil {
		return err
	}
	return nil
}

func (qp *QuicQProducer) Post(topic string, key, payload []byte) (*quicq.Response, error) {
	req := &quicq.Request{
		RequestType: quicq.RequestType_REQUEST_POST,
		Request: &quicq.Request_PostRequest{
			PostRequest: &quicq.PostRequest{
				Topic: topic,
				Record: &quicq.Record{
					Key:       key,
					Payload:   payload,
					Timestamp: timestamppb.New(time.Now()),
				},
			},
		},
	}
	return qp.do(req)
}

func (qp *QuicQProducer) do(req *quicq.Request) (*quicq.Response, error) {
	bs, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	bs = append(bs, '\xee')

	stream, err := qp.client.RequestStream()
	if err != nil {
		return nil, err
	}
	if _, err = stream.Send(bs); err != nil {
		return nil, err
	}

	var responseBytes [1024]byte
	if _, err = stream.Rcv(responseBytes[:]); err != nil {
		return nil, err
	}

	rb := bytes.Trim(responseBytes[:], "\x00")
	resp := new(quicq.Response)
	if err = proto.Unmarshal(rb, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
