// https://leetcode.com/problems/merge-k-sorted-lists/

package hard

import "github.com/emirpasic/gods/v2/queues/priorityqueue"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	minHeap := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(*ListNode).Val - b.(*ListNode).Val
	})

	for _, list := range lists {
		if list != nil {
			minHeap.Enqueue(list)
		}
	}

	result := &ListNode{Val: 0}
	curr := result

	for !minHeap.Empty() {
		node, _ := minHeap.Dequeue()
		curr.Next = node.(*ListNode)
		curr = curr.Next
		if curr.Next != nil {
			minHeap.Enqueue(curr.Next)
		}
	}

	return result.Next
}

// Time complexity: O(Nlogk) where k is the number of linked lists.
// Space complexity: O(n) Creating a new linked list costs O(n) space.
