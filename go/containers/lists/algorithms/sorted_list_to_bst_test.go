package algorithms

import (
	"testing"

	trees "github.com/kidlj/practice/go/trees/algorithms"
)

func Test_sortedListToBST(t *testing.T) {
	testCases := []struct {
		l        *listNode
		expected *trees.TreeNode
	}{
		{
			l:        fromSlice([]int{-10, -3, 0, 5, 9}),
			expected: trees.FromSlice([]int{0, -3, 9, -10, 5}),
		},
		{
			l:        nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		r := sortedListToBST(tc.l)
		if !tc.expected.Equal(r) {
			t.Errorf("not equal")
		}
	}
}
