package sorting

import (
	"reflect"
	"testing"
)

func TestSortings(t *testing.T) {
	funcs := map[string]func([]int){
		"SelectionSort": SelectionSort,
		"BubbleSort":    BubbleSort,
		"InsertionSort": InsertionSort,
		"ShellSort":     ShellSort,
		"HeapSort":      HeapSort,
		"QuickSort":     QuickSort,
		"MergeSort":     MergeSort,
	}

	for name, f := range funcs {
		testSorting(t, name, f)
	}
}

func testSorting(t *testing.T, name string, f func([]int)) {
	testCases := []struct {
		src      []int
		expected []int
	}{
		{
			src:      []int{0, 1},
			expected: []int{0, 1},
		},
		{
			src:      []int{1, 1, 1, 1, 1},
			expected: []int{1, 1, 1, 1, 1},
		},
		{
			src:      []int{3, 4, 7, 9, 2, 6, 5, 8, 1, 0},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			src:      []int{0, 0, 0, 0, 0, 1},
			expected: []int{0, 0, 0, 0, 0, 1},
		},
		{
			src:      []int{0, 0, 1, 0, 0, 1},
			expected: []int{0, 0, 0, 0, 1, 1},
		},
	}
	for _, tc := range testCases {
		f(tc.src)
		if !reflect.DeepEqual(tc.src, tc.expected) {
			t.Errorf("%s: expected %v, got %v", name, tc.expected, tc.src)
		}
	}
}
