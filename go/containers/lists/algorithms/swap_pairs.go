package algorithms

/*
	  dummy			n			n'
		O --> o --> o --> o --> o
	    l     r  	l'    r'
*/
func swapPairs(head *listNode) *listNode {
	dummy := &listNode{next: head}
	left := dummy
	right := dummy.next

	for left != nil && right != nil && right.next != nil {
		// swap
		left.next = right.next
		right.next = left.next.next
		left.next.next = right

		// next move
		left = right
		right = left.next
	}

	return dummy.next
}

/*
		 o --> o --> o --> o
	    head  next
*/
func swapPairsRecursive(head *listNode) *listNode {
	if head == nil || head.next == nil {
		return head
	}
	next := head.next
	head.next = swapPairsRecursive(next.next)
	next.next = head
	return next
}
