// https://leetcode.com/problems/subarray-sum-equals-k/

package medium

func SubarraySum(nums []int, k int) int {
	prefixSums := map[int]int{0: 1}

	sum := 0
	result := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		diff := sum - k
		if _, ok := prefixSums[diff]; ok {
			result += prefixSums[diff]
		}
		prefixSums[sum]++
	}

	return result
}
