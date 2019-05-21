package types

import "fmt"

func Example_LeftShift() {
	// Arithmetically, a left shift x<<n is equivalent to multiplication by 2^n
	// and a right shift x>>n is equivalent to the floor of division by 2^n.
	// Go book, p54.
	// ???
	var a uint8 = 64
	var b uint8 = a << 1
	var c uint8 = a << 2
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", c)
	// Output:
	// 01000000
	// 10000000
	// 00000000
	// 64
	// 128
	// 0
}

func Example_BitSets() {
	var a uint8 = 1 << 1
	var b uint8 = 1 << 2
	var c uint8 = 1 << 5
	var x uint8 = a | b
	var y uint8 = a | c
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Printf("%08b\n", x)    // the set {1, 2}
	fmt.Printf("%08b\n", y)    // the set {1, 5}
	fmt.Printf("%08b\n", x&y)  // the intersection {1}
	fmt.Printf("%08b\n", x|y)  // the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // the difference {2}
	// Output:
	// 00000010
	// 00000100
	// 00100000
	// 00000110
	// 00100010
	// 00000010
	// 00100110
	// 00100100
	// 00000100
}
