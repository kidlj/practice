package algorithms

func addTwoNumbers(l1, l2 *listNode) *listNode {
	dummy := &listNode{}
	cursor := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		l1Value := 0
		if l1 != nil {
			l1Value = l1.val
		}
		l2Value := 0
		if l2 != nil {
			l2Value = l2.val
		}

		sum := l1Value + l2Value + carry
		carry = sum / 10

		cursor.next = &listNode{val: sum % 10}
		cursor = cursor.next

		if l1 != nil {
			l1 = l1.next
		}
		if l2 != nil {
			l2 = l2.next
		}
	}

	return dummy.next
}
