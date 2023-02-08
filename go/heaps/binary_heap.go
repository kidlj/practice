package heaps

import (
	"fmt"
	"math"
)

type BinaryHeap struct {
	elements []int
	size     int
}

func (h *BinaryHeap) init() {
	if h.elements == nil {
		h.elements = []int{math.MinInt} // dummy node
		h.size = 0
	}
}

func (h *BinaryHeap) Size() int {
	return h.size
}

func (h *BinaryHeap) IsEmpty() bool {
	return h.size == 0
}

func (h *BinaryHeap) FindMin() (int, error) {
	if h.IsEmpty() {
		return -1, fmt.Errorf("empty heap")
	}
	return h.elements[1], nil
}

func (h *BinaryHeap) DeleteMin() (int, error) {
	if h.IsEmpty() {
		return -1, fmt.Errorf("empty heap")
	}
	min := h.elements[1]
	h.elements[1] = h.elements[h.size]
	h.size--

	h.percolateDown(1)

	return min, nil
}

func (h *BinaryHeap) percolateDown(i int) {
	v := h.elements[i]
	for child := i * 2; i*2 <= h.size; i, child = child, i*2 {
		if child != h.size && h.elements[child+1] < h.elements[child] {
			child++
		}
		if v > h.elements[child] {
			h.elements[i] = h.elements[child]
		} else {
			break
		}
	}

	h.elements[i] = v
}

func (h *BinaryHeap) Insert(v int) {
	h.init()

	h.elements = append(h.elements, v)
	h.size++

	h.percolateUp(h.size)
}

func (h *BinaryHeap) percolateUp(i int) {
	v := h.elements[i]
	for ; h.elements[i/2] > v; i /= 2 {
		h.elements[i] = h.elements[i/2]
	}
	h.elements[i] = v
}

func (h *BinaryHeap) IncreaseKey(key int, delta int) {
	h.elements[key] += delta
	h.percolateDown(key)
}

func (h *BinaryHeap) DecreaseKey(key int, delta int) {
	h.elements[key] -= delta
	h.percolateUp(key)
}

func (h *BinaryHeap) Delete(key int) (int, error) {
	if key > h.size || key <= 0 {
		return -1, fmt.Errorf("invalid key")
	}

	v := h.elements[key]
	h.elements[key] = math.MinInt
	h.percolateUp(key)
	_, err := h.DeleteMin()
	if err != nil {
		return -1, err
	}
	return v, nil
}

// O(N)
func BuildBinaryHeap(v []int) *BinaryHeap {
	h := &BinaryHeap{}
	h.init()
	h.elements = append(h.elements, v...)
	h.size = len(v)

	for i := h.size / 2; i >= 1; i-- {
		h.percolateDown(i)
	}
	return h
}

func (h *BinaryHeap) Print() string {
	return fmt.Sprintf("%v", h.elements[1:h.Size()+1])
}
