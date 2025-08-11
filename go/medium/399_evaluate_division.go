// https://leetcode.com/problems/evaluate-division/
package medium

func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	seen := make(map[string]bool)

	for i := 0; i < len(equations); i++ {
		a, b := equations[i][0], equations[i][1]

		if _, ok := graph[a]; !ok {
			graph[a] = make(map[string]float64)
		}
		graph[a][b] = values[i]

		if _, ok := graph[b]; !ok {
			graph[b] = make(map[string]float64)
		}
		graph[b][a] = 1 / values[i]
	}

	var dfs func(string, string) float64
	dfs = func(a, b string) float64 {
		if a == b {
			return 1.0
		}
		seen[a] = true
		for k, v := range graph[a] {
			if seen[k] {
				continue
			}
			result := dfs(k, b)
			if result > 0 {
				return result * v
			}
		}
		return -1.0
	}

	var isExists func(a, b string) bool
	isExists = func(a, b string) bool {
		if _, ok := graph[a]; !ok {
			return false
		}
		if _, ok := graph[b]; !ok {
			return false
		}
		return true
	}

	result := make([]float64, 0)
	for _, q := range queries {
		a, b := q[0], q[1]
		if !isExists(a, b) {
			result = append(result, -1.0)
			continue
		}

		seen = make(map[string]bool)
		result = append(result, dfs(a, b))
	}

	return result
}
