package slice

import "testing"

func Test_binarySearch(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7}

	testCases := []struct {
		v        int
		expected int
	}{
		{
			v:        2,
			expected: 2,
		},
		{
			v:        3,
			expected: 3,
		},
		{
			v:        4,
			expected: 4,
		},
		{
			v:        7,
			expected: 7,
		},
		{
			v:        8,
			expected: -1,
		},
	}

	for _, tc := range testCases {
		r := binarySearch(s, tc.v)
		if r != tc.expected {
			t.Errorf("expected %d, got %d", tc.expected, r)
		}
	}
}
