package lists

import "fmt"

func ExampleArrayList_Insert() {
	var list ArrayList
	list.Insert(0, "a")
	fmt.Println(list.Size())
	// Output:
	// 1
}

func ExampleArrayList_Delete() {
	var list ArrayList
	list.Insert(0, "a")
	list.Delete(0)
	fmt.Println(list.Size())
	// Output:
	// 0
}
