package algorithms

import "fmt"

func Example_toposort() {
	order := topoSort(prereqs)
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	// Output:
	// 1:	discrete math
	// 2:	data structures
	// 3:	algorithms
	// 4:	linear algebra
	// 5:	calculus
	// 6:	formal languages
	// 7:	computer organization
	// 8:	compilers
	// 9:	databases
	// 10:	intro to programming
	// 11:	discret math
	// 12:	operating systems
	// 13:	networks
	// 14:	programming languages
}
