// https://leetcode.com/problems/path-sum-iii/
package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func PathSumIII(root *TreeNode, targetSum int) int {
	count := 0
	m := map[int]int{}

	var helper func(node *TreeNode, currSum int)
	helper = func(node *TreeNode, currSum int) {
		if node == nil {
			return
		}

		currSum += node.Val
		if currSum == targetSum {
			count++
		}

		count += m[currSum-targetSum]
		m[currSum]++

		helper(node.Left, currSum)
		helper(node.Right, currSum)

		m[currSum]--
	}
	helper(root, 0)

	return count
}
