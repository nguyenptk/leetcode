// https://leetcode.com/problems/palindrome-linked-list/
package easy

func IsPalindromeLinkedList(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	l := head
	r := prev
	for r != nil {
		if l.Val != r.Val {
			return false
		}
		l = l.Next
		r = r.Next
	}

	return true
}
