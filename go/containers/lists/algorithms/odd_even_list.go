package algorithms

func oddEvenList(head *listNode) *listNode {
	if head == nil || head.next == nil {
		return head
	}
	o := head
	p := head.next
	e := p
	for o.next != nil && e.next != nil {
		o.next = e.next
		o = o.next
		e.next = o.next
		e = e.next
	}
	o.next = p

	return head
}
