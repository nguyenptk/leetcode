// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
package easy

import "math"

func MaxProfit(prices []int) int {
	result := 0
	minPrice := math.MaxInt
	for i := range prices {
		minPrice = min(minPrice, prices[i])
		result = max(result, prices[i]-minPrice)
	}

	return result
}
