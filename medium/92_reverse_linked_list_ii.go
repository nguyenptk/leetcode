// https://leetcode.com/problems/reverse-linked-list-ii/
package medium

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	result := &ListNode{
		Val:  0,
		Next: head,
	}

	// prev of the sublist [left, right]
	prev := result

	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	// tail of the sublist [left, right]
	tail := prev.Next

	// reserse the sublist [left, right] one by one
	for i := 0; i < right-left; i++ {
		cache := tail.Next
		tail.Next = cache.Next
		cache.Next = prev.Next
		prev.Next = cache
	}

	return result.Next
}
