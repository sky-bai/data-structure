package main

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func recursion(root *TreeNode) []int {

	var per func(node *TreeNode)
	m1 := make([]int, 0)
	per = func(node *TreeNode) {
		if node == nil {
			return
		}
		m1 = append(m1, node.val)
		per(node.left)
		per(node.right)

	}

	return m1

}
