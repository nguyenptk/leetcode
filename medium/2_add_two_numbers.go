// https://leetcode.com/problems/add-two-numbers/
package medium

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{}
	current := head
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Next = &ListNode{
			Val: sum % 10,
		}
		current = current.Next
	}
	if carry != 0 {
		current.Next = &ListNode{
			Val: carry,
		}
	}
	return head.Next
}
