// https://leetcode.com/problems/count-subarrays-of-length-three-with-a-condition/
package easy

func CountSubarrays(nums []int) int {
	result := 0
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] == (nums[i-1]+nums[i+1])*2 {
			result++
		}
	}

	return result
}
