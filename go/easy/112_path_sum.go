// https://leetcode.com/problems/path-sum/
package easy

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val

	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	return HasPathSum(root.Left, targetSum) || HasPathSum(root.Right, targetSum)
}
