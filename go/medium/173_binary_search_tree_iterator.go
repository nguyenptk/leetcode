// https://leetcode.com/problems/binary-search-tree-iterator/
package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack []*TreeNode
}

func ConstructorBSTIterator(root *TreeNode) BSTIterator {
	f := BSTIterator{stack: make([]*TreeNode, 0)}
	f.pushLeft(root)
	return f
}

func (f *BSTIterator) pushLeft(node *TreeNode) {
	for node != nil {
		f.stack = append(f.stack, node)
		node = node.Left
	}
}

func (f *BSTIterator) Next() int {
	node := f.stack[len(f.stack)-1]
	f.stack = f.stack[:len(f.stack)-1]
	if node.Right != nil {
		f.pushLeft(node.Right)
	}
	return node.Val
}

func (f *BSTIterator) HasNext() bool {
	return len(f.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
