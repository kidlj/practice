package heaps

type Heaps interface {
	DeleteMin() (int, error)
	FindMin() (int, error)
	Insert(v int)
}

var _ Heaps = &BinaryHeap{}
