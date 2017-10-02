package algorithm

import (
	"fmt"
	"testing"
)

func Test_minDepth(t *testing.T) {
	tree1Node1 := TreeNode{Val: 1}
	tree1Node2 := TreeNode{Val: 2}
	tree1Node3 := TreeNode{Val: 3}
	tree1Node4 := TreeNode{Val: 4}
	tree1Node5 := TreeNode{Val: 5}
	tree1Node6 := TreeNode{Val: 6}

	tree1Node1.Left = &tree1Node2
	tree1Node2.Left = &tree1Node3
	tree1Node2.Right = &tree1Node4
	tree1Node3.Left = &tree1Node5
	tree1Node4.Right = &tree1Node6

	res := minDepth(&tree1Node1)
	fmt.Println(res)
}
