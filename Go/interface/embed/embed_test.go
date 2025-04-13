package embed

import (
	"fmt"
	"reflect"
)

func Example_embed() {
	// nil dereference panic.
	// test := Test{}
	// test.Put()

	fmt.Println(typ(Test.Put))
	// Output:
	// Test

}

func typ(t interface{}) string {
	if rt := reflect.TypeOf(t); rt.NumIn() > 0 {
		return rt.In(0).Name()
	}
	return ""
}
