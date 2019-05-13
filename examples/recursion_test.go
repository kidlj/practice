package examples

import (
	"fmt"
	"strings"
	"text/scanner"
)

// Clear is better than clever;
// This is not clear.
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
			fmt.Printf("catch: %s\n", string(ch))
			return false
		}
		ch = scan.Next()
		fmt.Printf("%s\n", string(ch))
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

func ExampleIsBalancedRecursive() {
	a := []string{
		"",
		"[]",
		"[][]",
		"[[]]",
		"[[]][]",
		"[[[[][]]][]]",
		"[",
		"]",
		"[[",
		"][",
		"]]",
		"[]]",
		"[a]",
	}

	for _, s := range a {
		fmt.Println(IsBalancedRecursive(s))
	}
	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// false
	// false
	// false
	// false
	// false
	// false
	// false
}
