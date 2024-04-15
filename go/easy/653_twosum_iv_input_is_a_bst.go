// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
package easy

func FindTargetBST(root *TreeNode, k int) bool {
	queue := []TreeNode{*root}
	exist := map[int]bool{}
	for len(queue) > 0 {
		curr := queue[len(queue)-1]
		queue = queue[:len(queue)-1] // POP
		check := k - curr.Val
		if _, ok := exist[check]; ok {
			return true
		} else {
			exist[curr.Val] = true
		}
		if curr.Left != nil {
			queue = append(queue, *curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, *curr.Right)
		}
	}
	return false
}
