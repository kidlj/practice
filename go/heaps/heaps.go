package heaps

type Heap interface {
	Size() int
	DeleteMin() (int, error)
	FindMin() (int, error)
	Insert(v int)
	Print() string
}

var _ Heap = &BinaryHeap{}
var _ Heap = &GenericBinaryHeap[int]{}
