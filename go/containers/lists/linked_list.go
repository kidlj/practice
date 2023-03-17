package lists

import (
	"fmt"

	"github.com/kidlj/practice/go/containers"
)

type node struct {
	item any
	succ *node
}

type LinkedList struct {
	head        *node
	cursorIndex int
	cursorPtr   *node
	count       int
}

type LinkedListIterator struct {
	list    *LinkedList
	current *node
}

func NewLinkedList() *LinkedList {
	l := new(LinkedList)
	l.init()
	return l
}

func (l *LinkedList) init() {
	if l.head == nil {
		dummy := &node{}
		dummy.succ = dummy
		l.head = dummy
		l.cursorPtr = dummy
		l.cursorIndex = -1
		l.count = 0
	}
}

func (l *LinkedList) setCursor(i int) {
	if l.cursorIndex >= i {
		l.cursorIndex = -1
		l.cursorPtr = l.head
	}

	for l.cursorIndex < i-1 {
		l.cursorIndex++
		l.cursorPtr = l.cursorPtr.succ
	}
}

func (l *LinkedList) Size() int {
	return l.count
}

func (l *LinkedList) IsEmpty() bool {
	return l.count == 0
}

func (l *LinkedList) Clear() {
	l.count = 0
	l.head = nil
}

func (l *LinkedList) Insert(i int, e any) error {
	if i < 0 || i > l.count {
		return fmt.Errorf("Insert: index out of bounds: %d", i)
	}
	l.setCursor(i)
	n := &node{item: e, succ: l.cursorPtr.succ}
	l.cursorPtr.succ = n
	l.count++
	return nil
}

func (l *LinkedList) Get(i int) (any, error) {
	if i < 0 || i >= l.count {
		return nil, fmt.Errorf("Get: Index out of bounds: %d", i)
	}
	l.setCursor(i)
	return l.cursorPtr.succ.item, nil
}

func (l *LinkedList) Append(e any) {
	l.Insert(l.count, e)
}

func (l *LinkedList) Put(i int, e any) error {
	if i < 0 || i >= l.count {
		return fmt.Errorf("Put: index out of bounds: %d", i)
	}

	l.setCursor(i)
	l.cursorPtr.succ.item = e
	return nil
}

func (l *LinkedList) Delete(i int) (any, error) {
	if i < 0 || i >= l.count {
		return nil, fmt.Errorf("Delete: index out of bounds: %d", i)
	}

	l.setCursor(i)
	result := l.cursorPtr.succ.item
	l.cursorPtr.succ = l.cursorPtr.succ.succ
	l.count--
	return result, nil
}

func (l *LinkedList) Slice(i, j int) (List, error) {
	if i < 0 || i > j || j > l.count {
		return nil, fmt.Errorf("Slice: index out of bounds: %d,%d", i, j)
	}

	result := NewLinkedList()

	for ; i < j; i++ {
		e, _ := l.Get(i)
		result.Append(e)
	}

	return result, nil
}

func (l *LinkedList) Contains(e any) bool {
	for n := l.head.succ; n != l.head; n = n.succ {
		if n.item == e {
			return true
		}
	}

	return false
}

func (l *LinkedList) Index(e any) (int, bool) {
	for n, i := l.head.succ, 0; n != l.head; n, i = n.succ, i+1 {
		if n.item == e {
			return i, true
		}
	}

	return -1, false
}

func (l *LinkedList) Equal(list List) bool {
	if l.count != list.Size() {
		return false
	}

	iter := list.NewIterator()
	v, ok := iter.Next()
	for n := l.head.succ; n != l.head; n = n.succ {
		if !ok || n.item != v {
			return false
		}
		v, ok = iter.Next()
	}

	return true
}

func (l *LinkedList) Apply(f func(e any)) {
	for n := l.head.succ; n != l.head; n = n.succ {
		f(n.item)
	}
}

func (l *LinkedList) NewIterator() containers.Iterator {
	return &LinkedListIterator{list: l, current: l.head.succ}
}

func (iter *LinkedListIterator) Next() (any, bool) {
	if iter.current == iter.list.head {
		return nil, false
	}

	result := iter.current.item
	iter.current = iter.current.succ
	return result, true
}

func (iter *LinkedListIterator) Reset() {
	iter.current = iter.list.head.succ
}

func (iter *LinkedListIterator) Done() bool {
	return iter.current == iter.list.head
}

func (l *LinkedList) printLots(seqList List) {
	iter := seqList.NewIterator()
	v, ok := iter.Next()
	n := l.head.succ
	seq := 1
	for n != l.head && ok {
		if seq == v {
			fmt.Printf("%v ", n.item)
			v, ok = iter.Next()
		}
		seq++
		n = n.succ
	}
}
