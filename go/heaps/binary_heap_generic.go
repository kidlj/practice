package heaps

import "fmt"

type GenericBinaryHeap[T any] struct {
	cmp      func(a, b T) int
	elements []T
}

func NewGenericBinaryHeap[T any](cmp func(a, b T) int) *GenericBinaryHeap[T] {
	return &GenericBinaryHeap[T]{
		cmp:      cmp,
		elements: make([]T, 0),
	}
}

func FromSlice[T any](cmp func(a, b T) int, src []T) *GenericBinaryHeap[T] {
	h := &GenericBinaryHeap[T]{
		cmp:      cmp,
		elements: src,
	}
	for i := len(src)/2 - 1; i >= 0; i-- {
		h.percolateDown(i)
	}
	return h
}

func From[T any](cmp func(a, b T) int, t ...T) *GenericBinaryHeap[T] {
	return FromSlice(cmp, t)
}

func (h *GenericBinaryHeap[T]) Size() int {
	return len(h.elements)
}

func (h *GenericBinaryHeap[T]) IsEmpty() bool {
	return len(h.elements) == 0
}

func (h *GenericBinaryHeap[T]) DeleteMin() (T, error) {
	var result T
	if h.IsEmpty() {
		return result, fmt.Errorf("heap empty")
	}
	result = h.elements[0]
	h.elements[0] = h.elements[h.Size()-1]
	h.elements = h.elements[:h.Size()-1]
	h.percolateDown(0)
	return result, nil
}

func (h *GenericBinaryHeap[T]) FindMin() (T, error) {
	var result T
	if h.IsEmpty() {
		return result, fmt.Errorf("heap empty")
	}
	return h.elements[0], nil
}

func (h *GenericBinaryHeap[T]) Insert(v T) {
	h.elements = append(h.elements, v)
	h.percolateUp(h.Size() - 1)
}

func (h *GenericBinaryHeap[T]) percolateDown(i int) {
	smallest := i
	l := leftChildIndex(i)
	r := rightChildIndex(i)

	if l < h.Size() && h.cmp(h.elements[i], h.elements[l]) > 0 {
		smallest = l
	}
	if r < h.Size() && h.cmp(h.elements[smallest], h.elements[r]) > 0 {
		smallest = r
	}
	if smallest != i {
		h.swap(smallest, i)
		h.percolateDown(smallest)
	}
}

// NOTE: -1 / 2 = 0
func (h *GenericBinaryHeap[T]) percolateUp(i int) {
	for h.cmp(h.elements[i], h.elements[parentIndex(i)]) < 0 {
		h.swap(i, parentIndex(i))
		i = parentIndex(i)
	}
}

func leftChildIndex(i int) int {
	return i*2 + 1
}

func rightChildIndex(i int) int {
	return leftChildIndex(i) + 1
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *GenericBinaryHeap[T]) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *GenericBinaryHeap[T]) Print() string {
	return fmt.Sprintf("%v", h.elements)
}
