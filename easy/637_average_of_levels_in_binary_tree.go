// https://leetcode.com/problems/average-of-levels-in-binary-tree/

package easy

func AverageOfLevels(root *TreeNode) []float64 {
	result := []float64{}
	queue := []TreeNode{*root}
	for len(queue) > 0 {
		count := len(queue)
		sum := float64(0)
		for i := 0; i < count; i++ {
			curr := queue[0]  // FIRST
			queue = queue[1:] // POLL
			sum += float64(curr.Val)
			if curr.Left != nil {
				queue = append(queue, *curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, *curr.Right)
			}
		}
		result = append(result, sum/float64(count))
	}
	return result
}
