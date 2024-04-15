// https://leetcode.com/problems/binary-tree-level-order-traversal/
package medium

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	q := []TreeNode{*root}

	for len(q) > 0 {
		level := make([]int, 0)
		for i := len(q); i > 0; i-- {
			node := q[0] // First
			q = q[1:]    // Poll
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, *node.Left)
			}
			if node.Right != nil {
				q = append(q, *node.Right)
			}
		}
		result = append(result, level)
	}

	return result
}
