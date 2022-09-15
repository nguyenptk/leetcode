// https://leetcode.com/problems/binary-tree-level-order-traversal/
package medium

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	// Initialize 2D result array and queue
	result := [][]int{}
	q := []TreeNode{*root}

	// Iterate queue
	for len(q) > 0 {
		currLevel := []int{}
		for i := len(q); i > 0; i-- {
			node := q[0]                            // FIRST
			q = q[1:]                               // POLL
			currLevel = append(currLevel, node.Val) // Add current value to current array
			if node.Left != nil {
				q = append(q, *node.Left) // Add back node.Left to queue
			}
			if node.Right != nil {
				q = append(q, *node.Right) // Add back node.Right to queue
			}
		}
		result = append(result, currLevel) // Add current array to 2D result array
	}

	return result
}
