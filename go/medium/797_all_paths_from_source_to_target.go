// https://leetcode.com/problems/all-paths-from-source-to-target
package medium

func AllPathsSourceTarget(graph [][]int) [][]int {
	result := [][]int{}
	findAllPaths(graph, 0, []int{0}, &result)
	return result
}

func findAllPaths(graph [][]int, u int, path []int, result *[][]int) {
	if u == len(graph)-1 {
		*result = append(*result, append(path))
		return
	}

	for _, v := range graph[u] {
		path = append(append([]int{}, path...), v)
		findAllPaths(graph, v, path, result)
		path = path[:len(path)-1] // POP
	}
}
