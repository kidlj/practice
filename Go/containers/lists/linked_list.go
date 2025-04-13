package lists

import (
	"fmt"

	"github.com/kidlj/practice/go/containers"
)

type node struct {
	item any
	succ *node
}

type linkedList struct {
	head        *node
	cursorIndex int
	cursorPtr   *node
	count       int
}

type linkedListIterator struct {
	list    *linkedList
	current *node
}

func NewLinkedList() *linkedList {
	l := new(linkedList)
	l.init()
	return l
}

func (l *linkedList) init() {
	if l.head == nil {
		dummy := &node{}
		l.head = dummy
		l.cursorPtr = dummy
		l.cursorIndex = -1
		l.count = 0
	}
}

// set the cursor to the previous position before the index
func (l *linkedList) setCursor(i int) {
	if l.cursorIndex >= i {
		l.cursorIndex = -1
		l.cursorPtr = l.head
	}

	for l.cursorIndex < i-1 {
		l.cursorIndex++
		l.cursorPtr = l.cursorPtr.succ
	}
}

func (l *linkedList) Size() int {
	return l.count
}

func (l *linkedList) IsEmpty() bool {
	return l.count == 0
}

func (l *linkedList) Clear() {
	l.count = 0
	l.head = nil
	l.init()
}

func (l *linkedList) Insert(i int, e any) error {
	if i < 0 || i > l.count {
		return fmt.Errorf("Insert: index out of bounds: %d", i)
	}
	l.setCursor(i)
	n := &node{item: e, succ: l.cursorPtr.succ}
	l.cursorPtr.succ = n
	l.count++
	return nil
}

func (l *linkedList) Get(i int) (any, error) {
	if i < 0 || i >= l.count {
		return nil, fmt.Errorf("Get: Index out of bounds: %d", i)
	}
	l.setCursor(i)
	return l.cursorPtr.succ.item, nil
}

func (l *linkedList) Put(i int, e any) error {
	if i < 0 || i >= l.count {
		return fmt.Errorf("Put: index out of bounds: %d", i)
	}

	l.setCursor(i)
	l.cursorPtr.succ.item = e
	return nil
}

func (l *linkedList) Delete(i int) (any, error) {
	if i < 0 || i >= l.count {
		return nil, fmt.Errorf("Delete: index out of bounds: %d", i)
	}

	l.setCursor(i)
	result := l.cursorPtr.succ.item
	l.cursorPtr.succ = l.cursorPtr.succ.succ
	l.count--
	return result, nil
}

func (l *linkedList) Slice(i, j int) (List, error) {
	if i < 0 || i > j || j > l.count {
		return nil, fmt.Errorf("Slice: index out of bounds: %d,%d", i, j)
	}

	result := NewLinkedList()

	for ; i < j; i++ {
		e, _ := l.Get(i)
		_ = result.Insert(result.Size(), e)
	}

	return result, nil
}

func (l *linkedList) Contains(e any) bool {
	for n := l.head.succ; n != nil; n = n.succ {
		if n.item == e {
			return true
		}
	}

	return false
}

func (l *linkedList) Index(e any) (int, bool) {
	for n, i := l.head.succ, 0; n != nil; n, i = n.succ, i+1 {
		if n.item == e {
			return i, true
		}
	}

	return -1, false
}

func (l *linkedList) Equal(list List) bool {
	if l.count != list.Size() {
		return false
	}

	iter := list.NewIterator()
	v, ok := iter.Next()
	for n := l.head.succ; n != nil; n = n.succ {
		if !ok || n.item != v {
			return false
		}
		v, ok = iter.Next()
	}

	return true
}

func (l *linkedList) Apply(f func(e any)) {
	for n := l.head.succ; n != nil; n = n.succ {
		f(n.item)
	}
}

func (l *linkedList) NewIterator() containers.Iterator {
	return &linkedListIterator{list: l, current: l.head.succ}
}

func (iter *linkedListIterator) Next() (any, bool) {
	if iter.current == nil {
		return nil, false
	}

	result := iter.current.item
	iter.current = iter.current.succ
	return result, true
}

func (iter *linkedListIterator) Reset() {
	iter.current = iter.list.head.succ
}

func (iter *linkedListIterator) Done() bool {
	return iter.current == nil
}

// Algorithms
func (l *linkedList) String() string {
	var s string
	for n := l.head.succ; n != nil; n = n.succ {
		s = s + fmt.Sprintf(" %v", n.item)
	}

	return s
}
