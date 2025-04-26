// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/
package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func PairSum(head *ListNode) int {
	result := 0
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	curr := slow
	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	start := head
	for prev != nil {
		result = max(result, start.Val+prev.Val)
		prev = prev.Next
		start = start.Next
	}

	return result
}
