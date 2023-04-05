package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImplementation_Queue(t *testing.T) {
	q := New(10)
	Put(&q, 20)
	require.Equal(t, 20, q.Get())

	Put(&q, 30)
	Put(&q, 40)
	require.Equal(t, 4, q.Len())

	Remove(&q)
	require.Equal(t, 3, q.Len())
	require.Equal(t, 30, q.Get())

	Purge(&q)
	require.Equal(t, 0, q.Len())

	Put(&q, 69)
	require.Equal(t, 69, q.Get())

}

func BenchmarkQueueInsertBench(b *testing.B) {
	q := New(0)
	for i := 1; i < b.N; i++ {
		Put(&q, i)
	}

	for i := 0; i < q.Len(); i++ {
		Remove(&q)
	}
}
