package algorithms

func deleteDuplicates(head *listNode) *listNode {
	if head == nil || head.next == nil {
		return head
	}

	n := head
	for n.next != nil {
		if n.val == n.next.val {
			n.next = n.next.next
		}
		n = n.next
	}

	return head
}

func deleteDuplicatesRecursive(head *listNode) *listNode {
	if head == nil || head.next == nil {
		return head
	}

	head.next = deleteDuplicatesRecursive(head.next)
	if head.val == head.next.val {
		head = head.next
	}
	return head
}
