// https://leetcode.com/problems/copy-list-with-random-pointer/
package medium

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func CopyRandomList(head *Node) *Node {
	oldToCopy := map[*Node]*Node{}

	curr := head
	for curr != nil {
		clone := &Node{Val: curr.Val}
		oldToCopy[curr] = clone
		curr = curr.Next
	}

	for curr != nil {
		clone := oldToCopy[curr]
		clone.Next = oldToCopy[curr.Next]
		clone.Random = oldToCopy[curr.Random]
		curr = curr.Next
	}

	return oldToCopy[head]
}
