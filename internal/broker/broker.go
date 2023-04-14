package broker

import (
	"bytes"
	"context"
	"errors"
	"log"
	"sync"

	"github.com/threadedstream/quicthing/internal/conn"
	"github.com/threadedstream/quicthing/internal/protocol"
	"github.com/threadedstream/quicthing/internal/server"
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
	"google.golang.org/protobuf/proto"
)

const (
	defaultBrokerAddr = "0.0.0.0:9999"
)

var (
	brokerClosedErr = errors.New("closed")
)

type Broker interface {
	Run(context.Context) error
	HandleCommand(context.Context, protocol.Request) error
}

type QuicQBroker struct {
	server        *server.QuicServer
	mu            *sync.Mutex
	subscriptions map[int64][]string
}

func New() *QuicQBroker {
	broker := &QuicQBroker{
		server:        new(server.QuicServer),
		mu:            new(sync.Mutex),
		subscriptions: make(map[int64][]string),
	}
	return broker
}

func (qb *QuicQBroker) Run(ctx context.Context) error {
	return qb.run(ctx)
}

func (qb *QuicQBroker) run(ctx context.Context) error {
	if e := qb.server.Serve(defaultBrokerAddr); e != nil {
		log.Fatalf("unable to serve: %s", e.Error())
	}

	for {
		select {
		case <-ctx.Done():
			qb.server.Shutdown()
			return brokerClosedErr
		default:
			conn, err := qb.server.AcceptClient(ctx)
			if err != nil {
				// it's highly discouraged, but we're good w/ that for the purpose of learning
				log.Println("failed to accept: " + err.Error())
				continue
			}
			log.Printf("got a new connection: %s\n", conn.RemoteAddr().String())
			go qb.handleConnection(ctx, conn)
		}
	}
}

func (qb *QuicQBroker) handleConnection(ctx context.Context, conn conn.Connection) {
	stream, err := conn.AcceptStream(ctx)
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return
	}
	streamCtx := stream.Context()
	for {
		select {
		case <-streamCtx.Done():
			stream.Close()
			return
		default:
			var p [512]byte
			_, err = stream.Rcv(p[:])
			if err != nil {
				conn.Log("Failed to read a message: %s\n", err.Error())
				continue
			}
			bs := bytes.Trim(p[:], "\x00")
			req, err := qb.decodeRequest(bs)
			if err != nil {
				stream.Send([]byte("gfy maaan"))
				continue
			}
			qb.executeRequest(ctx, req)
			// write the message back
			if _, err = stream.Send(p[:]); err != nil {
				conn.Log("Failed to write a message: %s\n", err.Error())
				continue
			}
		}
	}
}

func (qb *QuicQBroker) executeRequest(ctx context.Context, req *quicq.Request) {
	switch req.GetRequestType() {
	case quicq.RequestType_REQUEST_SUBSCRIBE:
		qb.doSubscribe(req.GetSubscribeRequest())
	case quicq.RequestType_REQUEST_UNSUBSCRIBE:
		qb.doUnsubscribe(req.GetUnsubscribeRequest())
	}
}

func (qb *QuicQBroker) doSubscribe(req *quicq.SubscribeRequest) (*quicq.Response, error) {
	qb.mu.Lock()
	defer qb.mu.Unlock()
	qb.subscriptions[req.GetConsumerID()] = append(qb.subscriptions[req.GetConsumerID()], req.GetTopic())
	return &quicq.Response{
		ResponseType: quicq.ResponseType_RESPONSE_SUBSCRIBE,
	}, nil
}

func (qb *QuicQBroker) doUnsubscribe(req *quicq.UnsubscribeRequest) (*quicq.Response, error) {
	qb.mu.Lock()
	defer qb.mu.Unlock()
	topics := qb.subscriptions[req.GetConsumerID()]
	for i := range topics {
		if topics[i] == req.GetTopic() {
			oldTopics := topics
			topics = oldTopics[:i]
			topics = append(topics, oldTopics[i+1:]...)
		}
	}

	return &quicq.Response{
		ResponseType: quicq.ResponseType_RESPONSE_UNSUBSCRIBE,
	}, nil
}

func (qb *QuicQBroker) decodeRequest(bs []byte) (*quicq.Request, error) {
	req := new(quicq.Request)
	if err := proto.Unmarshal(bs, req); err != nil {
		return nil, err
	}
	return req, nil
}
