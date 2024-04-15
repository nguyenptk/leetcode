// https://leetcode.com/problems/binary-tree-inorder-traversal/
package easy

func InorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []TreeNode{}

	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, *curr)
			curr = curr.Left
		}
		curr = &stack[len(stack)-1]  // TOP
		stack = stack[:len(stack)-1] // POP
		result = append(result, curr.Val)
		curr = curr.Right
	}

	return result
}
