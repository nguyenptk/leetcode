// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
package medium

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > root.Val && q.Val > root.Val {
		return LowestCommonAncestor(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return LowestCommonAncestor(root.Left, p, q)
	}

	return root
}
