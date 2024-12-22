// https://leetcode.com/problems/set-matrix-zeroes/
package medium

func setZeroes(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])
	rowZero := false
	colZero := false

	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			rowZero = true
			break
		}
	}

	for j := 0; j < cols; j++ {
		if matrix[0][j] == 0 {
			colZero = true
			break
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if rowZero {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
	if colZero {
		for j := 0; j < cols; j++ {
			matrix[0][j] = 0
		}
	}
}
