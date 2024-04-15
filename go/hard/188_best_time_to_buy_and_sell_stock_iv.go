// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
package hard

import "math"

func MaxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	// Init result
	result := make([][]int, k+1)
	for i := range result {
		result[i] = make([]int, n+1)
	}

	// Fill the table in bottom-up
	for i := 1; i <= k; i++ {
		prevDiff := math.MinInt
		for j := 1; j < n; j++ {
			if prevDiff < (result[i-1][j-1] - prices[j-1]) {
				prevDiff = result[i-1][j-1] - prices[j-1]
			}
			if result[i][j-1] < prices[j]+prevDiff {
				result[i][j] = prices[j] + prevDiff
			} else {
				result[i][j] = result[i][j-1]
			}
		}
	}

	return result[k][n-1]
}
