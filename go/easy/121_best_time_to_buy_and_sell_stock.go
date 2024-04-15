// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
package easy

func MaxProfit(prices []int) int {
	l, r := 0, 1
	maxP := 0
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxP = max(maxP, profit)
		} else {
			l = r
		}
		r += 1
	}

	return maxP
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
