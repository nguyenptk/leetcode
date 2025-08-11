// https://leetcode.com/problems/water-and-jug-problem/
package medium

func CanMeasureWater(x, y, target int) bool {
	cap := x + y
	if target > cap {
		return false
	}

	visited := make(map[int]bool)
	var dfs func(total int) bool
	dfs = func(total int) bool {
		if total == target {
			return true
		}
		if total < 0 || total > cap || visited[total] {
			return false
		}
		visited[total] = true
		return dfs(total+x) ||
			dfs(total-x) ||
			dfs(total+y) ||
			dfs(total-y)
	}

	return dfs(0)
}
