package algorithms

func deleteDuplicates(l *listNode) *listNode {
	if l == nil {
		return l
	}

	n := l
	for n.next != nil {
		if n.val == n.next.val {
			n.next = n.next.next
		} else {
			n = n.next
		}
	}

	return l
}

func deleteDuplicatesRecursive(l *listNode) *listNode {
	if l == nil || l.next == nil {
		return l
	}

	l.next = deleteDuplicatesRecursive(l.next)
	if l.val == l.next.val {
		l = l.next
	}
	return l
}
