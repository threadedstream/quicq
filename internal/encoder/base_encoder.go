package encoder

import (
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
)

type Encoder interface {
	EncodeRequest(*quicq.Request) ([]byte, error)
	EncodeResponse(*quicq.Response) ([]byte, error)
}

type Decoder interface {
	DecodeRequest([]byte) (*quicq.Request, error)
	DecodeResponse([]byte) (*quicq.Response, error)
}
