// https://leetcode.com/problems/triangle/
package medium

func MinimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}
