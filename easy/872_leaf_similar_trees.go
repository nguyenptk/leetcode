// https: //leetcode.com/problems/leaf-similar-trees/
package easy

import "reflect"

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var l1 []int
	var l2 []int
	dfs(root1, &l1)
	dfs(root2, &l2)

	return reflect.DeepEqual(l1, l2)
}

func dfs(node *TreeNode, l *[]int) {
	if node != nil {
		if node.Left == nil && node.Right == nil {
			*l = append(*l, node.Val)
		}

		dfs(node.Left, l)
		dfs(node.Right, l)
	}
}
