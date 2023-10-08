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

const (
	// QueueSizeMax specifies the max size of a queue
	QueueSizeMax = 256
)
