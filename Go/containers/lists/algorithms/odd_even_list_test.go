package algorithms

import (
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/odd-even-linked-list/
func Test_oddEvenList(t *testing.T) {
	testCases := []struct {
		l        *listNode
		expected []int
	}{
		{
			l:        fromSlice([]int{0, 1, 2, 3, 4}),
			expected: []int{0, 2, 4, 1, 3},
		},
		{
			l:        fromSlice([]int{0, 1}),
			expected: []int{0, 1},
		},
		{
			l:        fromSlice([]int{}),
			expected: nil,
		},
		{
			l:        fromSlice([]int{0}),
			expected: []int{0},
		},
	}

	for _, tc := range testCases {
		result := oddEvenList(tc.l)
		if !reflect.DeepEqual(tc.expected, result.toSlice()) {
			t.Errorf("expected %v, got %v", tc.expected, result.toSlice())
		}
	}
}
