package algorithms

import (
	"reflect"
	"testing"

	trees "github.com/kidlj/practice/go/trees/algorithms"
)

func Test_flatten(t *testing.T) {
	testCases := []struct {
		t        *trees.TreeNode
		expected []int
	}{
		{
			t:        trees.FromSlice([]int{0, -3, 9, 5, -10}),
			expected: []int{0, -3, -10, 9, 5},
		},
		{
			t:        nil,
			expected: nil,
		},
		{
			t:        trees.FromSlice([]int{1}),
			expected: []int{1},
		},
	}

	for _, tc := range testCases {
		flatten(tc.t)
		if !reflect.DeepEqual(tc.t.RightMost(), tc.expected) {
			t.Errorf("expected %v, got %v", tc.expected, tc.t.RightMost())
		}
	}
}
