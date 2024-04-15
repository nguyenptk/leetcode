// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/

package medium

type Node struct {
	Val      int
	Left     *Node
	Right    *Node
	Next     *Node
	Children []*Node
	Random   *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) != 0 {
		size := len(queue)
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		for i := 1; i < size; i++ {
			prev := queue[i-1]
			node := queue[i]

			prev.Next = node
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}

	return root
}
