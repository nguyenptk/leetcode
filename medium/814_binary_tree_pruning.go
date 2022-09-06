// https://leetcode.com/problems/binary-tree-pruning/
package medium

func PruneTree(root *TreeNode) *TreeNode {
	if postOrderPrune(root) {
		return root
	}
	return nil
}

func postOrderPrune(root *TreeNode) bool {
	if root == nil {
		return false
	}

	left := postOrderPrune(root.Left)
	right := postOrderPrune(root.Right)

	if !left {
		root.Left = nil
	}

	if !right {
		root.Right = nil
	}

	return root.Val == 1 || left || right
}
