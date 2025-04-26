// https://leetcode.com/problems/range-sum-of-bst/
package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func RangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	result := 0
	if root.Val >= low && root.Val <= high {
		result = root.Val
	}
	if root.Val >= low {
		result += RangeSumBST(root.Left, low, high)
	}
	if root.Val <= high {
		result += RangeSumBST(root.Right, low, high)
	}

	return result
}
