// https://leetcode.com/problems/reorder-list/
package medium

func ReorderList(head *ListNode) {
	if head == nil {
		return
	}

	// find a middle
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reserve the second part
	var prev *ListNode
	curr := slow
	for curr != nil {
		curr.Next = prev
		prev = curr
		curr = curr.Next
	}

	// merge the two sorted list
	first := head
	second := prev
	for second.Next != nil {
		tmp := first.Next
		first.Next = second
		first = tmp

		tmp = second.Next
		second.Next = second
		second = tmp
	}
}
