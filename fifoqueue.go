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

func (q *Queue) Front() interface{} {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.list.Front().Value
}

func (q *Queue) Back() interface{} {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.list.Back().Value
}

func (q *Queue) Length() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.list.Len()
}

func (q *Queue) cap() int {
	q.capLock.RLock()
	defer q.capLock.RUnlock()
	return q.capacity
}

func (q *Queue) Elements() []interface{} {
	elements := make([]interface{}, 0)
	q.lock.RLock()
	defer q.lock.RUnlock()
	for e := q.list.Front(); e != nil; e = e.Next() {
		elements = append(elements, e)
	}
	return elements
}

func (q *Queue) Exists(v interface{}) bool {
	q.lock.RLock()
	defer q.lock.RUnlock()
	for e := q.list.Front(); e != nil; e = e.Next() {
		if v == e.Value {
			return true
		}
	}
	return false
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

