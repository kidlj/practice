package lists

import (
	"fmt"

	"github.com/kidlj/practice/go/containers"
)

type doublyLinkedListNode struct {
	item interface{}
	pred *doublyLinkedListNode
	succ *doublyLinkedListNode
}

// doublyLinkedListIterator is the data structure for a LinkedList external iterator.
type doublyLinkedListIterator struct {
	list    *doublyLinkedList
	current *doublyLinkedListNode
}

// doublyLinkedList is a doubly-linked implementation of a list
type doublyLinkedList struct {
	cursorIndex int
	cursorPtr   *doublyLinkedListNode
	head        *doublyLinkedListNode
	count       int
}

func NewDoubleyLinkedList() List {
	l := new(doublyLinkedList)
	l.init()
	return l
}

func (list *doublyLinkedList) init() {
	if list.head == nil {
		dummy := &doublyLinkedListNode{item: 0}
		dummy.pred, dummy.succ = dummy, dummy
		list.cursorIndex = -1
		list.cursorPtr, list.head = dummy, dummy
	}
}

func (list *doublyLinkedList) setCursor(index int) {
	if index <= list.count/2 {
		// compare the distance of target index between head and between current cursorIndex
		if index+1 < list.cursorIndex-index {
			list.cursorIndex = -1
			list.cursorPtr = list.head
		}
	} else {
		// compare the distance of target index between tail and between current cursorIndex
		if list.count-index < index-list.cursorIndex {
			list.cursorIndex = list.count
			list.cursorPtr = list.head
		}
	}
	for list.cursorIndex < index {
		list.cursorIndex++
		list.cursorPtr = list.cursorPtr.succ
	}
	for index < list.cursorIndex {
		list.cursorIndex--
		list.cursorPtr = list.cursorPtr.pred
	}
}

// Size indicates how many elements are in the list.
func (list *doublyLinkedList) Size() int {
	return list.count
}

// IsEmpty determines if the list is empty.
func (list *doublyLinkedList) IsEmpty() bool {
	return list.count == 0
}

// Clear remove all elements from the list.
func (list *doublyLinkedList) Clear() {
	list.count = 0
	list.head = nil
	list.init()
}

// Contains determines wheter e is in the list.
func (list *doublyLinkedList) Contains(e interface{}) bool {
	for ptr := list.head.succ; ptr != list.head; ptr = ptr.succ {
		if ptr.item == e {
			return true
		}
	}
	return false
}

// Insert puts element e into the list at location i.
func (list *doublyLinkedList) Insert(i int, e interface{}) error {
	if i < 0 || i > list.count {
		return fmt.Errorf("Insert: index out of bounds: %d", i)
	}
	list.setCursor(i)
	node := &doublyLinkedListNode{e, list.cursorPtr.pred, list.cursorPtr}
	list.cursorPtr.pred.succ = node
	list.cursorPtr.pred = node
	list.cursorPtr = node
	list.count++
	return nil
}

// Delete removes and returns the element at location i.
func (list *doublyLinkedList) Delete(i int) (interface{}, error) {
	if i < 0 || i >= list.count {
		return nil, fmt.Errorf("Delete: index out of bounds: %d", i)
	}
	list.setCursor(i)
	result := list.cursorPtr.item
	list.cursorPtr.pred.succ = list.cursorPtr.succ
	list.cursorPtr.succ.pred = list.cursorPtr.pred
	list.cursorPtr = list.cursorPtr.succ
	list.count--
	return result, nil
}

// Get returns the element at location i.
func (list *doublyLinkedList) Get(i int) (interface{}, error) {
	if i < 0 || i >= list.count {
		return nil, fmt.Errorf("Get: index out of bounds: %d", i)
	}
	list.setCursor(i)
	return list.cursorPtr.item, nil
}

// Put changes element i.
func (list *doublyLinkedList) Put(i int, e interface{}) error {
	if i < 0 || i >= list.count {
		return fmt.Errorf("Get: index out of bounds: %d", i)
	}
	list.setCursor(i)
	list.cursorPtr.item = e
	return nil
}

// Index returns the location of e.
// If e is not present, return -1 and false, otherwise return the location and true.
func (list *doublyLinkedList) Index(e interface{}) (int, bool) {
	for ptr, i := list.head.succ, 0; ptr != list.head; ptr, i = ptr.succ, i+1 {
		if ptr.item == e {
			return i, true
		}
	}
	return -1, false
}

// Slice makes a new list duplicating part of the list.
func (list *doublyLinkedList) Slice(i, j int) (List, error) {
	if i < 0 || i > j || j > list.count {
		return nil, fmt.Errorf("Slice: illegal indecies: %d, %d", i, j)
	}
	result := NewDoubleyLinkedList()
	for srcIndex, dstIndex := i, 0; srcIndex < j; srcIndex, dstIndex = srcIndex+1, dstIndex+1 {
		v, _ := list.Get(srcIndex)
		result.Insert(dstIndex, v)
	}
	return result, nil
}

// Equal indicates whether another List is identical to this one.
// Two Lists are identical if they are the same size and have the same
// elements in the same order.
// Precondition: element can be compared using ==.
// Precondition violation: panic.
func (list *doublyLinkedList) Equal(l List) bool {
	if list.count != l.Size() {
		return false
	}
	iter := l.NewIterator()
	v, ok := iter.Next()
	for ptr := list.head.succ; ptr != list.head; ptr = ptr.succ {
		if !ok || v != ptr.item {
			return false
		}
		v, ok = iter.Next()
	}
	return true
}

// NewIterator returns an external iterator.
func (list *doublyLinkedList) NewIterator() containers.Iterator {
	return &doublyLinkedListIterator{list, list.head.succ}
}

// Apply calls function f on every element in the list.
func (list *doublyLinkedList) Apply(f func(e interface{})) {
	for ptr := list.head.succ; ptr != list.head; ptr = ptr.succ {
		f(ptr.item)
	}
}

// Done indicates whether the iteration is complete.
func (iter *doublyLinkedListIterator) Done() bool {
	return iter.current == iter.list.head
}

// Reset prepares an iterator o traverse its associated Collection.
func (iter *doublyLinkedListIterator) Reset() {
	iter.current = iter.list.head.succ
}

// Next returns an element and an indication of whether iteration is in process.
func (iter *doublyLinkedListIterator) Next() (interface{}, bool) {
	if iter.current == iter.list.head {
		return nil, false
	}
	result := iter.current.item
	iter.current = iter.current.succ
	return result, true
}
