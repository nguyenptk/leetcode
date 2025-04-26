// https://leetcode.com/problems/find-the-highest-altitude/
package easy

import "math"

func LargestAltitude(gain []int) int {
	n := len(gain)
	prefix := make([]int, n+1)
	prefix[0] = 0

	for i := 1; i <= n; i++ {
		prefix[i] = gain[i-1] + prefix[i-1]
	}

	result := math.MinInt
	for i := 0; i <= n; i++ {
		result = max(result, prefix[i])
	}

	return result
}
