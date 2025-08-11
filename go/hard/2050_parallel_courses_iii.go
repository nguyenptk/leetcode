// https://leetcode.com/problems/parallel-courses-iii/
package hard

func MinimumTime(n int, relations [][]int, time []int) int {
	graph := make(map[int][]int)
	for _, r := range relations {
		from := r[0] - 1
		to := r[1] - 1
		graph[from] = append(graph[from], to)
	}

	result := 0
	memo := make(map[int]int)

	var dfs func(idx int) int
	dfs = func(idx int) int {
		if val, ok := memo[idx]; ok {
			return val
		}

		if _, ok := graph[idx]; !ok {
			return time[idx]
		}

		count := 0
		for _, nei := range graph[idx] {
			count = max(count, dfs(nei))
		}

		memo[idx] = time[idx] + count
		return memo[idx]
	}

	for i := 0; i < n; i++ {
		result = max(result, dfs(i))
	}

	return result
}
