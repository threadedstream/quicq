package common

import "sync/atomic"

var (
	// IDManager handles id generation for consumer
	IDManager = &idManager{}
)

type idManager struct {
	nextID atomic.Int64
}

func (im *idManager) GetNextID() uint64 {
	val := im.nextID.Add(1)
	return uint64(val)
}

// Message is a common message format for both producer and consumer
type Message struct {
	Key     []byte
	Payload []byte
}

const (
	// QueueSizeMax specifies the max size of a queue
	QueueSizeMax = 256
)
