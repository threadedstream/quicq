package chan_queue

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/threadedstream/quicthing/internal/queue"
)

func TestChanQueue_Implementation(t *testing.T) {
	const queueCap = 256
	q := new(ChanQueue)
	q.SetCap(queueCap)
	require.Equal(t, queueCap, q.Cap())

	q.Push(20)
	q.Push(40)
	q.Push(50)
	q.Push(100)
	expectedTwenty, _ := q.Get()
	require.Equal(t, 20, expectedTwenty)
	expectedForty, _ := q.Get()
	require.Equal(t, 40, expectedForty)
	expectedFifty, _ := q.Get()
	require.Equal(t, 50, expectedFifty)
	expectedHundred, _ := q.Get()
	require.Equal(t, 100, expectedHundred)

	_, emptyQErr := q.Get()
	require.Equal(t, queue.QEmptyErr, emptyQErr)
	require.Equal(t, 0, q.Len())

	const smallCapacity = 2
	smallCappedQ := new(ChanQueue)
	smallCappedQ.SetCap(smallCapacity)

	smallCappedQ.Push(&struct{}{})
	smallCappedQ.Push(&struct{}{})
	fullQErr := q.Push(&struct{}{})
	require.Equal(t, fullQErr, queue.QFullErr)
}
