// https://leetcode.com/problems/path-sum/
package easy

func HasPathSum(root *TreeNode, targetSum int) bool {
	return recurPathSum(root, targetSum)
}

func recurPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	result := false
	subSum := targetSum - root.Val

	if subSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil {
		result = result || recurPathSum(root.Left, subSum)
	}
	if root.Right != nil {
		result = result || recurPathSum(root.Right, subSum)
	}

	return result
}
