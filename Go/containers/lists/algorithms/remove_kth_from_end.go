package algorithms

func removeKthFromEnd(head *listNode, k int) *listNode {
	cursor := 0
	var f func(*listNode, int) *listNode
	f = func(head *listNode, k int) *listNode {
		if head == nil {
			return nil
		}
		head.next = f(head.next, k)
		cursor++
		if cursor == k {
			return head.next
		}
		return head
	}
	return f(head, k)
}

func removeKthFromEndPtr(head *listNode, k int) *listNode {
	fast := head
	slow := head
	for i := 0; i < k; i++ {
		fast = fast.next
	}

	if fast == nil {
		return head.next
	}

	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}

	slow.next = slow.next.next
	return head
}
