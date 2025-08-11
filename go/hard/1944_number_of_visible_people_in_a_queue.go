// https://leetcode.com/problems/number-of-visible-people-in-a-queue/
package hard

func CanSeePersonsCount(heights []int) []int {
	result := make([]int, len(heights))
	stack := make([]int, 0)
	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] <= heights[i] {
			stack = stack[:len(stack)-1]
			result[i]++
		}
		if len(stack) > 0 {
			result[i]++
		}
		stack = append(stack, i)
	}
	return result
}
