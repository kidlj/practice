package queues

import "errors"

type Queue interface {
	IsEmpty() bool
	Size() int
	Clear()
	Enter(any)
	Leave() (any, error)
	Front() (any, error)
}

// ArrayQueue is the array implementation
type ArrayQueue struct {
	count      int
	frontIndex int
	store      []interface{}
}

func NewArrayQueue() *ArrayQueue {
	return new(ArrayQueue)
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.count == 0
}

func (q *ArrayQueue) Size() int {
	return q.count
}

func (q *ArrayQueue) Clear() {
	q.count = 0
}

func (q *ArrayQueue) Enter(e interface{}) {
	if q.store == nil {
		q.store = make([]interface{}, 4)
	}

	if q.count == len(q.store) {
		q.store = append(q.store, q.store...)
	}

	q.store[(q.frontIndex+q.count)%len(q.store)] = e
	q.count++
}

func (q *ArrayQueue) Leave() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	result := q.store[q.frontIndex]
	q.frontIndex = (q.frontIndex + 1) % len(q.store)
	q.count--
	return result, nil
}

func (q *ArrayQueue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	return q.store[q.frontIndex], nil
}

// LinkedQueue is the linked implementation
type LinkedQueue struct {
	frontPtr *node
	rearPtr  *node
	count    int
}

type node struct {
	item interface{}
	next *node
}

func NewLinkedQueue() *LinkedQueue {
	return new(LinkedQueue)
}

func (q *LinkedQueue) IsEmpty() bool {
	return q.count == 0
}

func (q *LinkedQueue) Size() int {
	return q.count
}

func (q *LinkedQueue) Clear() {
	q.frontPtr = nil
	q.rearPtr = nil
	q.count = 0
}

func (q *LinkedQueue) Enter(e interface{}) {
	n := &node{e, nil}
	if q.IsEmpty() {
		q.frontPtr = n
	} else {
		q.rearPtr.next = n
	}
	q.rearPtr = n
	q.count++
}

func (q *LinkedQueue) Leave() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	result := q.frontPtr.item
	q.frontPtr = q.frontPtr.next
	if q.frontPtr == nil {
		q.rearPtr = nil
	}
	q.count--
	return result, nil
}

func (q *LinkedQueue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	return q.frontPtr.item, nil
}
