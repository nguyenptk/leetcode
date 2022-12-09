// https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/
package medium

func MaxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return findMax(root, root.Val, root.Val)
}

func findMax(node *TreeNode, currMax, currMin int) int {
	if node == nil {
		return currMax - currMin
	}

	if currMax < node.Val {
		currMax = node.Val
	}
	if currMin > node.Val {
		currMin = node.Val
	}

	l := findMax(node.Left, currMax, currMin)
	r := findMax(node.Right, currMax, currMin)

	if l < r {
		return r
	}
	return l
}
