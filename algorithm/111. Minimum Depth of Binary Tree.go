package algorithm

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	return backtrace(root)
}

func backtrace(root *TreeNode) int {

	if root == nil {
		return 0
	}

	depMin1 := backtrace(root.Left)
	depMin2 := backtrace(root.Right)

	m := int(math.Min(float64(depMin1), float64(depMin2)))
	k := 0
	if m > 0 {
		k = m
	} else {
		k = int(math.Max(float64(depMin1), float64(depMin2)))
	}
	return 1 + k

}
