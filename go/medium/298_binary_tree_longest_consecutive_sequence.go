// https://leetcode.com/problems/binary-tree-longest-consecutive-sequence/
package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func LongestConsecutiveSequence(root *TreeNode) int {
	maxLen := 0
	var dfs func(node *TreeNode, expected, length int)
	dfs = func(node *TreeNode, expected, length int) {
		if node == nil {
			return
		}
		if node.Val == expected {
			length++
		} else {
			length = 1
		}
		maxLen = max(maxLen, length)
		dfs(node.Left, node.Val+1, length)
		dfs(node.Right, node.Val+1, length)
	}

	dfs(root, root.Val, 0)
	return maxLen
}
