// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
package medium

var mp map[int]int
var preIndex int

func BuildTree(preorder []int, inorder []int) *TreeNode {
	// Initialize map for storage and preIndex
	mp = map[int]int{}
	preIndex = 0

	n := len(inorder)
	for i := 0; i < n; i++ {
		mp[inorder[i]] = i
	}
	return construct(preorder, inorder, 0, n-1, mp)
}

func construct(preorder, inorder []int, start, end int, mp map[int]int) *TreeNode {
	if start > end {
		return nil
	}

	// Pick the current node from Preorder traversal using preIndex and increment preIndex
	curr := preorder[preIndex]
	preIndex++
	node := &TreeNode{
		Val: curr,
	}

	// If this node has no childrent then return
	if start == end {
		return node
	}

	// Find the index of this node in Inorder traversal
	index := mp[curr]

	// Using index in Inorder traversal, construct left and right subtress
	node.Left = construct(preorder, inorder, start, index-1, mp)
	node.Right = construct(preorder, inorder, index+1, end, mp)

	return node
}
