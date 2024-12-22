// https://leetcode.com/problems/n-queens
package medium

func SolveNQueens(n int) [][]string {
	col := make(map[int]int)
	diag := make(map[int]int)
	antiDiag := make(map[int]int)
	result := make([][]string, 0)
	curr := make([]string, 0)

	var backtrack func(y int)
	backtrack = func(y int) {
		if y == n {
			result = append(result, append([]string{}, curr...))
		}
		for x := 0; x < n; x++ {
			if col[x] > 0 || diag[x+y] > 0 || antiDiag[n+x-y-1] > 0 {
				continue
			}
			// mark Q
			col[x] = 1
			diag[x+y] = 1
			antiDiag[n+x-y-1] = 1
			s := ""
			for i := 0; i < n; i++ {
				if i == x {
					s += "Q"
				} else {
					s += "."
				}
			}
			curr = append(curr, s)
			backtrack(y + 1)
			// clear mark Q
			col[x] = 0
			diag[x+y] = 0
			antiDiag[n+x-y-1] = 0
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(0)

	return result
}
