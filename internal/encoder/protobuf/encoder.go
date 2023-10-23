package protobuf

import (
	"bytes"
	"encoding/gob"
	"errors"
	"log"

	"github.com/threadedstream/quicthing/internal/encoder"
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
	"google.golang.org/protobuf/proto"
)

const (
	terminatingByte = '\xee'
)

var (
	ErrNoTrailingTerminatingCharacter   = errors.New("no trailing '\\xee' character")
	ErrPacketDataNotEqualPacketDataSize = errors.New("len(packet.Data) != len(packet.DataSize)")
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
	packet := new(encoder.Packet)
	buf := bytes.Buffer{}

	data, err := proto.Marshal(obj)
	if err != nil {
		return nil, err
	}

	gobEnc := gob.NewEncoder(&buf)

	data = append(data, terminatingByte)
	packet.Data = data
	packet.DataSize = int32(len(data))

	if err = gobEnc.Encode(packet); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
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
	packet := new(encoder.Packet)
	buf := bytes.NewBuffer(bs)

	gobDec := gob.NewDecoder(buf)

	if err := gobDec.Decode(packet); err != nil {
		log.Println("LEN(BS) = ", len(bs))
		return err
	}

	if int32(len(packet.Data)) != packet.DataSize {
		return ErrPacketDataNotEqualPacketDataSize
	}

	// remove trailing \xee
	if packet.Data[packet.DataSize-1] == terminatingByte {
		packet.Data = packet.Data[:packet.DataSize-1]
		if err := proto.Unmarshal(packet.Data, message); err != nil {
			return err
		}
		return nil
	}

	return ErrNoTrailingTerminatingCharacter
}
