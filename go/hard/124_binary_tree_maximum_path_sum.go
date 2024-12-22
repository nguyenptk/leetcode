// https://leetcode.com/problems/binary-tree-maximum-path-sum

package hard

import "math"

func MaxPathSum(root *TreeNode) int {
	maxSum := math.MinInt
	dfsMaxPathSum(root, &maxSum)
	return maxSum
}

func dfsMaxPathSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	left := max(dfsMaxPathSum(root.Left, maxSum), 0)
	right := max(dfsMaxPathSum(root.Right, maxSum), 0)
	*maxSum = max(*maxSum, left+right+root.Val)

	return max(left, right) + root.Val
}
