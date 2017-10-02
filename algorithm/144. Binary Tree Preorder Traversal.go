package algorithm

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0, 1000)
	nodeStack := make([]*TreeNode, 0, 1000)
	if root == nil {
		return nil
	}

	iterNode := root
	for iterNode != nil || len(nodeStack) > 0 {
		for iterNode != nil {
			nodeStack = append(nodeStack, iterNode)
			result = append(result, iterNode.Val)
			iterNode = iterNode.Left
		}
		if len(result) > 0 {
			iterNode = nodeStack[len(nodeStack)-1].Right
			nodeStack = nodeStack[:len(nodeStack)-1]
		}
	}
	return result
}
