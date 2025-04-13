package stacks

import "testing"

func TestStacks(t *testing.T) {
	testStack(t, new(ArrayStack), "ArrayStack ")
	testStack(t, NewLinkedListStack(), "LinkedListStack ")
}

func testStack(t *testing.T, s Stack, name string) {
	if !s.IsEmpty() || s.Size() != 0 {
		t.Error(name + "should be empty and size should be 0 when new")
	}

	if v, err := s.Top(); err == nil {
		t.Errorf(name+"Top should fail on emtpy stack, instead returns %v", v)
	}

	if v, err := s.Pop(); err == nil {
		t.Errorf(name+"Pop should fail on empty stack, instead returns %v", v)
	}

	for i := 1; i <= 10; i++ {
		s.Push(i)
	}
	for i := 10; i > 0; i-- {
		if v, err := s.Top(); err == nil {
			if v != i {
				t.Errorf(name+"Top value error: %v should be %v", v, i)
			}
		} else {
			t.Errorf(name + "Top should not fail when stack not empty")
		}
		if v, err := s.Pop(); err == nil {
			if v != i {
				t.Errorf(name+"Pop value error: %v should be %v", v, i)
			}
		} else {
			t.Errorf(name + "Pop should not fail when stack not empty")
		}
	}
	if !s.IsEmpty() || s.Size() != 0 {
		t.Errorf(name+"stack should be emtpy when all data is popped, instead returns %v", s.Size())
	}

	for i := float64(1); i <= 10; i++ {
		s.Push(i)
	}
	s.Clear()
	if !s.IsEmpty() || s.Size() != 0 {
		t.Errorf(name+"stack should be emtpy when all data is cleared, instead returns %v", s.Size())
	}
}
