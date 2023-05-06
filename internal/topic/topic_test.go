package topic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTopic_Implementation(t *testing.T) {
	topic := New("test-topic", 256)
	require.NotNil(t, topic)
}
