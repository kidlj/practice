package lists

import (
	"fmt"

	"github.com/kidlj/demo/containers"
)

// List is the interface for lists in the container hierarchy.
type List interface {
	containers.Collection
	Insert(i int, e interface{}) error
	Delete(i int) (interface{}, error)
	Get(i int) (interface{}, error)
	Put(i int, e interface{}) error
	Index(e interface{}) (int, bool)
	Slice(i, j int) (List, error)
	Equal(l List) bool
}

// ArrayList is a contiguous implemention of a list.
type ArrayList struct {
	store []interface{}
	count int
}

// Size returns the list size.
func (list *ArrayList) Size() int {
	return len(list.store)
}

// IsEmpty returns if the list is empty.
func (list *ArrayList) IsEmpty() bool {
	return len(list.store) == 0
}

// Clear clears the list.
func (list *ArrayList) Clear() {
	list.store = nil
}

// Insert adds an element in the i index.
func (list *ArrayList) Insert(i int, e interface{}) error {
	if i < 0 || i > list.count {
		return fmt.Errorf("Insert: index out of bounds: %d", i)
	}
	list.store = append(list.store, 0)
	copy(list.store[i+1:], list.store[i:list.count])
	list.store[i] = e
	list.count++
	return nil
}

// Delete removes an element in the i index.
func (list *ArrayList) Delete(i int) (interface{}, error) {
	if i < 0 || i >= list.count {
		return nil, fmt.Errorf("Delete: index out of bounds: %d", i)
	}
	result := list.store[i]
	copy(list.store[i:], list.store[i+1:])
	list.count--
	list.store = list.store[:list.count]
	return result, nil
}
