package algorithms

func purgeDuplicated(head *listNode) *listNode {
	if head == nil || head.next == nil {
		return head
	}

	dummy := &listNode{next: head}
	pre := dummy
	cur := head
	for cur != nil && cur.next != nil {
		if cur.val == cur.next.val {
			for cur.next != nil && cur.val == cur.next.val {
				cur = cur.next
			}
			pre.next = cur.next
			cur = cur.next
		} else {
			pre = cur
			cur = cur.next
		}
	}

	return dummy.next
}
