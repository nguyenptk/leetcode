// https://leetcode.com/problems/maximum-difference-between-adjacent-elements-in-a-circular-array/
package easy

func MaxAdjacentDistance(nums []int) int {
	result := abs(nums[0] - nums[len(nums)-1])
	for i := 0; i < len(nums)-1; i++ {
		result = max(result, abs(nums[i]-nums[i+1]))
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
