package stacks

import (
	"fmt"
	"strconv"
)

type node struct {
	val   string
	left  *node
	right *node
}

type ExpressionTreeCalculator struct {
	stack []*node
}

func (t *ExpressionTreeCalculator) push(n *node) {
	t.stack = append(t.stack, n)
}

func (t *ExpressionTreeCalculator) pop() (*node, error) {
	if len(t.stack) == 0 {
		return nil, fmt.Errorf("pop: stack empty")
	}
	result := t.stack[len(t.stack)-1]
	t.stack = t.stack[:len(t.stack)-1]
	return result, nil
}

func (t *ExpressionTreeCalculator) size() int {
	return len(t.stack)
}

func (t *ExpressionTreeCalculator) Parse(expr string) (*node, error) {
	i := 0
	for i < len(expr) {
		op := ""
		op, i = getOP(expr, i)

		switch op {
		case "+", "-", "*", "/":
			right, err := t.pop()
			if err != nil {
				return nil, syntaxErr
			}
			left, err := t.pop()
			if err != nil {
				return nil, syntaxErr
			}
			n := &node{val: op}
			n.left = left
			n.right = right
			t.push(n)
		default:
			n := &node{val: op}
			t.push(n)
		}
	}
	root, err := t.pop()
	if err != nil {
		return nil, syntaxErr
	}
	if t.size() != 0 {
		return nil, syntaxErr
	}

	return root, nil
}

func (t *ExpressionTreeCalculator) Evaluate(expr string) (float64, error) {
	root, err := t.Parse(expr)
	if err != nil {
		return -1, err
	}
	v, err := t.evaluate(root)
	if err != nil {
		return -1, err
	}

	return v, nil
}

func (t *ExpressionTreeCalculator) evaluate(n *node) (float64, error) {
	if n == nil {
		return 0, nil
	}
	left, err := t.evaluate(n.left)
	if err != nil {
		return -1, syntaxErr
	}
	right, err := t.evaluate(n.right)
	if err != nil {
		return -1, syntaxErr
	}
	op := n.val
	switch op {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		if right == 0 {
			return -1, fmt.Errorf("zero divisor")
		}
		return left / right, nil
	default:
		n, err := strconv.ParseFloat(op, 64)
		if err != nil {
			return -1, syntaxErr
		}
		return n, nil
	}
}

func (t *ExpressionTreeCalculator) Reset() {
	t.stack = nil
}
