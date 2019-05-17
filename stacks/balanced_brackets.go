package stacks

import (
	"strings"
	"text/scanner"
)

func IsBalancedStack(s string) bool {
	var scan scanner.Scanner
	var stack ArrayStack

	scan.Init(strings.NewReader(s))

	for ch := scan.Next(); ch != scanner.EOF; ch = scan.Next() {
		if ch == '[' {
			stack.Push(ch)
		} else if ch == ']' {
			if stack.IsEmpty() {
				return false
			}
			stack.Pop()
		} else {
			return false
		}
	}

	return stack.IsEmpty()
}
