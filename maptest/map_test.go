package maptest

import "fmt"

func Example_map() {
	states := make(map[int]int)
	for i := 0; i < 3; i++ {
		states[i]++
	}
	fmt.Println(states)
	// Output:
	// bad
}

func Trim(n int) int {
	states := make(map[int]int)
	for i := 0; i < n; i++ {
		state := Trim(i)
		states[state]++
	}
	fmt.Println("post states:", states)
	return n
}

func Example_Trim() {
	Trim(2)
	// Output:
	// bad
}
