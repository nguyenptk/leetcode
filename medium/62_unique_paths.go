// https://leetcode.com/problems/unique-paths/
package medium

var cache [][]int

func UniquePaths(m int, n int) int {
	cache = make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}
	return findRoute(m, n, 0, 0)
}

func findRoute(m, n, row, col int) int {
	if row >= m || col >= n {
		return 0
	}

	if row == m-1 && col == n-1 {
		return 1
	}

	if cache[row][col] > 0 {
		return cache[row][col]
	}

	down := findRoute(m, n, row+1, col)
	right := findRoute(m, n, row, col+1)

	cache[row][col] = down + right

	return cache[row][col]
}
