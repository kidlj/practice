package algorithms

import trees "github.com/kidlj/practice/go/trees/algorithms"

// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
// - 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null。
// - 展开后的单链表应该与二叉树 先序遍历 顺序相同。
func flatten(root *trees.TreeNode) {
	var last *trees.TreeNode

	var f func(*trees.TreeNode)
	f = func(root *trees.TreeNode) {
		if root == nil {
			return
		}

		f(root.Right)
		f(root.Left)

		root.Left = nil
		root.Right = last
		last = root
	}

	f(root)
}
