package algorithms

func reverseList(head *listNode) *listNode {
	var rev *listNode
	for head != nil {
		next := head.next
		head.next = rev
		rev = head
		head = next
	}
	return rev
}
