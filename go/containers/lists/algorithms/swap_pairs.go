package algorithms

/*
	  dummy
		O --> o --> o --> o --> o
	    p     o  	e     n
*/
func swapPairs(l *listNode) *listNode {
	dummy := &listNode{next: l}
	p := dummy
	for p.next != nil && p.next.next != nil {
		o := p.next
		e := p.next.next
		n := p.next.next.next
		p.next = e
		e.next = o
		o.next = n
		p = o
	}

	return dummy.next
}

/*
		 o --> o --> o --> o
	     l     n
*/
func swapPairsRecursive(l *listNode) *listNode {
	if l == nil || l.next == nil {
		return l
	}
	n := l.next
	l.next = swapPairsRecursive(n.next)
	n.next = l
	return n
}
