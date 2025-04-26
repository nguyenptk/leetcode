// https://leetcode.com/problems/binary-tree-right-side-view/
package medium

func RightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(0, root, &result)
	return result
}

func dfs(depth int, root *TreeNode, result *[]int) {
	if root != nil {
		if depth == len(*result) {
			*result = append(*result, root.Val)
		}
		dfs(depth+1, root.Right, result)
		dfs(depth+1, root.Left, result)
	}
}

func bfs(root *TreeNode, result *[]int) {
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[len(queue)-1]

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		*result = append(*result, node.Val)
	}
}
