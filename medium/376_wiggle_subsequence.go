// https://leetcode.com/problems/wiggle-subsequence/
package medium

func WiggleMaxLength(nums []int) int {
	n := len(nums)
	i := 1

	// Check duplicated numbers
	for i < n && nums[i] == nums[i-1] {
		i++
	}
	if i == n {
		return 1
	}

	// Initialize up-flag & count
	up := nums[i-1] > nums[i]
	count := 1
	for ; i < n; i++ {
		if (up && nums[i] < nums[i-1]) || (!up && nums[i] > nums[i-1]) {
			up = !up
			count++
		}
	}
	return count
}
