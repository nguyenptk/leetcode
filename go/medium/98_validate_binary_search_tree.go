// https://leetcode.com/problems/validate-binary-search-tree/
package medium

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	return recValidate(root, math.MinInt, math.MaxInt)
}

func recValidate(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if min >= root.Val {
		return false
	}

	if max <= root.Val {
		return false
	}

	return recValidate(root.Left, min, root.Val) &&
		recValidate(root.Right, root.Val, max)
}
