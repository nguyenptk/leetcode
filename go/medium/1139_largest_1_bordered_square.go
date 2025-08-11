// https://leetcode.com/problems/largest-1-bordered-square/
package medium

func largest1BorderedSquare(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dpL := make([][]int, n)
	dpU := make([][]int, n)
	for i := range dpL {
		dpL[i] = make([]int, m)
		dpU[i] = make([]int, m)
	}

	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				if j > 0 {
					dpL[i][j] = dpL[i][j-1] + 1
				} else {
					dpL[i][j] = 1
				}
				if i > 0 {
					dpU[i][j] = dpU[i-1][j] + 1
				} else {
					dpU[i][j] = 1
				}
				// try all sizes from the largest possible down to best+1
				// side = length of the square's border
				maxSide := min(dpL[i][j], dpU[i][j])
				for side := maxSide; side > result; side-- {
					if dpL[i-side+1][j] >= side && dpU[i][j-side+1] >= side {
						result = side
						break
					}
				}
			}
		}
	}
	return result * result
}

// Time complexity: O(n·m·min(n,m))
// Space complexity: O(n·m)
