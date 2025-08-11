// https://leetcode.com/problems/binary-tree-right-side-view/
package medium

// DFS
func RightSideView(root *TreeNode) []int {
	result := make([]int, 0)

	var rec func(node *TreeNode, level int)
	rec = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(result) == level {
			result = append(result, node.Val)
		}

		rec(node.Right, level+1)
		rec(node.Left, level+1)
	}

	rec(root, 0)
	return result
}

// Time complexity: O(N) since one has to visit each node.
// Space complexity: O(H) to keep the recursion stack, where H is a tree height.

// BFS
// func RightSideView(root *TreeNode) []int {
// 	result := make([]int, 0)
// 	queue := []*TreeNode{root}

// 	for len(queue) > 0 {
// 		node := queue[len(queue)-1]

// 		if node.Left != nil {
// 			queue = append(queue, node.Left)
// 		}
// 		if node.Right != nil {
// 			queue = append(queue, node.Right)
// 		}

// 		result = append(result, node.Val)
// 	}
// 	return result
// }
