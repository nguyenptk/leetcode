// https://leetcode.com/problems/make-costs-of-paths-equal-in-a-binary-tree/
package medium

func MinIncrements(n int, cost []int) int {
	result := 0

	var dfs func(idx int) int
	dfs = func(idx int) int {
		if idx*2 > n {
			return cost[idx-1]
		}
		l := dfs(idx * 2)
		r := dfs(idx*2 + 1)
		result += abs(l - r)
		return max(l, r) + cost[idx-1]
	}

	dfs(1)
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
