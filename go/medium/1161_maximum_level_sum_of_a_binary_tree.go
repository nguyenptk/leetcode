// https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
package medium

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func MaxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := math.MinInt
	result := 1
	level := 1

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelSum := 0
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if levelSum > maxSum {
			maxSum = levelSum
			result = level
		}
		level++
	}

	return result
}
