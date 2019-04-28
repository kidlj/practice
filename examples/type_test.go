package examples

import (
	"fmt"
	"strconv"
)

type StringArray []string

func (a StringArray) Get(index int) string {
	return a[index]
}

// This example shows that:
// "The underlying type of a named type determines its structure and representation, and also the set of intrinsic operations it supports,
// which are the same as if the underlying type had been used directly." -- The Go Book p40
func ExampleStringArray_Get() {
	var a = StringArray{"hello", "world"}
	fmt.Println(a.Get(1))
	// Output:
	// world
}

type Errno uintptr

var errorTable = [...]string{
	"file not exist",
	"permission denied",
}

func (e Errno) Error() string {
	if int(e) > 0 && int(e) < len(errorTable) {
		// Array index must be of `integer type` or an untyped constant
		// https://golang.org/ref/spec#Index_expressions
		// `IntegerType` is here for the purposes of documentation only. It is a stand-in for any integer type: int, uint, int8 etc.
		// https://golang.org/pkg/builtin/#IntegerType
		s := errorTable[e]
		if s != "" {
			return s
		}
	}
	return "errno " + strconv.Itoa(int(e))
}

// This example shows that interface value can be compared to a concrete value that implements it,
// But I can't find this specification in Go spec doc.
// https://golang.org/ref/spec#Comparison_operators
// https://medium.com/learning-the-go-programming-language/comparing-values-in-go-8f7b002e767a
func ExampleErrno_Error() {
	e := Errno(1)
	fmt.Println(e.Error())
	var err error = e
	if err == e {
		fmt.Println("good")
	} else {
		fmt.Println("bad")
	}
	var f = "hello"
	var x interface{} = f
	if x == f {
		fmt.Println("good")
	} else {
		fmt.Println("bad")
	}
	// Output:
	// permission denied
	// good
	// good
}

// This example demostrates that slice in Go is _not_ dynamic array in
// Javascript.
// And append() is defferent to modify, append() always returns a new slice,
// with a new length, the underlying array may be still the same one.
// Updating the slice variable is required not just when calling append, but
// for any function that may change the length or capacity of a slice or make
// it refer to a different underlying array. (Go book p91)
func ExampleSliceAppendSlice() {
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
