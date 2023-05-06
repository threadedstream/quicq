package chan_queue

import (
	"github.com/threadedstream/quicthing/internal/queue"
)

// ChanQueue is an implementation of Queue interface
type ChanQueue struct {
	internalQ chan any
}

// SetCap initializes internalQ
func (cq *ChanQueue) SetCap(cap int) {
	cq.internalQ = make(chan any, cap)
}

// Push pushes val onto a queue
func (cq *ChanQueue) Push(val any) error {
	if cq.Len() == cq.Cap() {
		return queue.QFullErr
	}
	cq.internalQ <- val
	return nil
}

// Get pulls item from a queue
func (cq *ChanQueue) Get() (any, error) {
	if cq.Len() == 0 {
		return nil, queue.QEmptyErr
	}
	return <-cq.internalQ, nil
}

// Purge empties a whole queue
func (cq *ChanQueue) Purge() {
	for _, e := cq.Get(); e == nil; {
	}
}

// Remove removes item from a queue
func (cq *ChanQueue) Remove() error {
	_, e := cq.Get()
	return e
}

// Len returns len of internal channel
func (cq *ChanQueue) Len() int {
	return len(cq.internalQ)
}

// Cap returns capacity of internal channel
func (cq *ChanQueue) Cap() int {
	return cap(cq.internalQ)
}
