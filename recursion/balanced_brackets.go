package recursion

import (
	"strings"
	"text/scanner"
)

// The book implementation
func IsBalancedRecursiveBook(s string) bool {
	var (
		scan       scanner.Scanner
		ch         rune
		isBalanced func() bool
	)
	isBalanced = func() bool {
		if ch == scanner.EOF {
			return true
		}
		if ch != '[' {
			return false
		}
		ch = scan.Next()
		if ch == '[' {
			if !isBalanced() {
				return false
			}
		}
		if ch != ']' {
			return false
		}
		ch = scan.Next()
		if ch == '[' {
			return isBalanced()
		}
		return true
	}

	scan.Init(strings.NewReader(s))
	ch = scan.Next()
	return isBalanced() && ch == scanner.EOF
}

// My implementation
func IsBalancedRecursive(s string) bool {
	var (
		scan       scanner.Scanner
		ch         rune
		isBalanced func() bool
	)

	isBalanced = func() bool {
		ch = scan.Next()
		if ch != '[' {
			return false
		}

		for isBalanced() {
		}

		if ch != ']' {
			return false
		}

		return true
	}

	// main
	scan.Init(strings.NewReader(s))

	for scan.Peek() != scanner.EOF {
		if !isBalanced() {
			return false
		}
	}

	return true
}
