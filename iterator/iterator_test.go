package iterator

import (
	"fmt"
)

func ExampleInts_IterClassic() {
	ints := Ints{1, 2, 3}
	for it := ints.IterClassic(); it.HasNext(); {
		fmt.Println(it.Next())
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExampleInts_IterClosure() {
	ints := Ints{1, 2, 3}
	it := ints.IterClosure()
	for {
		val, ok := it()
		if !ok {
			break
		}
		fmt.Println(val)
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExampleInts_IterChannel() {
	ints := Ints{1, 2, 3}
	for val := range ints.IterChannel() {
		fmt.Println(val)
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExampleInts_IterDo() {
	ints := Ints{1, 2, 3}
	ints.IterDo(func(val int) {
		fmt.Println(val)
	})
	// Output:
	// 1
	// 2
	// 3
}
