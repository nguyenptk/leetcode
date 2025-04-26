// https://leetcode.com/problems/symmetric-tree/
package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsSymmetric(root *TreeNode) bool {
	var isMirror func(t1, t2 *TreeNode) bool
	isMirror = func(t1, t2 *TreeNode) bool {
		if t1 == nil && t2 == nil {
			return true
		}
		if t1 == nil || t2 == nil {
			return false
		}
		return t1.Val == t2.Val && isMirror(t1.Right, t2.Left) && isMirror(t1.Left, t2.Right)
	}
	return isMirror(root, root)
}
