// https://leetcode.com/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level/
package medium

import (
	"sort"
)

func MinimumOperations(root *TreeNode) int {
	queue := []*TreeNode{root}
	totalSwaps := 0

	// BFS to process tree level by level
	for len(queue) > 0 {
		levelSize := len(queue)
		levelValues := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]  // Get front of the queue (FIFO)
			queue = queue[1:] // Remove front element
			levelValues[i] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Calculate minimum swaps for the current level
		totalSwaps += getMinSwaps(levelValues)
	}
	return totalSwaps
}

func getMinSwaps(original []int) int {
	swaps := 0
	target := make([]int, len(original))
	copy(target, original)
	sort.Ints(target)

	// Position map for current indices of values
	pos := make(map[int]int, len(original))
	for i, val := range original {
		pos[val] = i
	}

	// Swap elements until array matches the target
	for i := 0; i < len(original); i++ {
		for original[i] != target[i] {
			swaps++
			currPos := pos[target[i]]
			pos[original[i]] = currPos
			pos[target[i]] = i
			original[currPos], original[i] = original[i], original[currPos]
		}
	}
	return swaps
}
