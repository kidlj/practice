package algorithms

type listNode struct {
	val  int
	next *listNode
}

func fromSlice(s []int) *listNode {
	dummy := &listNode{}
	n := dummy
	for _, v := range s {
		n.next = &listNode{val: v}
		n = n.next
	}
	return dummy.next
}

func (l *listNode) toSlice() (result []int) {
	for n := l; n != nil; n = n.next {
		result = append(result, n.val)
	}
	return
}
