package queue

import (
	"errors"
)

type Queue interface {
	Push(val any) error
	Get() (any, error)
	Purge()
	Remove() error
	Len() int
	Cap() int
	SetCap(cap int)
}

var (
	// QFullErr is returned when there's no free space to push additional items onto queue
	QFullErr = errors.New("queue's full")
	// QEmptyErr is returned when there's no items to return from queue
	QEmptyErr = errors.New("queue's empty")
)

// NewQueue creates new instance of relevant queue
func NewQueue[T Queue](cap int) {
	// TBD
}
