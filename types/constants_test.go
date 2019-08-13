package types

import "fmt"

func Example_PrintRune() {
	s := 'h'
	fmt.Printf("%q", s)
	// Output:
	// 'h'
}

// https://twitter.com/davecheney/status/1145275398478893056
func Example_Zero() {
	fmt.Printf("%T\n", "0"[0])
	fmt.Printf("%U\n", byte('0'))
	fmt.Printf("%d\n", "0"[0])

	fmt.Printf("%T\n", byte('0'))
	fmt.Printf("%U\n", byte('0'))

	fmt.Printf("%T\n", '0')
	fmt.Printf("%U\n", '0')

	fmt.Printf("%T\n", 0)
	fmt.Printf("%U\n", 0)
	fmt.Printf("%d\n", 0)

	fmt.Printf("%T\n", 1)
	fmt.Printf("%U\n", 1)
	fmt.Printf("%d\n", 1)

	fmt.Println(byte('0') == "0"[0])
	fmt.Println(0 < "0"[0])
	// Output:
	// uint8
	// U+0030
	// 48
	// uint8
	// U+0030
	// int32
	// U+0030
	// int
	// U+0000
	// 0
	// int
	// U+0001
	// 1
	// true
	// true
}
