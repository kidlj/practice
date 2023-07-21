package algorithms

import "testing"

// https://leetcode.cn/problems/scramble-string/
func Test_isScramble(t *testing.T) {
	testCases := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{
			s1:       "abcd",
			s2:       "bcda",
			expected: true,
		},
		{
			s1:       "abcd",
			s2:       "bcad",
			expected: true,
		},
	}

	for _, tc := range testCases {
		if isScramble(tc.s1, tc.s2) != tc.expected {
			t.Errorf("Test failed: s1: %s, s2: %s, expected: %v", tc.s1, tc.s2, tc.expected)
		}
	}
}
