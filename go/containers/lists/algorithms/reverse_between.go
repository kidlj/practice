package algorithms

// 1 <= l <= r <= n
func reverseBetween(head *listNode, l, r int) *listNode {
	dummy := &listNode{next: head}
	pre := dummy
	cur := head

	for i := 1; i < l; i++ {
		pre = pre.next
		cur = cur.next
	}

	for i := l; i < r; i++ {
		next := cur.next
		cur.next = next.next

		next.next = pre.next
		pre.next = next
	}

	return dummy.next
}
