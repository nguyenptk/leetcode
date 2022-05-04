// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
package easy

func MaxProfit(prices []int) int {
	minPrice := 10000
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
