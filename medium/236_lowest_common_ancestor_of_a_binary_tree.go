// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
package medium

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	l := LowestCommonAncestor(root.Left, p, q)
	r := LowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	} else if l == nil && r == nil {
		return nil
	} else {
		if l == nil {
			return r
		}
		return l
	}
}
