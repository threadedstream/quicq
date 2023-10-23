package encoder

import (
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Packet struct {
	proto.Message
	Data     []byte
	DataSize int32
}

func (Packet) ProtoReflect() protoreflect.Message {
	return nil
}

type Encoder interface {
	EncodeRequest(*quicq.Request) ([]byte, error)
	EncodeResponse(*quicq.Response) ([]byte, error)
}

type Decoder interface {
	DecodeRequest([]byte) (*quicq.Request, error)
	DecodeResponse([]byte) (*quicq.Response, error)
}
