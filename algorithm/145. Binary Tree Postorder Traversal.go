package algorithm

func postorderTraversal(root *TreeNode) []int {
	nodeStack := make([]*TreeNode, 0, 100000)
	result := make([]int, 0, 100000)

	if root == nil {
		return result
	}

	nodeStack = append(nodeStack, root)
	//result = append(result, root.Val)
	iterNode := root
	var prevNode *TreeNode

	for len(nodeStack) > 0 {
		iterNode = nodeStack[len(nodeStack)-1]

		//fmt.Println("val:", iterNode.Val, "nodeStack:", nodeStack)
		// 某节点没有左右节点,或者其左右节点都已经被访问过, 则可输出
		if (iterNode.Left == nil && iterNode.Right == nil) ||
			(prevNode != nil && (prevNode == iterNode.Left || prevNode == iterNode.Right)) {
			result = append(result, iterNode.Val)
			nodeStack = nodeStack[:len(nodeStack)-1]
			prevNode = iterNode
		} else {
			if iterNode.Right != nil {
				nodeStack = append(nodeStack, iterNode.Right)
			}
			if iterNode.Left != nil {
				nodeStack = append(nodeStack, iterNode.Left)
			}
		}
	}
	return result
}
