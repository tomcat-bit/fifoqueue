package fifoqueue

import (
	_"fmt"
	"sync"
	"errors"
	"container/list"
)

type Queue struct {
	list *list.List
	lock sync.RWMutex
	capacity int
	capLock sync.RWMutex
}

var (
	errNegativeCapacity = errors.New("Capacity must be greater than zero")
)

func NewFIFOQueue(capacity int) (*Queue, error) {
	if capacity <= 0 {
		return nil, errNegativeCapacity
	}

	return &Queue{
		list: list.New(),
		capacity: capacity,
	}, nil
}

func (q *Queue) Insert(v interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	// If cap is reached, remove the element at the beginning
	if q.list.Len() >= q.cap() {
		e := q.list.Front()
		q.list.Remove(e)
	}

	q.list.PushBack(v)
}

func (q *Queue) Front() *list.Element {
	return q.list.Front()
}

func (q *Queue) Back() *list.Element {
	return q.list.Back()
}

func (q *Queue) Length() int {
	return q.list.Len()
}

func (q *Queue) cap() int {
	q.capLock.RLock()
	defer q.capLock.RUnlock()
	return q.capacity
}

func (q *Queue) SetCapacity(cap int) error {
	if q.cap() <= 0 {
		return errNegativeCapacity
	}

	q.capLock.Lock()
	defer q.capLock.Unlock()
	q.capacity = cap
	return nil
}
