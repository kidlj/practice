package queue

import "errors"

type ArrayQueue struct {
	count      int
	frontIndex int
	store      []interface{}
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
