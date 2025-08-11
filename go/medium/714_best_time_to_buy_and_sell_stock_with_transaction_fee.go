// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
package medium

func MaxProfitWithTransactionFee(prices []int, fee int) int {
	cash := 0
	hold := -prices[0]

	for i := 1; i < len(prices); i++ {
		cash = max(cash, hold+prices[i]-fee)
		hold = max(hold, cash-prices[i])
	}

	return cash
}

// Time: O(n)
// Space: O(1) (only two variables: cash and hold)
