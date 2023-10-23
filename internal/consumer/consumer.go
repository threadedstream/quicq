package consumer

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
)

// Consumer is a consumer interface
type Consumer interface {
	Connect(ctx context.Context) error
	Subscribe(topic string) (*quicq.Response, error)
	Unsubscribe(topic string) (*quicq.Response, error)
	FetchTopicMetadata() (*quicq.Response, error)
	Poll() (*quicq.Response, error)
}

// QuicQConsumer is a Consumer implementation
type QuicQConsumer struct {
	id      int64
	client  *client.QuicQClient
	encoder encoder.Encoder
	decoder encoder.Decoder
}

// New initializes a QuicQConsumer object
func New() *QuicQConsumer {
	return &QuicQConsumer{
		id:      int64(common.IDManager.GetNextID()),
		encoder: protoenc.NewProtoEncoder(),
		decoder: protoenc.NewProtoDecoder(),
	}
}

// Connect connects to broker
func (qc *QuicQConsumer) Connect(ctx context.Context) error {
	return qc.connect(ctx)
}

func (qc *QuicQConsumer) connect(ctx context.Context) error {
	qc.client = client.New()
	if err := qc.client.Dial(ctx, config.BrokerConfig.BrokerAddr()); err != nil {
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

func (qc *QuicQConsumer) Poll() (*quicq.Response, error) {
	req := &quicq.Request{
		RequestType: quicq.RequestType_REQUEST_POLL,
		Request: &quicq.Request_PollRequest{
			PollRequest: &quicq.PollRequest{
				ConsumerID: qc.id,
			},
		},
	}

	return qc.do(req)
}

// Notify is not really a request, but an attempt to set up push model
func (qc *QuicQConsumer) Notify(ctx context.Context, pollPeriod time.Duration) (chan *quicq.Response, chan error) {
	dataChan := make(chan *quicq.Response, common.QueueSizeMax)
	errChan := make(chan error, common.QueueSizeMax)

	go func() {
	outer:
		for {
			select {
			case <-ctx.Done():
				close(dataChan)
				close(errChan)
				// context canceled
				return
			case <-time.After(pollPeriod):
				resp, err := qc.Poll()
				if err != nil {
					errChan <- err
					continue outer
				}
				dataChan <- resp
				continue outer
			}
		}
	}()

	return dataChan, errChan
}

func (qc *QuicQConsumer) do(req *quicq.Request) (*quicq.Response, error) {
	bs, err := qc.encoder.EncodeRequest(req)
	if err != nil {
		log.Println("failed to encode request: ", err)
		return nil, err
	}

	stream, err := qc.client.RequestStream()
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
		return nil, err
	}

	log.Println("received exactly ", len(responseBytes), " bytes")

	resp, err := qc.decoder.DecodeResponse(responseBytes)
	if err != nil {
		log.Println("failed to decode response: ", err)
		return nil, err
	}

	if resp.ResponseType == quicq.ResponseType_RESPONSE_ERROR {
		return nil, errors.New(resp.GetErrResponse().GetDetails())
	}
	return resp, nil
}
