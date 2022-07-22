// https://leetcode.com/problems/partition-list/
package medium

func Partition(head *ListNode, x int) *ListNode {
	frontDummy := &ListNode{
		Val: 0,
	}
	front := frontDummy

	backDummy := &ListNode{
		Val: 0,
	}
	back := backDummy

	curr := head

	for curr != nil {
		if curr.Val < x {
			front.Next = curr
			front = curr
		} else {
			back.Next = curr
			back = curr
		}
		curr = curr.Next
	}
	front.Next = backDummy.Next
	back.Next = nil

	return frontDummy.Next
}
