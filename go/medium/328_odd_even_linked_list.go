// https://leetcode.com/problems/odd-even-linked-list/
package medium

func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd, even := head, head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next, even.Next = even.Next, even.Next.Next
		odd, even = odd.Next, even.Next
	}

	odd.Next = evenHead
	return head
}
