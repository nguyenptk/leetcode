// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	isOrderLeft := true

	for len(queue) > 0 {
		size := len(queue)
		group := make([]int, 0)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]

			if isOrderLeft {
				group = append(group, node.Val)
			} else {
				group = append([]int{node.Val}, group...)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			size--
		}
		result = append(result, group)
		isOrderLeft = !isOrderLeft
	}

	return result
}
