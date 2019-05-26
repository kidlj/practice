package lists

import (
	"fmt"

	"github.com/kidlj/demo/containers"
)

// ArrayList is a contiguous implementation of a list.
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

// Get gets the element at index i.
func (list *ArrayList) Get(i int) (interface{}, error) {
	if i < 0 || i >= list.count {
		return nil, fmt.Errorf("Get: index out of bounds: %d", i)
	}
	return list.store[i], nil
}

// Put changes element i.
func (list *ArrayList) Put(i int, e interface{}) error {
	if i < 0 || i >= list.count {
		return fmt.Errorf("Get: index out of bounds: %d", i)
	}
	list.store[i] = e
	return nil
}

// Index returns the index of the provided element in the list.
func (list *ArrayList) Index(e interface{}) (int, bool) {
	for i := 0; i < list.count; i++ {
		if list.store[i] == e {
			return i, true
		}
	}
	return 0, false
}

// Contains returns true if element e is in the list.
func (list *ArrayList) Contains(e interface{}) bool {
	for i := 0; i < list.count; i++ {
		if list.store[i] == e {
			return true
		}
	}
	return false
}

// Slice makes a new list duplicating part of this list.
func (list *ArrayList) Slice(i, j int) (List, error) {
	if i < 0 || i > j || j > list.count {
		return nil, fmt.Errorf("Slice: illegal indecies: i = %d, j = %d", i, j)
	}
	result := new(ArrayList)
	result.count = j - i
	result.store = make([]interface{}, result.count)
	copy(result.store[0:], list.store[i:j])
	return result, nil
}

// Equal determines whether another List is identical to this one.
// Two Lists are identical if they are the same size and have the same elements
// in the same order.
func (list *ArrayList) Equal(l List) bool {
	if list.count != l.Size() {
		return false
	}

	iter := l.NewIterator()
	v, ok := iter.Next()
	for i := 0; i < list.count; i++ {
		if !ok || list.store[i] != v {
			return false
		}
		v, ok = iter.Next()
	}

	return true
}

type arrayListIterator struct {
	list *ArrayList
	next int
}

func (iter *arrayListIterator) Reset() {
	iter.next = 0
}

func (iter *arrayListIterator) Done() bool {
	return iter.next >= iter.list.Size()
}

func (iter *arrayListIterator) Next() (interface{}, bool) {
	if iter.next < iter.list.Size() {
		result, _ := iter.list.Get(iter.next)
		iter.next++
		return result, true
	}
	return nil, false
}

// NewIterator returns an Iterator.
func (list *ArrayList) NewIterator() containers.Iterator {
	result := new(arrayListIterator)
	result.list = list
	return result
}

// Apply calls function f on every element in the list.
func (list *ArrayList) Apply(f func(interface{})) {
	for i := 0; i < list.count; i++ {
		f(list.store[i])
	}
}
