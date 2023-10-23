package publisher

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/threadedstream/quicthing/internal/client"
	"github.com/threadedstream/quicthing/internal/common"
	"github.com/threadedstream/quicthing/internal/config"
	"github.com/threadedstream/quicthing/internal/encoder"
	protoenc "github.com/threadedstream/quicthing/internal/encoder/protobuf"
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Publisher interface {
	Post(topic string, message common.Message) (*quicq.Response, error)
	PostBulk(topic string, messages []common.Message) (*quicq.Response, error)
}

// QuicQProducer is a Consumer implementation
type QuicQProducer struct {
	client  *client.QuicQClient
	encoder encoder.Encoder
	decoder encoder.Decoder
}

// New initializes a QuicQProducer object
func New() *QuicQProducer {
	return &QuicQProducer{
		encoder: protoenc.NewProtoEncoder(),
		decoder: protoenc.NewProtoDecoder(),
	}
}

// Connect connects to broker
func (qp *QuicQProducer) Connect(ctx context.Context) error {
	return qp.connect(ctx)
}

func (qp *QuicQProducer) connect(ctx context.Context) error {
	qp.client = client.New()
	if err := qp.client.Dial(ctx, config.BrokerConfig.BrokerAddr()); err != nil {
		return err
	}
	return nil
}

func (qp *QuicQProducer) Post(topic string, message common.Message) (*quicq.Response, error) {
	req := &quicq.Request{
		RequestType: quicq.RequestType_REQUEST_POST,
		Request: &quicq.Request_PostRequest{
			PostRequest: &quicq.PostRequest{
				Topic: topic,
				Record: &quicq.Record{
					Key:       message.Key,
					Payload:   message.Payload,
					Timestamp: timestamppb.New(time.Now()),
				},
			},
		},
	}
	return qp.do(req)
}

// PostBulk - the slow version
func (qp *QuicQProducer) PostBulk(topic string, messages []common.Message) (*quicq.Response, error) {
	var records []*quicq.Record
	for _, m := range messages {
		records = append(records, &quicq.Record{
			Key:       m.Key,
			Payload:   m.Payload,
			Timestamp: timestamppb.New(time.Now()),
		})
	}
	req := &quicq.Request{
		RequestType: quicq.RequestType_REQUEST_POST_BULK,
		Request: &quicq.Request_PostBulkRequest{
			PostBulkRequest: &quicq.PostBulkRequest{
				Topic:   topic,
				Records: records,
			},
		},
	}
	return qp.do(req)
}

func (qp *QuicQProducer) do(req *quicq.Request) (*quicq.Response, error) {
	bs, err := qp.encoder.EncodeRequest(req)
	if err != nil {
		log.Println("failed to encode request: ", err)
		return nil, err
	}

	stream, err := qp.client.RequestStream()
	if err != nil {
		log.Println("failed to request stream: ", err)
		return nil, err
	}

	defer func() {
		_ = stream.Shutdown()
	}()

	var bytesSent int
	if bytesSent, err = stream.Send(bs); err != nil {
		log.Println("failed to send stream of bytes over a stream: ", err)
		return nil, err
	}

	if bytesSent != len(bs) {
		log.Println("bytesSent != len(bs)")
	}

	log.Println("sent exactly ", bytesSent, " bytes")

	var responseBytes []byte
	if responseBytes, err = stream.Rcv(); err != nil {
		log.Println("failed to receive stream of bytes from stream: ", err)
		return nil, err
	}

	log.Println("received exactly ", len(responseBytes), " bytes")

	resp, err := qp.decoder.DecodeResponse(responseBytes)
	if err != nil {
		log.Println("failed to decode response: ", err)
		return nil, err
	}
	if resp.ResponseType == quicq.ResponseType_RESPONSE_ERROR {
		return nil, errors.New(resp.GetErrResponse().GetDetails())
	}
	return resp, nil
}
