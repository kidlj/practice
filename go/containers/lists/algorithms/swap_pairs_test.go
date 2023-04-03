package algorithms

import (
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	testSwapPairs(t, "swapPairs", swapPairs)
	testSwapPairs(t, "swapPairsRecursive", swapPairsRecursive)
}

func testSwapPairs(t *testing.T, name string, f func(*listNode) *listNode) {
	testCases := []struct {
		l        *listNode
		expected []int
	}{
		{
			l:        fromSlice([]int{1, 2, 3, 4, 5, 6}),
			expected: []int{2, 1, 4, 3, 6, 5},
		},
		{
			l:        fromSlice([]int{1, 2, 3, 4, 5}),
			expected: []int{2, 1, 4, 3, 5},
		},
		{
			l:        fromSlice([]int{1}),
			expected: []int{1},
		},
		{
			l:        nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		r := f(tc.l)
		if !reflect.DeepEqual(tc.expected, r.toSlice()) {
			t.Errorf("%s: expected %v, got %v", name, tc.expected, r.toSlice())
		}
	}
}
