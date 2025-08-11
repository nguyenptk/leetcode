// https://leetcode.com/problems/reverse-nodes-in-k-group/
package hard

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func ReverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	groupPrev := dummy

	for {
		kth := getKthNode(groupPrev, k)
		if kth == nil {
			break
		}
		groupNext := kth.Next

		// Reverse
		prev, curr := kth.Next, groupPrev.Next
		for curr != groupNext {
			tmp := curr.Next
			curr.Next = prev
			prev = curr
			curr = tmp
		}

		tmp := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = tmp
	}

	return dummy.Next
}

func getKthNode(curr *ListNode, k int) *ListNode {
	for curr != nil && k > 0 {
		curr = curr.Next
		k--
	}
	return curr
}

// Time Complexity: O(N)
// Space Complexity: O(N/k) used up by the recursion stack

// func getKthNode(curr *ListNode, k int) *ListNode {
//     for curr != nil && k > 0 {
//         curr = curr.Next
//         k--
//     }
//     return curr
// }

// func reverseLinkedList(head *ListNode, k int) *ListNode {
//     var newHead *ListNode = nil
//     ptr := head
//     for k > 0 {
//         nextNode := ptr.Next
//         ptr.Next = newHead
//         newHead = ptr
//         ptr = nextNode
//         k--
//     }
//     return newHead
// }

// func reverseKGroup(head *ListNode, k int) *ListNode {
//     count := 0
//     ptr := head
//     for count < k && ptr != nil {
//         ptr = ptr.Next
//         count++
//     }
//     if count == k {
//         reversedHead := reverseLinkedList(head, k)
//         head.Next = reverseKGroup(ptr, k)
//         return reversedHead
//     }
//     return head
// }
