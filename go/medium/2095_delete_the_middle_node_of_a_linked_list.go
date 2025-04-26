// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DeleteMiddle(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	slow := dummy
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
