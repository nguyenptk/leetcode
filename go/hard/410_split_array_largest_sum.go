// https://leetcode.com/problems/split-array-largest-sum/
package hard

import "math"

func SplitArray(nums []int, k int) int {
	n := len(nums)

	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	memo := make(map[[2]int]int)

	var minimum func(idx int, count int) int
	minimum = func(idx int, count int) int {
		if count == 1 {
			return prefixSum[n] - prefixSum[idx]
		}

		key := [2]int{idx, count}
		if val, ok := memo[key]; ok {
			return val
		}

		minSum := math.MaxInt
		for i := idx; i <= n-count; i++ {
			first := prefixSum[i+1] - prefixSum[idx]
			maxSum := max(first, minimum(i+1, count-1))
			minSum = min(minSum, maxSum)
			if first >= minSum {
				break
			}
		}
		memo[key] = minSum

		return memo[key]
	}

	return minimum(0, k)
}

// Time complexity: O(N^2⋅M)
// Space complexity: O(N⋅M)
