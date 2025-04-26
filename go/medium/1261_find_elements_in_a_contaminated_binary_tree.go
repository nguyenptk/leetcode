// https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/
package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
	root *TreeNode
	seen map[int]bool
}

func ConstructorFindElements(root *TreeNode) FindElements {
	return FindElements{
		root: root,
		seen: map[int]bool{},
	}
}

func (f *FindElements) Find(target int) bool {
	var dfs func(node *TreeNode, value int)
	dfs = func(node *TreeNode, value int) {
		if node == nil {
			return
		}
		if f.seen[value] {
			return
		}
		f.seen[value] = true
		dfs(node.Left, value*2+1)
		dfs(node.Right, value*2+2)
	}
	dfs(f.root, 0)

	return f.seen[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
