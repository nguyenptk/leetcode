// https://leetcode.com/problems/maximum-ascending-subarray-sum/
package easy

func MaxAscendingSum(nums []int) int {
	result := 0
	count := nums[0]
	for r := 1; r < len(nums); r++ {
		if nums[r] > nums[r-1] {
			count += nums[r]
		} else {
			result = max(result, count)
			count = nums[r]
		}
	}

	return max(result, count)
}
