// https://leetcode.com/problems/odd-even-linked-list/
package medium

func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	curr := head
	next := head.Next

	for next != nil && next.Next != nil {
		tmp := next.Next
		tmp2 := curr.Next

		curr.Next = tmp
		next.Next = tmp.Next
		tmp.Next = tmp2

		curr = curr.Next
		next = next.Next
	}

	return head
}
