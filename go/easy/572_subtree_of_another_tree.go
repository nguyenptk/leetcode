package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return dfsIsSubtree(root, subRoot) || IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}

func dfsIsSubtree(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil || node2 == nil {
		return node1 == nil && node2 == nil
	}
	if node1 == nil || node2 == nil || node1.Val != node2.Val {
		return false
	}
	return dfsIsSubtree(node1.Left, node2.Left) && dfsIsSubtree(node1.Right, node2.Right)
}
