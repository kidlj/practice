package lists

func ExamplelinkedList_printLots() {
	// 0 1 2 3 4
	list := NewLinkedList()
	_ = list.Insert(0, 3)
	_ = list.Insert(0, 1)
	_ = list.Insert(2, 4)
	_ = list.Insert(1, 2)
	_ = list.Insert(0, 0)

	// 2 3 5
	seqList := NewLinkedList()
	_ = seqList.Insert(0, 5)
	_ = seqList.Insert(0, 3)
	_ = seqList.Insert(0, 2)

	list.printLots(seqList)
	// Output:
	// 1 2 4
}
