package algorithms

import "testing"

func Test_maxDepth(t *testing.T) {
	testCases := []struct {
		t        *TreeNode
		expected int
	}{
		{
			t:        FromSlice([]int{0, -3, 9, -10, 5}),
			expected: 3,
		},
		{
			t:        nil,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		r := maxDepth(tc.t)
		if r != tc.expected {
			t.Errorf("expected %v, got %v", tc.expected, r)
		}
	}
}
