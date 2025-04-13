package algorithms

import (
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/add-two-numbers/
func Test_addTwoNumbers(t *testing.T) {
	testCases := []struct {
		l1       *listNode
		l2       *listNode
		expected []int
	}{
		{
			l1:       fromSlice([]int{2, 4, 3}),
			l2:       fromSlice([]int{5, 6, 1}),
			expected: []int{7, 0, 5},
		},
		{
			l1:       fromSlice([]int{2, 4, 8}),
			l2:       fromSlice([]int{5, 6, 1}),
			expected: []int{7, 0, 0, 1},
		},
		{
			l1:       fromSlice([]int{0}),
			l2:       fromSlice([]int{1, 2, 3}),
			expected: []int{1, 2, 3},
		},
		{
			l1:       fromSlice([]int{1, 2, 3}),
			l2:       fromSlice([]int{0}),
			expected: []int{1, 2, 3},
		},
		{
			l1:       fromSlice([]int{0}),
			l2:       fromSlice([]int{0}),
			expected: []int{0},
		},
	}

	for _, tc := range testCases {
		c := addTwoNumbers(tc.l1, tc.l2)
		if !reflect.DeepEqual(tc.expected, c.toSlice()) {
			t.Errorf("expected %v, got %v", tc.expected, c.toSlice())
		}
	}
}
