package stacks

import (
	"testing"
)

func TestCalulator(t *testing.T) {
	testCalulator(t, new(Calc))
	testCalulator(t, new(ExpressionTreeCalculator))
}

func testCalulator(t *testing.T, cal Calculator) {
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
			s:      "  35  2    *",
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
		n, err := cal.Evaluate(c.s)
		if err != nil {
			if !c.errored {
				t.Errorf("expression should error: %s", c.s)
			}
			cal.Reset()
			continue
		}
		if n != c.result {
			t.Errorf("expression evaluation error: %s, expected: %f, got %f", c.s, c.result, n)
		}
		cal.Reset()
	}
}
