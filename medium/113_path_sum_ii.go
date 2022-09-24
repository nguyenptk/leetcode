// https://leetcode.com/problems/path-sum-ii/
package medium

func PathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	dfsPathSum(root, targetSum, []int{}, &result)
	return result
}

func dfsPathSum(root *TreeNode, targetSum int, path []int, result *[][]int) {
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		*result = append(*result, append(path, root.Val))
		return
	}
	path = append(append([]int{}, path...), root.Val)
	if root.Left != nil {
		dfsPathSum(root.Left, targetSum-root.Val, path, result)
	}
	if root.Right != nil {
		dfsPathSum(root.Right, targetSum-root.Val, path, result)
	}
}
