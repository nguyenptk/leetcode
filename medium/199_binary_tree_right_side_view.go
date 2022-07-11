// https://leetcode.com/problems/binary-tree-right-side-view/
package medium

var result []int

func RightSideView(root *TreeNode) []int {
	result = []int{}
	dfs(0, root)
	return result
}

func dfs(depth int, root *TreeNode) {
	if root != nil {
		if depth == len(result) {
			result = append(result, root.Val)
		}
		dfs(depth+1, root.Right)
		dfs(depth+1, root.Left)
	}
}
