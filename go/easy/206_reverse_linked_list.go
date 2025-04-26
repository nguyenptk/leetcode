// https://leetcode.com/problems/reverse-linked-list
package easy

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}
	return prev
}

func ReverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := ReverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
