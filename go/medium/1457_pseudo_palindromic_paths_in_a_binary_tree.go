// https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
package medium

func PseudoPalindromicPaths(root *TreeNode) int {
	result := 0
	dfsPalindromicPaths(root, 0, &result)
	return result
}

func dfsPalindromicPaths(root *TreeNode, path int, result *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		path ^= 1 << root.Val
		if (path & (path - 1)) == 0 {
			*result++
		}
		return
	}

	dfsPalindromicPaths(root.Left, path^1<<root.Val, result)
	dfsPalindromicPaths(root.Right, path^1<<root.Val, result)
}
