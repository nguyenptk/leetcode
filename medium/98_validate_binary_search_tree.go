// https://leetcode.com/problems/validate-binary-search-tree/
package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	return recValidate(root, nil, nil)
}

func recValidate(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && min.Val >= root.Val {
		return false
	}
	if max != nil && max.Val <= root.Val {
		return false
	}
	return recValidate(root.Left, min, root) && recValidate(root.Right, root, max)
}
