package broker

import (
	"context"
	"errors"
	"fmt"
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
	"io"
	"log"
	"sync"
	"time"

	errors2 "github.com/onsi/gomega/gstruct/errors"
	"github.com/threadedstream/quicthing/internal/config"
	"github.com/threadedstream/quicthing/internal/conn"
	"github.com/threadedstream/quicthing/internal/encoder"
	"github.com/threadedstream/quicthing/internal/server"
	serverTCP "github.com/threadedstream/quicthing/internal/server/tcp"
	"github.com/threadedstream/quicthing/internal/topic"
)

const (
	ctxPollTimeout = 100 * time.Millisecond
)

const (
	errOccurredFmt = "error cause: %s"
)

var (
	brokerClosedErr = errors.New("closed")
)

type Broker interface {
	Run(context.Context) error
}

type QuicQBroker struct {
	server        server.Server
	mu            *sync.Mutex
	subscriptions map[int64][]string
	topics        []topic.Topic
	topicQueryMap map[string]topic.Topic
	encoder       encoder.Encoder
	decoder       encoder.Decoder
}

func New() *QuicQBroker {
	broker := &QuicQBroker{
		server:        new(serverTCP.QuicQTCPServer),
		mu:            new(sync.Mutex),
		topics:        make([]topic.Topic, 0),
		subscriptions: make(map[int64][]string),
		topicQueryMap: make(map[string]topic.Topic),
		encoder:       encoder.NewProtoEncoder(),
		decoder:       encoder.NewProtoDecoder(),
	}
	return broker
}

func (qb *QuicQBroker) Run(ctx context.Context) error {
	return qb.run(ctx)
}

func (qb *QuicQBroker) Shutdown() error {
	return qb.server.Shutdown()
}

func (qb *QuicQBroker) run(ctx context.Context) error {
	if e := qb.server.Serve(config.BrokerConfig.BrokerAddr()); e != nil {
		log.Fatalf("unable to serve: %s", e.Error())
	}
loop:
	for {
		select {
		case <-ctx.Done():
			qb.server.Shutdown()
			return brokerClosedErr
		case <-time.After(ctxPollTimeout):
			c, err := qb.server.AcceptClient(ctx)
			if err != nil {
				// it's highly discouraged, but we're good w/ that for the purpose of learning
				log.Println("failed to accept: " + err.Error())
				continue loop
			}
			log.Printf("got a new connection: %s\n", c.RemoteAddr().String())
			go qb.handleConnection(ctx, c)
		}
	}
}

func (qb *QuicQBroker) handleConnection(ctx context.Context, conn conn.Connection) {
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case <-time.After(ctxPollTimeout):
			stream, err := conn.AcceptStream(ctx)
			if err != nil {
				log.Println("ERROR: ", err.Error())
				return
			}
			go qb.handleStream(stream)
		}
	}
}

func (qb *QuicQBroker) handleStream(stream conn.Stream) {
	streamCtx := stream.Context()
outer:
	for {
		select {
		case _, ok := <-streamCtx.Done():
			if !ok {
				println("closing stream...")
				return
			}
		case <-time.After(ctxPollTimeout):
			var p [1024]byte
			_, err := stream.Rcv(p[:])
			// TODO(threadedstream): tcp connection "streams" are not context-aware, so return from goroutine
			// once EOF is received
			if err != nil {
				if err == io.EOF {
					println("closing stream...")
					return
				}
				stream.Log("Failed to read a message: %s\n", err.Error())
				continue outer
			}
			log.Println("got some data in")
			req, err := qb.decoder.DecodeRequest(p[:])
			if err != nil {
				qb.sendErr(stream, fmt.Sprintf(errOccurredFmt, err.Error()))
				continue outer
			}

			resp, err := qb.executeRequest(streamCtx, req)
			if err != nil {
				qb.sendErr(stream, fmt.Sprintf(errOccurredFmt, err.Error()))
				continue outer
			}
			bs, err := qb.encoder.EncodeResponse(resp)
			if err != nil {
				log.Println("failed to encode response: ", err.Error())
				qb.sendErr(stream, fmt.Sprintf(errOccurredFmt, "internal server error"))
				continue outer
			}
			// write the message back
			if _, err = stream.Send(bs); err != nil {
				stream.Log("Failed to write a message: %s\n", err.Error())
				continue outer
			}
		}
	}
}

func (qb *QuicQBroker) executeRequest(ctx context.Context, req *quicq.Request) (*quicq.Response, error) {
	switch req.GetRequestType() {
	default:
		return nil, errors.New("unknown request")
	case quicq.RequestType_REQUEST_SUBSCRIBE:
		return qb.doSubscribe(req.GetSubscribeRequest())
	case quicq.RequestType_REQUEST_UNSUBSCRIBE:
		return qb.doUnsubscribe(req.GetUnsubscribeRequest())
	case quicq.RequestType_REQUEST_FETCH_TOPIC_METADATA:
		return qb.doFetchTopicMetadata(req.GetFetchTopicMetadataRequest())
	case quicq.RequestType_REQUEST_POLL:
		return qb.doPoll(req.GetPollRequest())
	case quicq.RequestType_REQUEST_POST:
		return qb.doPost(req.GetPostRequest())
	}
}

func (qb *QuicQBroker) doPost(req *quicq.PostRequest) (*quicq.Response, error) {
	// TODO(threadedstream): define different type of mutex for locking a queue
	qb.mu.Lock()
	defer qb.mu.Unlock()
	top := qb.getOrAddTopic(req.GetTopic())
	err := top.Push(req.GetRecord())
	if err != nil {
		return nil, err
	}
	return &quicq.Response{
		ResponseType: quicq.ResponseType_RESPONSE_POST,
		Response: &quicq.Response_PostResponse{
			PostResponse: &quicq.PostResponse{
				Offset: 10,
			},
		},
	}, nil
}

func (qb *QuicQBroker) doPoll(req *quicq.PollRequest) (*quicq.Response, error) {
	var records []*quicq.Record
	topicNames := qb.subscriptions[req.GetConsumerID()]

	var topics []topic.Topic
	for _, topicName := range topicNames {
		if top, ok := qb.topicQueryMap[topicName]; ok {
			topics = append(topics, top)
		}
	}

	var errs errors2.AggregateError
	for _, t := range topics {
		recs, err := t.GetRecordBatch()
		if err != nil {
			err = fmt.Errorf("<%s> = %s", t.Name(), err.Error())
			errs = append(errs, err)
			continue
		}
		records = append(records, recs...)
	}

	if errs != nil {
		return nil, errors.New(errs.Error())
	}

	return &quicq.Response{
		ResponseType: quicq.ResponseType_RESPONSE_POLL,
		Response: &quicq.Response_PollResponse{
			PollResponse: &quicq.PollResponse{
				Records: records,
			},
		},
	}, nil
}

func (qb *QuicQBroker) doFetchTopicMetadata(req *quicq.FetchTopicMetadataRequest) (*quicq.Response, error) {
	qb.mu.Lock()
	defer qb.mu.Unlock()
	return &quicq.Response{
		ResponseType: quicq.ResponseType_RESPONSE_FETCH_TOPIC_METADATA,
		Response: &quicq.Response_FetchTopicMetadataResponse{
			FetchTopicMetadataResponse: &quicq.FetchTopicMetadataResponse{
				Topics: qb.subscriptions[req.GetConsumerID()],
			},
		},
	}, nil
}

func (qb *QuicQBroker) doSubscribe(req *quicq.SubscribeRequest) (*quicq.Response, error) {
	qb.mu.Lock()
	defer qb.mu.Unlock()
	top := qb.getOrAddTopic(req.GetTopic())
	_ = top.TieConsumer(req.GetConsumerID())
	qb.subscriptions[req.GetConsumerID()] = append(qb.subscriptions[req.GetConsumerID()], req.GetTopic())
	return &quicq.Response{
		ResponseType: quicq.ResponseType_RESPONSE_SUBSCRIBE,
		Response: &quicq.Response_SubscribeResponse{
			SubscribeResponse: &quicq.SubscribeResponse{
				Topic: req.GetTopic(),
			},
		},
	}, nil
}

func (qb *QuicQBroker) doUnsubscribe(req *quicq.UnsubscribeRequest) (*quicq.Response, error) {
	qb.mu.Lock()
	defer qb.mu.Unlock()
	for _, t := range qb.topics {
		if t.Name() == req.GetTopic() {
			_ = t.EvictConsumer(req.GetConsumerID())
		}
	}

	topicNames := qb.subscriptions[req.GetConsumerID()]
	for i := range topicNames {
		if req.GetTopic() == topicNames[i] {
			topicNames[i] = topicNames[len(topicNames)-1]
			topicNames = topicNames[:len(topicNames)-1]
		}
	}
	qb.subscriptions[req.GetConsumerID()] = topicNames
	return &quicq.Response{
		ResponseType: quicq.ResponseType_RESPONSE_UNSUBSCRIBE,
	}, nil
}

func (qb *QuicQBroker) getOrAddTopic(name string) topic.Topic {
	var top topic.Topic
	if t, ok := qb.topicQueryMap[name]; ok {
		top = t
	} else {
		top = topic.New(name, int64(config.BrokerConfig.QueueLen()))
		qb.topicQueryMap[name] = top
		qb.topics = append(qb.topics, top)
	}
	return top
}

func (qb *QuicQBroker) sendErr(stream conn.Stream, cause string) error {
	errResp := qb.formErrResponse(cause)
	bs, err := qb.encoder.EncodeResponse(errResp)
	if err != nil {
		return err
	}
	return func() error { _, e := stream.Send(bs); return e }()
}

func (qb *QuicQBroker) formErrResponse(cause string) *quicq.Response {
	return &quicq.Response{
		ResponseType: quicq.ResponseType_RESPONSE_ERROR,
		Response: &quicq.Response_ErrResponse{
			ErrResponse: &quicq.ErrorResponse{
				Details: cause,
			},
		},
	}
}
