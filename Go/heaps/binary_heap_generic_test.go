package heaps

import "testing"

func buildGenericBinaryHeap() *GenericBinaryHeap[int] {
	v := []int{5, 7, 6, 3, 8, 2, 4}
	h := FromSlice(func(a, b int) int { return a - b }, v)
	return h
}

func TestBuildGenericBinaryHeap(t *testing.T) {
	h := buildGenericBinaryHeap()
	s := h.Print()
	expected := "[2 3 4 7 8 6 5]"
	if s != expected {
		t.Errorf("BuildBinaryHeap failed, expected %s, got %s", expected, s)
	}
}

func TestGenericBinaryHeap_DeleteMin(t *testing.T) {
	h := buildGenericBinaryHeap()
	min, err := h.DeleteMin()
	if err != nil {
		t.Errorf("DeleteMin should not error, got %v", err)
	}
	if min != 2 {
		t.Errorf("DeleteMin expected %v, got %v", 2, min)
	}
	s := h.Print()
	expected := "[3 5 4 7 8 6]"
	if s != expected {
		t.Errorf("DeleteMin failed, expected %s, got %s", expected, s)
	}
}

func TestGenericBinaryHeap_Insert(t *testing.T) {
	h := buildGenericBinaryHeap()
	h.Insert(0)
	if h.Size() != 8 {
		t.Errorf("Insert expected heap size %v, got %v", 8, h.Size())
	}
	s := h.Print()
	expected := "[0 2 4 3 8 6 5 7]"
	if s != expected {
		t.Errorf("DeleteMin failed, expected %s, got %s", expected, s)
	}
}
func TestGenericBinaryHeap_FindMin(t *testing.T) {
	h := buildGenericBinaryHeap()
	v, err := h.FindMin()
	if err != nil {
		t.Errorf("FindMin should not fail, got %v", err)
	}
	if v != 2 {
		t.Errorf("FindMin expected %v, got %v", 2, v)
	}
}
