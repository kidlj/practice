package main

import "fmt"

// IntSet is an user defined type.
type IntSet struct{}

func (s IntSet) String() string {
	return ""
}

func (s *IntSet) Error() string {
	return ""
}

func main() {
	var s IntSet

	var _ fmt.Stringer = s

	// This works,
	// &s 可用于 error 接口
	var _ fmt.Stringer = &s

	// This won't compile,
	// s 不可用于 error 接口.
	// var _ error = s

	var _ error = &s

	fmt.Println("OK")
}
