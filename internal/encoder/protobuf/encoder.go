package protobuf

import (
	"errors"

	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
	"google.golang.org/protobuf/proto"
)

const (
	terminatingByte = '\xee'
)

type Encoder struct{}

type Decoder struct{}

func NewProtoEncoder() *Encoder { return &Encoder{} }

func (enc *Encoder) EncodeRequest(req *quicq.Request) ([]byte, error) {
	return enc.encode(req)
}

func (enc *Encoder) EncodeResponse(resp *quicq.Response) ([]byte, error) {
	return enc.encode(resp)
}

func (enc *Encoder) encode(obj proto.Message) ([]byte, error) {
	bs, err := proto.Marshal(obj)
	if err != nil {
		return nil, err
	}
	// note: the idea borrowed from rabbitmq protocol implementation
	bs = append(bs, terminatingByte)
	return bs, nil
}

func NewProtoDecoder() *Decoder { return &Decoder{} }

func (pd *Decoder) DecodeRequest(bs []byte) (*quicq.Request, error) {
	req := new(quicq.Request)
	if err := pd.decode(bs, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (pd *Decoder) DecodeResponse(bs []byte) (*quicq.Response, error) {
	resp := new(quicq.Response)
	if err := pd.decode(bs, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (pd *Decoder) decode(bs []byte, message proto.Message) error {
	if bs[len(bs)-1] == '\xee' {
		bs = bs[:len(bs)-1]
		// remove trailing \xee
		if err := proto.Unmarshal(bs, message); err != nil {
			return err
		}
		return nil
	}
	return errors.New("no trailing \"\\xee\" character")
}
