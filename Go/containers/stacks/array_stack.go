package stacks

import "fmt"

type ArrayStack struct {
	store []any
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

func (s *ArrayStack) Push(e any) {
	s.store = append(s.store, e)
}

func (s *ArrayStack) Pop() (any, error) {
	if len(s.store) == 0 {
		return nil, fmt.Errorf("Pop: stack empty")
	}
	result := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return result, nil
}

func (s *ArrayStack) Top() (any, error) {
	if len(s.store) == 0 {
		return nil, fmt.Errorf("Top: stack empty")
	}
	result := s.store[len(s.store)-1]
	return result, nil
}
