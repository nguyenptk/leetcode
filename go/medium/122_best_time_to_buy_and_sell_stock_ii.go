// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
package medium

func MaxProfit(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			result += prices[i] - prices[i-1]
		}
	}

	return result
}
