// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func LowestCommonAncestorBTree(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	l := LowestCommonAncestorBTree(root.Left, p, q)
	r := LowestCommonAncestorBTree(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}

	if l != nil {
		return l
	}

	return r
}
