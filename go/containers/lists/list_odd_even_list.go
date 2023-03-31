package lists

func (l *linkedList) oddEvenList() *linkedList {
	if l.head.succ == nil || l.head.succ.succ == nil {
		return l
	}
	o := l.head.succ
	p := l.head.succ.succ
	e := p
	for o.succ != nil && e.succ != nil {
		o.succ = e.succ
		o = o.succ
		e.succ = o.succ
		e = e.succ
	}
	o.succ = p

	return l
}
