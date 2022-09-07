// https://leetcode.com/problems/606/
package easy

import "strconv"

func Tree2str(root *TreeNode) string {
	return dfsTree2str(root)
}

func dfsTree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := strconv.Itoa(root.Val)

	if root.Right != nil {
		return result + "(" + dfsTree2str(root.Left) + ")(" + dfsTree2str(root.Right) + ")"
	}

	if root.Left != nil {
		return result + "(" + dfsTree2str(root.Left) + ")"
	}

	return result + ""
}
