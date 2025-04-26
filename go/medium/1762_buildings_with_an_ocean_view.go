// https://leetcode.com/problems/buildings-with-an-ocean-view/
package medium

func FindBuildings(heights []int) []int {
	n := len(heights)
	stack := make([]int, 0)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[i] >= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return stack
}
