// https://leetcode.com/problems/deepest-leaves-sum/
package medium

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func DeepestLeavesSum(root *TreeNode) int {
	return sumMaxLevel(root)
}

func sumMaxLevel(root *TreeNode) int {
	// call to function to calculate
	// max depth
	maxDepth := maxDepth(root)

	return sumMaxLevelRec(root, maxDepth)
}

// maxDepth function to find the
// max depth of the tree
func maxDepth(root *TreeNode) int {
	// base case
	if root == nil {
		return 0
	}

	// either leftDepth of rightDepth is
	// greater add 1 to include height
	// of node at which call is
	return 1 + int(math.Max(float64(maxDepth(root.Left)),
		float64(maxDepth(root.Right))))
}

func sumMaxLevelRec(root *TreeNode, max int) int {
	// base case
	if root == nil {
		return 0
	}

	// max == 1 to track the node
	// at deepest level
	if max == 1 {
		return root.Val
	}

	// recursive call to left and right nodes
	return sumMaxLevelRec(root.Left, max-1) + sumMaxLevelRec(root.Right, max-1)
}
