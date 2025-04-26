// https://leetcode.com/problems/minimum-cost-for-tickets/
package medium

func MinCostTickets(days []int, costs []int) int {
	daySet := make(map[int]bool)
	for _, day := range days {
		daySet[day] = true
	}

	memo := make(map[int]int, 0)

	var dfs func(currDay int) int
	dfs = func(currDay int) int {
		if currDay > days[len(days)-1] {
			return 0
		}
		if val, ok := memo[currDay]; ok {
			return val
		}
		if !daySet[currDay] {
			return dfs(currDay + 1)
		}
		oneDay := costs[0] + dfs(currDay+1)
		sevenDay := costs[1] + dfs(currDay+7)
		thirtyDay := costs[2] + dfs(currDay+30)

		memo[currDay] = min(oneDay, min(sevenDay, thirtyDay))
		return memo[currDay]
	}

	return dfs(1)
}
