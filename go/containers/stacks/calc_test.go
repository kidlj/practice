package stacks

import (
	"testing"
)

func TestCal_Cal(t *testing.T) {
	cases := []struct {
		s       string
		result  float64
		errored bool
	}{
		{
			s:      "3 4 +",
			result: 7,
		},
		{
			s:      "3 4 -",
			result: -1,
		},
		{
			s:      "3 4 *",
			result: 12,
		},
		{
			s:      "8 4 /",
			result: 2,
		},
		{
			s:      "3 4 + 2 -",
			result: 5,
		},
		{
			s:      "2 4 * 4 -",
			result: 4,
		},
		{
			s:      "3.5 4 +",
			result: 7.5,
		},
		{
			s:      "3.5 2 *",
			result: 7,
		},
		{
			s:      "35 2 *",
			result: 70,
		},
		{
			s:       "4 4 4",
			errored: true,
		},
		{
			s:       "",
			errored: true,
		},
		{
			s:       "3 - 5",
			errored: true,
		},
		{
			s:       "3 0 /",
			errored: true,
		},
	}
	for _, c := range cases {
		cal := new(Calc)
		n, err := cal.Cal(c.s)
		if err != nil {
			if !c.errored {
				t.Errorf("expression should error: %s", c.s)
			}
			continue
		}
		if n != c.result {
			t.Errorf("expression: %s, expected: %f, got %f", c.s, c.result, n)
		}
	}
}
