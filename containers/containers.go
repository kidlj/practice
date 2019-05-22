package containers

// Container is the root type in the containers hierarchy.
type Container interface {
	Size() int
	IsEmpty() bool
	Clear()
}

// Collection is the ancestor type of all traversible containers in the hierarchy.
type Collection interface {
	Container
	Containes(e interface{}) bool
	NewIterator() Iterator
	Apply(f func(interface{}))
}

// Iterator is the interface for all Collection external iterators.
type Iterator interface {
	Reset()
	Done() bool
	Next() (interface{}, bool)
}
