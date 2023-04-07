package algorithms

import "testing"

func TestTreeNode(t *testing.T) {
	testCases := []struct {
		t1 *TreeNode
		t2 *TreeNode
	}{
		{
			nil,
			nil,
		},
		{
			FromSlice([]int{1, 2, 3}),
			FromSlice([]int{1, 2, 3}),
		},
		{
			FromSlice([]int{1, 7, 3, 3, 2, 5}),
			FromSlice([]int{1, 7, 3, 2, 5}),
		},
		{
			FromSlice([]int{0, -3, 9, -10, 5}),
			FromSlice([]int{0, -3, 9, 5, -10}),
		},
	}

	for _, tc := range testCases {
		if !tc.t1.Equal(tc.t2) {
			t.Errorf("t1 not equal t2")
		}
	}
}
