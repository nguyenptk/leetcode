// https://leetcode.com/problems/word-search
package medium

func Exist(board [][]byte, word string) bool {
	rows := len(board)
	if rows == 0 {
		return false
	}
	cols := len(board[0])

	var backtrack func(row, col, index int) bool
	backtrack = func(row, col, index int) bool {
		if row < 0 || col < 0 || row == rows || col == cols {
			return false
		}
		if board[row][col] != word[index] {
			return false
		}
		if board[row][col] == '*' {
			return false
		}
		if index == len(word)-1 {
			return true
		}

		tmp := board[row][col]
		board[row][col] = '*'
		result := backtrack(row+1, col, index+1) ||
			backtrack(row-1, col, index+1) ||
			backtrack(row, col+1, index+1) ||
			backtrack(row, col-1, index+1)
		board[row][col] = tmp

		return result
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}
