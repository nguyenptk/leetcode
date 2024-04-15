// https://leetcode.com/problems/jump-game-ii/
package medium

func Jump(nums []int) int {
	result := 0
	n := len(nums) - 1
	currEnd := 0
	currFar := 0

	for i := 0; i < n; i++ {
		if currFar < i+nums[i] {
			currFar = i + nums[i]
		}
		if i == currEnd {
			result++
			currEnd = currFar
		}
	}

	return result
}
