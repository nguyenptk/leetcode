// https://leetcode.com/problems/put-marbles-in-bags/
package hard

import "sort"

func PutMarbles(weights []int, k int) int64 {
	n := len(weights)
	pairs := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		pairs[i] = weights[i] + weights[i+1]
	}

	sort.Ints(pairs)

	result := int64(0)
	for i := 0; i < k-1; i++ {
		result += int64(pairs[n-2-i] - pairs[i])
	}

	return result
}
