// https://leetcode.com/problems/add-two-numbers/
package medium

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil || l2 != nil || carry != 0 {
		l1Val := 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		l2Val := 0
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		colSum := l1Val + l2Val + carry
		carry = colSum / 10
		new := &ListNode{Val: colSum % 10}
		curr.Next = new
		curr = curr.Next
	}

	return dummy.Next
}
