// https://leetcode.com/problems/n-ary-tree-level-order-traversal/
package medium

func NAryLevelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []Node{*root}

	for len(queue) > 0 {
		currLevel := []int{}
		for i := len(queue); i > 0; i-- {
			node := queue[0]  // First
			queue = queue[1:] // Poll
			currLevel = append(currLevel, node.Val)
			for _, v := range node.Children {
				queue = append(queue, *v)
			}
		}
		result = append(result, currLevel)
	}

	return result
}
