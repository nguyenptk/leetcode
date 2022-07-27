// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
package medium

func Flatten(root *TreeNode) {
	// return if root is nil or if it is a leaf node
	if root == nil || root.Left == nil && root.Right == nil {
		return
	}
	// if root->left exists then we have to make it
	// root->right
	if root.Left != nil {
		// move left recursively
		Flatten(root.Left)
		// store the node root->right
		tmpR := root.Right
		root.Right = root.Left
		root.Left = nil
		// find the position to insert the stored value
		r := root.Right
		for r.Right != nil {
			r = r.Right
		}
		// insert the stored value
		r.Right = tmpR
	}
	// move right recursively
	Flatten(root.Right)
}
