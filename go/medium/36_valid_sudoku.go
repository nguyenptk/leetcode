// https://leetcode.com/problems/valid-sudoku/
package medium

func IsValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			digit := board[r][c]
			// boxIndex represents the index of the box for the current row and column
			boxIndex := (r/3)*3 + c/3
			// check if digit already exists in the boxes, rows or columns
			if rows[r][digit] || cols[c][digit] || boxes[boxIndex][digit] {
				return false
			}
			rows[r][digit] = true
			cols[c][digit] = true
			boxes[boxIndex][digit] = true
		}
	}

	return true
}

// func isValidSudoku(board [][]byte) bool {
// 	rows := make([]map[byte]bool, 9)
// 	columns := make([]map[byte]bool, 9)
// 	boxes := make([]map[byte]bool, 9)

// 	for i := range rows {
// 		rows[i] = make(map[byte]bool)
// 		columns[i] = make(map[byte]bool)
// 		boxes[i] = make(map[byte]bool)
// 	}

// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			if board[i][j] != '.' {
// 				digit := board[i][j]

// 				boxIndex := (i/3)*3 + j/3

// 				rows[i][digit] = true
// 				columns[j][digit] = true
// 				boxes[boxIndex][digit] = true

// 				if rows[i][digit] && columns[j][digit] && boxes[boxIndex][digit] {
// 					return false
// 				}
// 			}
// 		}
// 	}

// 	return true
// }
