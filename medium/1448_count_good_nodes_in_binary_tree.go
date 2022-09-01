// https://leetcode.com/problems/count-good-nodes-in-binary-tree/
package medium

import "math"

func GoodNodes(root *TreeNode) int {
	return recGoodNodes(root, math.MinInt)
}

func recGoodNodes(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}

	newMax := max
	if newMax < root.Val {
		newMax = root.Val
	}

	if root.Val >= max {
		return 1 + recGoodNodes(root.Left, newMax) + recGoodNodes(root.Right, newMax)
	}

	return recGoodNodes(root.Left, newMax) + recGoodNodes(root.Right, newMax)
}
