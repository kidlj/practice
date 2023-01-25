package stacks

import (
	"fmt"
	"strconv"
)

type Calc struct {
	stack []float64
}

func (c *Calc) pop() (float64, error) {
	if len(c.stack) == 0 {
		return -1, fmt.Errorf("pop: stack empty")
	}
	result := c.stack[len(c.stack)-1]
	c.stack = c.stack[:len(c.stack)-1]
	return result, nil

}

func (c *Calc) push(n float64) {
	c.stack = append(c.stack, n)
}

func (c *Calc) size() int {
	return len(c.stack)
}

func (c *Calc) getOP(s string, i int) (string, int) {
	for ; i < len(s) && isSpace(s[i]); i++ {
		continue
	}
	if i < len(s) && !isDigit(s[i]) {
		c := s[i]
		i++
		return string(c), i
	}
	op := []byte{}
	for ; i < len(s) && isDigit(s[i]); i++ {
		op = append(op, s[i])
	}
	if i < len(s) && s[i] == '.' {
		op = append(op, s[i])
		i++
	}
	for ; i < len(s) && isDigit(s[i]); i++ {
		op = append(op, s[i])
	}
	return string(op), i
}

var syntaxErr = fmt.Errorf("bad syntax")

func (c *Calc) Cal(s string) (float64, error) {
	i := 0
	for i < len(s) {
		op := ""
		op, i = c.getOP(s, i)
		switch op {
		case "+":
			op1, err := c.pop()
			if err != nil {
				return -1, syntaxErr
			}
			op2, err := c.pop()
			if err != nil {
				return -1, syntaxErr
			}
			c.push(op2 + op1)
		case "-":
			op1, err := c.pop()
			if err != nil {
				return -1, syntaxErr
			}
			op2, err := c.pop()
			if err != nil {
				return -1, syntaxErr
			}
			c.push(op2 - op1)
		case "*":
			op1, err := c.pop()
			if err != nil {
				return -1, syntaxErr
			}
			op2, err := c.pop()
			if err != nil {
				return -1, syntaxErr
			}
			c.push(op2 * op1)
		case "/":
			op1, err := c.pop()
			if err != nil {
				return -1, syntaxErr
			}
			op2, err := c.pop()
			if err != nil {
				return -1, syntaxErr
			}

			if op1 == 0 {
				return -1, fmt.Errorf("zero divisor")
			}
			c.push(op2 / op1)
		default:
			n, err := strconv.ParseFloat(op, 64)
			if err != nil {
				return -1, syntaxErr
			}
			c.push(n)
		}
	}
	result, err := c.pop()
	if err != nil {
		return -1, syntaxErr
	}
	if c.size() != 0 {
		return -1, syntaxErr
	}

	return result, nil
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t'
}
