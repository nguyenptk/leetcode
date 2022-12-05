// https://leetcode.com/problems/middle-of-the-linked-list/
package easy

func MiddleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
