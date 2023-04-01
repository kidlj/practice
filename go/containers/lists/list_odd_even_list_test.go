package lists

import "fmt"

func Example_linkedList_oddEvenList() {
	// 0 1 2 3 4
	list := NewLinkedList()
	_ = list.Insert(0, 3)
	_ = list.Insert(0, 1)
	_ = list.Insert(2, 4)
	_ = list.Insert(1, 2)
	_ = list.Insert(0, 0)

	l := list.oddEvenList()
	fmt.Println(l.String())
	// Output:
	// 0 2 4 1 3
}
