// https://leetcode.com/problems/count-good-nodes-in-binary-tree/
package medium

func GoodNodes(root *TreeNode) int {
	return recGoodNodes(root, root.Val)
}

func recGoodNodes(root *TreeNode, parent int) int {
	if root == nil {
		return 0
	}

	result := 1
	max := root.Val
	if max < parent {
		max = parent
		result = 0
	}

	result += recGoodNodes(root.Left, max)
	result += recGoodNodes(root.Right, max)

	return result
}
