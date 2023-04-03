package algorithms

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	testDeleteDuplicates(t, "deleteDuplicates", deleteDuplicates)
	testDeleteDuplicates(t, "deleteDuplicatesRecursive", deleteDuplicatesRecursive)
}

func testDeleteDuplicates(t *testing.T, name string, f func(*listNode) *listNode) {
	testCases := []struct {
		l        *listNode
		expected []int
	}{
		{
			l:        fromSlice([]int{0, 1, 1, 2, 3, 3, 4}),
			expected: []int{0, 1, 2, 3, 4},
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
		r := deleteDuplicates(tc.l)
		if !reflect.DeepEqual(tc.expected, r.toSlice()) {
			t.Errorf("expected %v, got %v", tc.expected, r.toSlice())
		}
	}
}
