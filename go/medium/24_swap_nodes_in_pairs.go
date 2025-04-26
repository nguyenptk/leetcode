// https://leetcode.com/problems/swap-nodes-in-pairs/
package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	tail, first := dummy, head
	for first != nil && first.Next != nil {
		second := first.Next

		tail.Next = second
		first.Next = second.Next
		second.Next = first

		tail = first
		first = first.Next
	}

	return dummy.Next
}
