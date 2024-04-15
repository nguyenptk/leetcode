package easy

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BalancedTree struct {
	Balance bool
	Height  int
}

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfsIsBalanced(root).Balance
}

func dfsIsBalanced(root *TreeNode) BalancedTree {
	if root == nil {
		return BalancedTree{true, 0}
	}
	left := dfsIsBalanced(root.Left)
	right := dfsIsBalanced(root.Right)
	balance := left.Balance && right.Balance && math.Abs(float64(left.Height-right.Height)) <= 1
	height := 1 + max(left.Height, right.Height)
	return BalancedTree{balance, height}
}
