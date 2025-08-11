// https://leetcode.com/problems/next-greater-node-in-linked-list/
package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func NextLargerNodes(head *ListNode) []int {
	values := make([]int, 0)
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	stack := make([]int, 0)
	result := make([]int, len(values))

	for i := 0; i < len(values); i++ {
		for len(stack) > 0 && values[stack[len(stack)-1]] < values[i] {
			smaller := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[smaller] = values[i]
		}
		stack = append(stack, i)
	}

	return result
}

// Time complexity: O(n)
// Space complexity: O(n)
