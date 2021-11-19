package slice

import "fmt"

// This example demostrates that slice in Go is _not_ dynamic array in
// Javascript.
// And append() is defferent to modify, append() always returns a new slice,
// with a new length, the underlying array may be still the same one.
// Updating the slice variable is required not just when calling append, but
// for any function that may change the length or capacity of a slice or make
// it refer to a different underlying array. (Go book p91)
func ExampleSliceAppend() {
	a := make([]int, 0, 10000)
	b := a
	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(b)
	// Wrong Output:
	// [1]
	// [1] ✗

	// Output:
	// [1]
	// []

	// Javascript array:
	// a = [1, 2, 3]
	// b = a
	// a.push(4)
	// console.log(a)
	// console.log(b)

	// Javascript Output:
	// [1, 2, 3, 4]
	// [1, 2, 3, 4]
}

func ExampleSliceAppendSlice() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	a = append(a, b...)
	fmt.Println(a)
	// Output:
	// [1 2 3 4 5 6]
}

func ExampleSliceModify() {
	a := []int{1, 2, 3}
	b := a
	fmt.Println(a)
	fmt.Println(b)
	a[0] = 0
	fmt.Println(a)
	fmt.Println(b)
	// Output:
	// [1 2 3]
	// [1 2 3]
	// [0 2 3]
	// [0 2 3]
}

func ExampleArrayIndex() {
	a := []int{}
	fmt.Println(a[0])
	// Output:
	// panic: index out of range
}

func ExampleMakeSlice() {
	a := make([]int, 4)
	b := make([]int, 0, 4)
	a = append(a, 1)
	b = append(b, 1)
	fmt.Println(a)
	fmt.Println(b)
	// Output:
	// [0 0 0 0 1]
	// [1]
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func ExampleReverse() {
	s := []int{0, 1, 2, 3, 4}
	reverse(s)
	fmt.Println(s)
	// Output:
	// [4 3 2 1 0]
}

func ExampleRotate() {
	// Roteate s left by two positions.
	s := []int{0, 1, 2, 3, 4}
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)
	// Output:
	// [2 3 4 0 1]
}

func ExampleFilter() {
	strings := []string{"hello", "", "world", ""}
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	fmt.Println(out)
	// Output:
	// [hello world]
}

func ExampleCopy() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6, 7}
	n := copy(a, b)
	fmt.Println(a)
	fmt.Println(n)
	// Output:
	// [4 5 6]
	// 3
}

// https://www.v2ex.com/t/606165
func ExampleIndex() {
	a := []int{}
	b := []int{0}
	var c []int
	fmt.Println(a == nil)
	fmt.Println(c == nil)
	fmt.Println(a[0:])
	fmt.Println(b[1:])
	fmt.Println(c[0:])
	// Output:
	// false
	// true
	// []
	// []
	// []
}

// cap == len by default
func ExampleSliceCap() {
	s := make([]int, 3)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	// Output:
	// 3
	// 3
}

func Example_Array() {
	a := []int{1, 2, 3}
	b := append(a, 4)
	fmt.Println(a)
	fmt.Println(b)
	// Output:
	// [1 2 3]
	// [1 2 3 4]
}
