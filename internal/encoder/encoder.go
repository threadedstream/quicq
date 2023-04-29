package encoder

import (
	"bytes"
	"github.com/golang/protobuf/proto"
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

type ProtoEncoder struct{}

type ProtoDecoder struct{}

func NewProtoEncoder() *ProtoEncoder { return &ProtoEncoder{} }

func (pe *ProtoEncoder) EncodeRequest(req *quicq.Request) ([]byte, error) {
	return pe.encode(req)
}

func (pe *ProtoEncoder) EncodeResponse(resp *quicq.Response) ([]byte, error) {
	return pe.encode(resp)
}

func (pe *ProtoEncoder) encode(obj proto.Message) ([]byte, error) {
	bs, err := proto.Marshal(obj)
	if err != nil {
		return nil, err
	}
	// note: the idea borrowed from rabbitmq protocol implementation
	bs = append(bs, '\xee')
	return bs, nil
}

func NewProtoDecoder() *ProtoDecoder { return &ProtoDecoder{} }

func (pd *ProtoDecoder) DecodeRequest(bs []byte) (*quicq.Request, error) {
	req := new(quicq.Request)
	if err := pd.decode(bs, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (pd *ProtoDecoder) DecodeResponse(bs []byte) (*quicq.Response, error) {
	resp := new(quicq.Response)
	if err := pd.decode(bs, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (pd *ProtoDecoder) decode(bs []byte, message proto.Message) error {
	bs = bytes.Trim(bs, "\x00")
	// remove trailing \xee
	bs = bs[:len(bs)-1]
	if err := proto.Unmarshal(bs, message); err != nil {
		return err
	}
	return nil
}
