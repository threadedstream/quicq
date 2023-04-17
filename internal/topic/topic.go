package topic

import (
	"errors"
	"github.com/threadedstream/quicthing/pkg/proto/quicq/v1"
)

const (
	queueSizeMax = 256
)

var (
	duplicateConsumerErr = errors.New("duplicate consumer")
	queueEmptyErr        = errors.New("queue's empty")
	queueFullErr         = errors.New("queue's full")
)

type Topic interface {
	Name() string
	Push(record *quicq.Record) error
	GetRecordBatch() ([]*quicq.Record, error)
	TieConsumer(consumerID int64) error
	GetConsumers() []int64
	EvictConsumer(consumerID int64) error
}

// QuicQTopic is an implementation of a Topic interface
type QuicQTopic struct {
	name      string
	q         chan *quicq.Record
	head      int
	consumers []int64
}

// New returns fresh instance of QuicQTopic object
func New(name string, cap int64) *QuicQTopic {
	return &QuicQTopic{
		name: name,
		q:    make(chan *quicq.Record, cap),
	}
}

// Name returns name of a topic
func (qt *QuicQTopic) Name() string {
	return qt.name
}

// TieConsumer associates a specific consumer to a topic
func (qt *QuicQTopic) TieConsumer(consumerID int64) error {
	// find out if consumer's already subscribed
	// do it in O(n) time for now
	hasDup := false
	for _, c := range qt.consumers {
		if c == consumerID {
			hasDup = true
			break
		}
	}
	if hasDup {
		return duplicateConsumerErr
	}
	qt.consumers = append(qt.consumers, consumerID)
	return nil
}

// GetConsumers returns a list of subscribed consumers
func (qt *QuicQTopic) GetConsumers() []int64 {
	return qt.consumers
}

// Push pushes record onto a queue
func (qt *QuicQTopic) Push(record *quicq.Record) error {
	if len(qt.q) == queueSizeMax {
		return queueFullErr
	}
	qt.q <- record
	return nil
}

// EvictConsumer evicts consumer from a list of subscribed consumers
func (qt *QuicQTopic) EvictConsumer(consumerID int64) error {
	for i := range qt.consumers {
		if consumerID == qt.consumers[i] {
			qt.consumers[i] = qt.consumers[len(qt.consumers)-1]
			qt.consumers = qt.consumers[:len(qt.consumers)-1]
			break
		}
	}
	return nil
}

func (qt *QuicQTopic) GetRecordBatch() ([]*quicq.Record, error) {
	var recs []*quicq.Record
	if len(qt.q) == 0 {
		return nil, queueEmptyErr
	}
	for r := range qt.q {
		recs = append(recs, r)
		if len(qt.q) == 0 {
			break
		}
	}
	return recs, nil
}
