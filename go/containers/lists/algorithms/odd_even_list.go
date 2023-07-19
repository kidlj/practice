package algorithms

func oddEvenList(l *listNode) *listNode {
	if l == nil || l.next == nil {
		return l
	}
	o := l
	p := l.next
	e := p
	for e != nil && e.next != nil {
		o.next = e.next
		o = o.next
		e.next = o.next
		e = e.next
	}
	o.next = p

	return l
}
