// https://leetcode.com/problems/binary-tree-cameras/
package hard

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Stores the minimum number of
// cameras required
var count int

func MinCameraCover(root *TreeNode) int {
	count = 0
	val := find(root)
	if val == 2 {
		return count + 1
	}
	return count
}

// DFS find
func find(root *TreeNode) int {
	if root == nil {
		return 1
	}

	l := find(root.Left)
	r := find(root.Right)

	// Both the nodes are monitored
	if l == 1 && r == 1 {
		return 2

		// If one of the left and the
		// right subtree is not monitored
	} else if l == 2 || r == 2 {
		count++
		return 3
	}

	// If the root node is monitored
	return 1
}
