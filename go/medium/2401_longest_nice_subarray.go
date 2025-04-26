// https://leetcode.com/problems/longest-nice-subarray/
package medium

func LongestNiceSubarray(nums []int) int {
	result := 0
	usedBits := 0
	l := 0
	for r := 0; r < len(nums); r++ {
		for usedBits&nums[r] != 0 {
			usedBits ^= nums[l]
			l++
		}
		usedBits |= nums[r]
		result = max(result, r-l+1)
	}

	return result
}
