package algorithms

import "testing"

func Test_generateTrees(t *testing.T) {
	testCases := []struct {
		n        int
		expected []*TreeNode
	}{
		{
			n: 2,
			expected: []*TreeNode{
				FromSlice([]int{1, 2}),
				FromSlice([]int{2, 1}),
			},
		},
		{
			n: 3,
			expected: []*TreeNode{
				FromSlice([]int{1, 2, 3}),
				FromSlice([]int{1, 3, 2}),
				FromSlice([]int{2, 1, 3}),
				FromSlice([]int{3, 1, 2}),
				FromSlice([]int{3, 2, 1}),
			},
		},
	}

	for _, tc := range testCases {
		r := generateTrees(tc.n)
		if len(r) != len(tc.expected) {
			t.Errorf("expected length %d, got %d", len(tc.expected), len(r))
		}

		for i, n := range r {
			if !n.Equal(tc.expected[i]) {
				t.Errorf("tree not equal")
			}
		}
	}
}
