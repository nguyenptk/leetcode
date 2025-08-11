// https://leetcode.com/problems/step-by-step-directions-from-a-binary-tree-node-to-another/
package medium

import "strings"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getDirections(root *TreeNode, startValue int, destValue int) string {
	var dfs func(node *TreeNode, target int, path []byte) []byte
	dfs = func(node *TreeNode, target int, path []byte) []byte {
		if node == nil {
			return nil
		}
		if node.Val == target {
			return path
		}
		if left := dfs(node.Left, target, append(path, 'L')); left != nil {
			return left
		}
		if right := dfs(node.Right, target, append(path, 'R')); right != nil {
			return right
		}
		return nil
	}

	startArr := dfs(root, startValue, []byte{})
	destArr := dfs(root, destValue, []byte{})
	i := 0
	// chop off common prefix
	for len(startArr) > i && len(destArr) > i && startArr[i] == destArr[i] {
		i++
	}
	// up-moves for leftover startPath, then down-moves from destPath[i:]
	pathUp := strings.Repeat("U", len(startArr)-i)
	pathDown := string(destArr[i:])
	return pathUp + pathDown
}
