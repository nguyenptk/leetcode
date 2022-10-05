// https://leetcode.com/problems/add-one-row-to-tree/
package medium

func AddOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:   val,
			Left:  root,
			Right: nil,
		}
	}

	addRow(root, val, depth-1)

	return root
}

func addRow(parent *TreeNode, val, depth int) {
	if parent == nil {
		return
	}

	if depth == 1 {
		var temp *TreeNode

		temp = parent.Left
		parent.Left = &TreeNode{Val: val}
		parent.Left.Left = temp

		temp = parent.Right
		parent.Right = &TreeNode{Val: val}
		parent.Right.Right = temp

	} else {
		addRow(parent.Left, val, depth-1)
		addRow(parent.Right, val, depth-1)
	}
}
