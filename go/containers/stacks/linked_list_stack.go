package stacks

import (
	"fmt"

	"github.com/kidlj/playground/go/containers/lists"
)

type LinkedListStack struct {
	store lists.List
}

func NewLinkedListStack() *LinkedListStack {
	s := new(LinkedListStack)
	s.store = new(lists.LinkedList)
	return s
}

func (s *LinkedListStack) Size() int {
	return s.store.Size()
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.store.IsEmpty()
}

func (s *LinkedListStack) Clear() {
	s.store.Clear()
}

func (s *LinkedListStack) Push(e any) {
	_ = s.store.Insert(0, e)
}

func (s *LinkedListStack) Pop() (any, error) {
	if s.store.Size() == 0 {
		return nil, fmt.Errorf("Pop: stack empty")
	}

	result, _ := s.store.Delete(0)
	return result, nil
}

func (s *LinkedListStack) Top() (any, error) {
	if s.store.Size() == 0 {
		return nil, fmt.Errorf("Top: stack empty")
	}

	return s.store.Get(0)
}
