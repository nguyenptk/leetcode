// https://leetcode.com/problems/find-largest-value-in-each-tree-row/
package medium

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func LargestValues(root *TreeNode) []int {
	result := make([]int, 0)
	// dfsLargestValues(root, &result, 0)
	bfsLargestValues(root, &result)
	return result
}

func dfsLargestValues(root *TreeNode, result *[]int, depth int) {
	if root == nil {
		return
	}

	if len(*result) == depth {
		*result = append(*result, root.Val)
	} else {
		(*result)[depth] = max((*result)[depth], root.Val)
	}

	dfsLargestValues(root.Left, result, depth+1)
	dfsLargestValues(root.Right, result, depth+1)
}

func bfsLargestValues(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		currLength := len(queue)
		currMax := math.MinInt32

		for i := 0; i < currLength; i++ {
			node := queue[0]
			queue = queue[1:]
			currMax = max(currMax, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		*result = append(*result, currMax)
	}
}
