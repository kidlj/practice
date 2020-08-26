package reflect

import (
	"fmt"
	"reflect"
	"runtime"
)

func FuncTion() {
	fmt.Println("hello")
}

func Example_func() {
	fmt.Printf("func is %v", FuncName(reflect.ValueOf(FuncTion)))
	// Output:
	// a
}

func FuncName(v reflect.Value) (name string) {
	if v.Kind() == reflect.Func {
		name = runtime.FuncForPC(v.Pointer()).Name()
		return name
	}
	return
}
