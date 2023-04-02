package algorithms

// merge two nondecreasing lists as one, keeping order.
func mergeLists(a, b *listNode) *listNode {
	dummy := &listNode{}
	c := dummy

	for a != nil && b != nil {
		if a.val <= b.val {
			c.next = a
			c = c.next
			a = a.next
		} else {
			c.next = b
			c = c.next
			b = b.next
		}
	}

	if a == nil {
		c.next = b
	} else {
		c.next = a
	}

	return dummy.next
}
