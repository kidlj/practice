package lists

import (
	"fmt"

	"github.com/kidlj/practice/go/containers"
)

// arrayList is a contiguous implementation of a list.
type arrayList struct {
	store []interface{}
	count int
}

func NewArrayList() List {
	return new(arrayList)
}

// Size returns the list size.
func (list *arrayList) Size() int {
	return len(list.store)
}

// IsEmpty returns if the list is empty.
func (list *arrayList) IsEmpty() bool {
	return len(list.store) == 0
}

// Clear clears the list.
func (list *arrayList) Clear() {
	list.store = nil
	list.count = 0
}

// Insert adds an element in the i index.
func (list *arrayList) Insert(i int, e interface{}) error {
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
func (list *arrayList) Delete(i int) (interface{}, error) {
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
func (list *arrayList) Get(i int) (interface{}, error) {
	if i < 0 || i >= list.count {
		return nil, fmt.Errorf("Get: index out of bounds: %d", i)
	}
	return list.store[i], nil
}

// Put changes element i.
func (list *arrayList) Put(i int, e interface{}) error {
	if i < 0 || i >= list.count {
		return fmt.Errorf("Put: index out of bounds: %d", i)
	}
	list.store[i] = e
	return nil
}

// Index returns the index of the provided element in the list.
func (list *arrayList) Index(e interface{}) (int, bool) {
	for i := 0; i < list.count; i++ {
		if list.store[i] == e {
			return i, true
		}
	}
	return 0, false
}

func (list *arrayList) Append(e any) error {
	return list.Insert(list.Size(), e)
}

// Contains returns true if element e is in the list.
func (list *arrayList) Contains(e interface{}) bool {
	for i := 0; i < list.count; i++ {
		if list.store[i] == e {
			return true
		}
	}
	return false
}

// Slice makes a new list duplicating part of this list.
func (list *arrayList) Slice(i, j int) (List, error) {
	if i < 0 || i > j || j > list.count {
		return nil, fmt.Errorf("Slice: illegal indecies: i = %d, j = %d", i, j)
	}
	result := new(arrayList)
	result.count = j - i
	result.store = make([]interface{}, result.count)
	copy(result.store[0:], list.store[i:j])
	return result, nil
}

// Equal determines whether another List is identical to this one.
// Two Lists are identical if they are the same size and have the same elements
// in the same order.
func (list *arrayList) Equal(l List) bool {
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
	list *arrayList
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
func (list *arrayList) NewIterator() containers.Iterator {
	result := new(arrayListIterator)
	result.list = list
	return result
}

// Apply calls function f on every element in the list.
func (list *arrayList) Apply(f func(interface{})) {
	for i := 0; i < list.count; i++ {
		f(list.store[i])
	}
}
