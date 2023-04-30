package common

import "sync"

var (
	// IDManager handles id generation for consumer
	IDManager = &idManager{
		mu: new(sync.Mutex),
	}
)

type idManager struct {
	mu     *sync.Mutex
	nextID uint64
}

func (im *idManager) GetNextID() uint64 {
	im.mu.Lock()
	defer im.mu.Unlock()
	im.nextID++
	return im.nextID
}

const (
	// QueueSizeMax specifies the max size of a queue
	QueueSizeMax = 256
)
