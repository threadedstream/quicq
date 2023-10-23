package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/threadedstream/quicthing/internal/encoder/protobuf"
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
)

func BenchmarkEncodeDecode(b *testing.B) {
	enc := protobuf.NewProtoEncoder()
	dec := protobuf.NewProtoDecoder()

	var records []*quicq.Record
	for i := 1; i <= b.N; i++ {
		records = append(records, &quicq.Record{
			Key:     []byte(fmt.Sprintf("key%d", i)),
			Payload: []byte(fmt.Sprintf("value%d", i)),
		})
	}

	b.ResetTimer()

	resp := &quicq.Response{
		ResponseType: quicq.ResponseType_RESPONSE_POLL,
		Response: &quicq.Response_PollResponse{
			PollResponse: &quicq.PollResponse{
				Records: records,
			},
		},
	}

	for i := 0; i < b.N; i++ {
		bs, err := enc.EncodeResponse(resp)
		require.NoError(b, err)
		r, err := dec.DecodeResponse(bs)
		require.NoError(b, err)
		require.Equal(b, r.GetResponseType(), quicq.ResponseType_RESPONSE_POLL)
	}
}
