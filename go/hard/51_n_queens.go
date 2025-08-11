// https://leetcode.com/problems/n-queens/
package hard

import "strings"

func SolveNQueens(n int) [][]string {
	cols := make(map[int]int)
	diags := make(map[int]int)
	antiDiags := make(map[int]int)
	result := make([][]string, 0)
	curr := make([]string, 0)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			result = append(result, append([]string{}, curr...))
		}
		for col := 0; col < n; col++ {
			if cols[col] > 0 || diags[row+col] > 0 || antiDiags[n+col-row-1] > 0 {
				continue
			}
			// Mark Q
			cols[col] = 1
			diags[row+col] = 1
			antiDiags[n+col-row-1] = 1
			var s strings.Builder
			for i := 0; i < n; i++ {
				if i == col {
					s.WriteString("Q")
				} else {
					s.WriteString(".")
				}
			}
			curr = append(curr, s.String())
			backtrack(row + 1)
			// Clear marked Q
			cols[col] = 0
			diags[row+col] = 0
			antiDiags[n+col-row-1] = 0
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(0)

	return result
}

// Time complexity: O(N!)
// Space complexity: O(N^2)
