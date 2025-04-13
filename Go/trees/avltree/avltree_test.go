package avltree

import (
	"fmt"
	"testing"
)

func initTree() *AVLTree {
	/*
	   13
	   /
	  8
	*/
	tr := &AVLTree{}
	tr.Insert(13)
	tr.Insert(8)
	tr.Insert(13) // no-op

	return tr
}

func TestAVLTree_Insert(t *testing.T) {
	tr := initTree()
	if tr.Size() != 2 {
		t.Errorf("Expect tree size be 2, got %d", tr.Size())
	}
	if tr.Height() != 1 {
		t.Errorf("Expect tree height be 1, got %d", tr.Height())
	}

	// Left-Left
	/*
				13		  8
				/	->   / \
			   8        0  13
		      /
			 0
	*/
	tr.Insert(0)
	if tr.Size() != 3 {
		t.Errorf("Expect tree size be 3, got %d", tr.Size())
	}
	if tr.Height() != 1 {
		t.Errorf("Expect tree height be 1, got %d", tr.Height())
	}

	// Left-Right
	/*
			8 			  8
		   /  \	        /   \
		  0    13  ->  0    12
		  	  /			   /  \
			 11			  11   13
			  \
			   12
	*/
	tr.Insert(11)
	tr.Insert(12)
	if tr.Size() != 5 {
		t.Errorf("Expect tree size be 5, got %d", tr.Size())
	}
	if tr.Height() != 2 {
		t.Errorf("Expect tree height be 2, got %d", tr.Height())
	}

	// Right-Left
	/*
		       8				 11
			  /  \				/  \
			 0    12		   8    12
			     /  \   ->	  / \     \
			   11   13		 0  10    13
			   /
			 10
	*/
	tr.Insert(10)
	if tr.Size() != 6 {
		t.Errorf("Expect tree size be 6, got %d", tr.Size())
	}
	if tr.Height() != 2 {
		t.Errorf("Expect tree height be 2, got %d", tr.Height())
	}

	// Right-Right
	/*
			11					 11
		   /  \				   /   \
		  8    12    -> 	  8     13
		 / \    \            / \   /  \
		0  10   13			0  10 12  14
		          \
			      14
	*/
	tr.Insert(14)
	if tr.Size() != 7 {
		t.Errorf("Expect tree size be 7, got %d", tr.Size())
	}
	if tr.Height() != 2 {
		t.Errorf("Expect tree height be 2, got %d", tr.Height())
	}

	output := ""
	print := func(v int) { output = fmt.Sprintf("%s %v", output, v) }
	tr.VisitDepthOrder(print)
	if output != " 11 8 13 0 10 12 14" {
		t.Errorf("Tree shape is not right! Depth-order travesal got %s", output)
	}
}
