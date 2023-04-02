package algorithms

type listNode struct {
	val  int
	next *listNode
}

func fromSlice(s []int) *listNode {
	if len(s) == 0 {
		return nil
	}

	result := &listNode{val: s[0]}
	n := result
	for _, v := range s[1:] {
		n.next = &listNode{val: v}
		n = n.next
	}
	return result
}

func (l *listNode) toSlice() []int {
	result := []int{} // initialized value; for easy testing
	for n := l; n != nil; n = n.next {
		result = append(result, n.val)
	}
	return result
}
