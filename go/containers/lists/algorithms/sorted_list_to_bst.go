package algorithms

import (
	trees "github.com/kidlj/practice/go/trees/algorithms"
)

// sorted list to perfectly balanced binary search tree
func sortedListToBST(head *listNode) *trees.TreeNode {
	if head == nil {
		return nil
	}
	if head.next == nil {
		return &trees.TreeNode{Val: head.val}
	}

	pre := head
	p := pre.next
	q := p.next

	for q != nil && q.next != nil {
		pre = pre.next
		p = p.next
		q = q.next.next
	}

	root := &trees.TreeNode{Val: p.val}
	pre.next = nil
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(p.next)

	return root
}
