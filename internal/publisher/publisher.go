package publisher

import (
	"context"
	"errors"
	"github.com/threadedstream/quicthing/internal/client/tcp"
	"time"

	"github.com/threadedstream/quicthing/internal/config"
	"github.com/threadedstream/quicthing/internal/encoder"

	"github.com/threadedstream/quicthing/internal/client"
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Publisher interface {
	Post(topic string, key, payload []byte) (*quicq.Response, error)
}

// QuicQProducer is a Consumer implementation
type QuicQProducer struct {
	client  client.Client
	encoder encoder.Encoder
	decoder encoder.Decoder
}

// New initializes a QuicQProducer object
func New() *QuicQProducer {
	return &QuicQProducer{
		encoder: encoder.NewProtoEncoder(),
		decoder: encoder.NewProtoDecoder(),
	}
}

// Connect connects to broker
func (qp *QuicQProducer) Connect(ctx context.Context) error {
	return qp.connect(ctx)
}

func (qp *QuicQProducer) connect(ctx context.Context) error {
	var err error
	if qp.client, err = tcp.Dial(ctx, config.BrokerConfig.BrokerAddr()); err != nil {
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
	bs, err := qp.encoder.EncodeRequest(req)
	if err != nil {
		return nil, err
	}

	stream, err := qp.client.RequestStream()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = stream.Shutdown()
	}()

	var n int
	if n, err = stream.Send(bs); err != nil {
		return nil, err
	}
	println(n)

	var responseBytes [1024]byte
	if _, err = stream.Rcv(responseBytes[:]); err != nil {
		return nil, err
	}

	resp, err := qp.decoder.DecodeResponse(responseBytes[:])
	if err != nil {
		return nil, err
	}
	if resp.ResponseType == quicq.ResponseType_RESPONSE_ERROR {
		return nil, errors.New(resp.GetErrResponse().GetDetails())
	}
	return resp, nil
}
