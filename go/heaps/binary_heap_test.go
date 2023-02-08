package heaps

import "testing"

func buildBinaryHeap() *BinaryHeap {
	v := []int{5, 7, 6, 3, 8, 2, 4}
	h := BuildBinaryHeap(v)
	return h
}

func TestBuildBinaryHeap(t *testing.T) {
	h := buildBinaryHeap()
	s := h.Print()
	expected := "[2 3 4 7 8 6 5]"
	if s != expected {
		t.Errorf("BuildBinaryHeap failed, expected %s, got %s", expected, s)
	}
}

func TestBinaryHeap_DeleteMin(t *testing.T) {
	h := buildBinaryHeap()
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

func TestBinaryHeap_Insert(t *testing.T) {
	h := buildBinaryHeap()
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

func TestBinaryHeap_DecreaseKey(t *testing.T) {
	h := buildBinaryHeap()
	h.DecreaseKey(2, 1)
	expected := "[2 2 4 7 8 6 5]"
	s := h.Print()
	if s != expected {
		t.Errorf("DecreaseKey failed, expected %s, got %s", expected, s)
	}
	h.DecreaseKey(2, 1)
	expected = "[1 2 4 7 8 6 5]"
	s = h.Print()
	if s != expected {
		t.Errorf("DecreaseKey failed, expected %s, got %s", expected, s)
	}
}

func TestBinaryHeap_IncreaseKey(t *testing.T) {
	h := buildBinaryHeap()
	h.IncreaseKey(2, 1)
	expected := "[2 4 4 7 8 6 5]"
	s := h.Print()
	if s != expected {
		t.Errorf("DecreaseKey failed, expected %s, got %s", expected, s)
	}
	h.IncreaseKey(2, 4)
	expected = "[2 7 4 8 8 6 5]"
	s = h.Print()
	if s != expected {
		t.Errorf("DecreaseKey failed, expected %s, got %s", expected, s)
	}
}

func TestBinaryHeap_Delete(t *testing.T) {
	h := buildBinaryHeap()
	v, err := h.Delete(2)
	if err != nil {
		t.Errorf("Delete key 2 got error: %v", err)
	}
	if v != 3 {
		t.Errorf("Delete key 2 expected value %v, got %v", 3, v)
	}
	if h.Size() != 6 {
		t.Errorf("Delete key 2 expected heap size %v, got %v", 6, h.Size())
	}
	expected := "[2 5 4 7 8 6]"
	s := h.Print()
	if s != expected {
		t.Errorf("Delete key 2 failed, expected %s, got %s", expected, s)
	}
}

func TestBinaryHeap_FindMin(t *testing.T) {
	h := buildBinaryHeap()
	v, err := h.FindMin()
	if err != nil {
		t.Errorf("FindMin should not fail, got %v", err)
	}
	if v != 2 {
		t.Errorf("FindMin expected %v, got %v", 2, v)
	}
}
