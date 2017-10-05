package algorithm

import "testing"

func Test_postorderTraversal(t *testing.T) {
	root := TreeNode{Val: 1}
	node1 := TreeNode{Val: 2}
	node2 := TreeNode{Val: 3}

	root.Right = &node1
	node1.Left = &node2

	result := postorderTraversal(&root)
	t.Log(result)
}
