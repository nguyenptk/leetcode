// https://leetcode.com/problems/count-good-nodes-in-binary-tree/
package medium

func GoodNodes(root *TreeNode) int {
	var dfs func(root *TreeNode, curMax int) int
	dfs = func(root *TreeNode, curMax int) int {
		if root == nil {
			return 0
		}

		count := 0
		if root.Val >= curMax {
			count++
			curMax = root.Val
		}

		l := dfs(root.Left, curMax)
		r := dfs(root.Right, curMax)

		return l + r + count
	}

	return dfs(root, root.Val)
}
