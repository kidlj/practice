package stacks

import "errors"

type ArrayStack struct {
	store []interface{}
}

func (s *ArrayStack) Push(item interface{}) {
	s.store = append(s.store, item)
}

func (s *ArrayStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Pop: stack is empty")
	}

	result := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return result, nil
}

func (s *ArrayStack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Top: stack is empty")
	}

	return s.store[len(s.store)-1], nil
}

func (s *ArrayStack) IsEmpty() bool {
	return len(s.store) == 0
}

func (s *ArrayStack) Size() int {
	return len(s.store)
}

func (s *ArrayStack) Clear() {
	s.store = nil
}
