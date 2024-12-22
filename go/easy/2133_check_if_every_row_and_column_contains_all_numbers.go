// https://leetcode.com/problems/check-if-every-row-and-column-contains-all-numbers/

package easy

func CheckValid(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		row := make(map[int]int)
		col := make(map[int]int)
		for j := 0; j < len(matrix[0]); j++ {
			row[matrix[i][j]]++
			col[matrix[j][i]]++
		}
		if len(row) != len(matrix) || len(col) != len(matrix) {
			return false
		}
	}

	return true
}
