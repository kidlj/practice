package binarysearchtree

import (
	"fmt"
	"testing"
)

func initTree() *BinarySearchTree {
	tr := &BinarySearchTree{}
	/*
				 5
				/ \
			   /   \
			  3     7
			 / \   / \
			2   4 6   8
		   /
		  1
	*/
	tr.Insert(5)
	tr.Insert(3)
	tr.Insert(7)
	tr.Insert(2)
	tr.Insert(4)
	tr.Insert(6)
	tr.Insert(8)
	tr.Insert(1)

	return tr
}

func TestBinarySearchTree_Insert(t *testing.T) {
	tr := initTree()
	if tr.Size() != 8 {
		t.Errorf("Expecte tree size == 8, got %d", tr.Size())
	}
	if tr.Height() != 3 {
		t.Errorf("Expect tree size == 3, got %d", tr.Height())
	}
	tr.Insert(3)
	if tr.Size() != 8 {
		t.Errorf("Insert duplicated elements should do noting")
	}
}

func TestBinarySearchTree_Delete(t *testing.T) {
	tr := initTree()
	tr.Delete(9)
	if tr.Size() != 8 {
		t.Errorf("delete a non-existing element should do nothing")
	}
	tr.Delete(3)
	if tr.Size() != 7 {
		t.Errorf("delete an existing element should decrease tree size by 1")
	}
	if tr.Height() != 3 {
		t.Errorf("Expect tree size == 3, got %d", tr.Height())
	}
	tr.Delete(2)
	if tr.Size() != 6 {
		t.Errorf("delete an existing element should decrease tree size by 1")
	}
	if tr.Height() != 2 {
		t.Errorf("Expect tree size == 2, got %d", tr.Height())
	}
	tr.Delete(1)
	if tr.Size() != 5 {
		t.Errorf("delete an existing element should decrease tree size by 1")
	}
	if tr.Height() != 2 {
		t.Errorf("Expect tree size == 2, got %d", tr.Height())
	}
}

func TestBinarySearchTree_Find(t *testing.T) {
	tr := initTree()
	n := tr.Find(3)
	if n == nil {
		t.Errorf("3 should be in tree")
		return
	}
	if n.val != 3 {
		t.Errorf("Found 3, but its value is %d", n.val)
	}
}

func TestBinarySearchTree_FindMin(t *testing.T) {
	tr := initTree()
	n := tr.FindMin()
	if n == nil {
		t.Errorf("min should not be nil")
		return
	}
	if n.val != 1 {
		t.Errorf("Found min node and its value should be 1, but got %d", n.val)
	}
}

func ExampleBinarySearchTree_VisitPreOrder() {
	tr := initTree()
	print := func(v int) { fmt.Printf("%d ", v) }
	tr.VisitPreOrder(print)
	// Output:
	// 5 3 2 1 4 7 6 8
}

func ExampleBinarySearchTree_NewInOrderIterator() {
	tr := initTree()
	iter := tr.NewInOrderIterator()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("%v", iter.Done())
	// Output:
	// 1 2 3 4 5 6 7 8 true
}

func ExampleBinarySearchTree_NewPreOrderIterator() {
	tr := initTree()
	iter := tr.NewPreOrderIterator()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("%v", iter.Done())
	// Output:
	// 5 3 2 1 4 7 6 8 true
}

func ExampleBinarySearchTree_NewDepthOrderIterator() {
	tr := initTree()
	iter := tr.NewDepthOrderIterator()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		fmt.Printf("%d ", v)

	}
	fmt.Printf("%v", iter.Done())
	// Output:
	// 5 3 7 2 4 6 8 1 true
}
