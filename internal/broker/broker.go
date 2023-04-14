package broker

import (
	"bytes"
	"context"
	"errors"
	"log"

	"github.com/threadedstream/quicthing/internal/conn"
	"github.com/threadedstream/quicthing/internal/protocol"
	"github.com/threadedstream/quicthing/internal/server"
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
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
	server *server.QuicServer
}

func New() *QuicQBroker {
	broker := &QuicQBroker{
		server: new(server.QuicServer),
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
			wellFormed := bytes.Trim(p[:], "\x00")
			qb.decodeRequest()
			// write the message back
			if _, err = stream.Send(p[:]); err != nil {
				conn.Log("Failed to write a message: %s\n", err.Error())
				continue
			}
		}
	}
}

func (qb *QuicQBroker) decodeRequest(bs []byte) *quicq.Request {

}
