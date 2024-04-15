package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	return compare(p, q)
}

func compare(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return compare(p.Left, q.Left) && compare(p.Right, q.Right)
}
