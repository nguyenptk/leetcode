// https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/
package medium

func MaximumSubarraySum(nums []int, k int) int64 {
	result := 0
	l := 0
	sum := 0
	m := make(map[int]bool)

	for r := 0; r < len(nums); r++ {
		for m[nums[r]] {
			m[nums[l]] = false
			sum -= nums[l]
			l++
		}

		m[nums[r]] = true
		sum += nums[r]

		if r-l+1 == k {
			result = max(result, sum)
			m[nums[l]] = false
			sum -= nums[l]
			l++
		}
	}

	return int64(result)
}
