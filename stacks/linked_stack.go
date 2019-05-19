package stacks

import "errors"

type LinkedStack struct {
	topPtr *node
	count  int
}

type node struct {
	item interface{}
	next *node
}

func (s *LinkedStack) IsEmpty() bool {
	return s.count == 0
}

func (s *LinkedStack) Size() int {
	return s.count
}

func (s *LinkedStack) Clear() {
	s.topPtr = nil
	s.count = 0
}

func (s *LinkedStack) Push(item interface{}) {
	s.topPtr = &node{item, s.topPtr}
	s.count++
}

func (s *LinkedStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Pop: stack is empty")
	}
	result := s.topPtr.item
	s.topPtr = s.topPtr.next
	s.count--
	return result, nil
}

func (s *LinkedStack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Top: stack is empty")
	}
	return s.topPtr.item, nil
}
