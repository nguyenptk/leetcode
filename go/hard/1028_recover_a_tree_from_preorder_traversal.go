// https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/
package hard

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func RecoverFromPreorder(traversal string) *TreeNode {
	stack := make([]*TreeNode, 0)
	idx := 0

	for idx < len(traversal) {
		depth := 0
		for idx < len(traversal) && traversal[idx] == '-' {
			depth++
			idx++
		}

		value := 0
		for idx < len(traversal) &&
			traversal[idx] >= '0' && traversal[idx] <= '9' {
			value = value*10 + int(traversal[idx]-'0')
			idx++
		}

		node := &TreeNode{Val: value}
		for len(stack) > depth {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			if stack[len(stack)-1].Left == nil {
				stack[len(stack)-1].Left = node
			} else {
				stack[len(stack)-1].Right = node
			}
		}

		stack = append(stack, node)
	}

	return stack[0]
}
