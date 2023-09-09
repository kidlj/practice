package types

import (
	"fmt"
	"strconv"
)

type StringArray []string

func (a StringArray) Get(index int) string {
	return a[index]
}

func (a StringArray) String() string {
	return "hello"
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

func ExampleStructLiteral() {
	type kidlj []string

	blog := "blog"
	read := "read"
	wiki := "wiki"

	site := kidlj{blog, read, wiki}
	fmt.Println(site)
	// Output:
	// [blog read wiki]
}

func Example_Nil_Slice() {
	var s []int
	fmt.Printf("s == nil: %t\n", s == nil)
	// Output:
	// s == nil: true
}

func Example_Nil_Pointer() {
	var s StringArray
	fmt.Printf("s == nil: %t\n", s == nil)
	fmt.Println(s.String())
	// Output:
	// s == nil: true
	// hello
}

type Employee struct {
	ID     int
	Name   string
	Salary int
}

func EmployeeByID(id int) Employee {
	return Employee{
		ID:     1,
		Name:   "m",
		Salary: 100,
	}
}

func (e *Employee) SetSalary(salary int) {
	e.Salary = salary
}

// gopl p100.
func Example_Value_Variable() {
	// cannot assign to EmployeeByID(3).Salary (value of type int) compiler
	// EmployeeByID(3).Salary = 0
}

func Example_AutoPointer() {
	e := Employee{}
	e.SetSalary(100)
	fmt.Println(e.Salary)
	// Output:
	// 100
}
