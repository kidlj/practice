package algorithms

import (
	"reflect"
	"testing"
)

func Test_removeKthFromEnd(t *testing.T) {
	testRemoveKthFromEnd(t, "removeKthFromEndRecursive", removeKthFromEnd)
	testRemoveKthFromEnd(t, "removeKthFromEndPtr", removeKthFromEndPtr)
}

func testRemoveKthFromEnd(t *testing.T, name string, f func(*listNode, int) *listNode) {
	testCases := []struct {
		l        *listNode
		k        int
		expected []int
	}{
		{
			l:        fromSlice([]int{0, 1, 2, 3, 4}),
			k:        1,
			expected: []int{0, 1, 2, 3},
		},
		{
			l:        fromSlice([]int{0, 1, 2, 3, 4}),
			k:        2,
			expected: []int{0, 1, 2, 4},
		},
		{
			l:        fromSlice([]int{0, 1, 2, 3, 4}),
			k:        5,
			expected: []int{1, 2, 3, 4},
		},
		{
			l:        fromSlice([]int{1}),
			k:        1,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		r := f(tc.l, tc.k)
		if !reflect.DeepEqual(tc.expected, r.toSlice()) {
			t.Errorf("%s: expected %v, got %v", name, tc.expected, r.toSlice())
		}
	}
}
