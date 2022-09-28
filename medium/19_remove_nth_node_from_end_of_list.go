// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
package medium

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
