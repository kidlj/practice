package types

import (
	"fmt"
	"reflect"
)

func Example_Pointer_Equal() {
	var a *int
	var b *int
	fmt.Println(a == b)
	// fmt.Println(*a == *b) // panic
	fmt.Println(reflect.DeepEqual(a, b))

	s := 3
	q := 3
	a = &s
	b = &q
	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(a == b)
	// Output:
	// true
	// true
	// true
	// false
}

func Example_nil() {
	var a = new(int)
	fmt.Println(*a)
	fmt.Println(a == nil)
	var b *int
	// panic !!!
	*b = 3
	// Output:
	// 0
	// false
}

type S struct{}

type T struct {
	S S
}

func (s S) value() {
	fmt.Println("value")
}

func (s *S) pointer() {
	fmt.Println("pointer")
}

// func Example_receiver() {
// 	S{}.value()
// 	[]S{{}}[0].value()
// 	T{S{}}.S.value()
// 	map[string]S{"a": {}}["a"].value()

// 	S{}.pointer()                        // compiler: cannot call pointer method ptr on S
// 	[]S{{}}[0].pointer()                 // ok for slice/array literal elements
// 	T{S{}}.S.pointer()                   // compiler: cannot call pointer method ptr on S
// 	map[string]S{"a": {}}["a"].pointer() // compiler: cannot call pointer method ptr on S

// 	// Output:
// 	// value
// 	// value
// 	// value
// 	// value
// }
