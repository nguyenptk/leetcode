// https://leetcode.com/problems/surrounded-regions
package medium

func Solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	// 1. Capture unsurrounded regions (O -> *) by cols
	for c := 0; c < cols; c++ {
		dfsSolve(board, rows, cols, 0, c)
		dfsSolve(board, rows, cols, rows-1, c)
	}

	// 2. Capture unsurrounded regions (O -> *) by rows
	for r := 0; r < rows; r++ {
		dfsSolve(board, rows, cols, r, 0)
		dfsSolve(board, rows, cols, r, cols-1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// 3. Capture surrounded regions (O -> X)
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			}
			// 4. Uncapture unsurrounded regions (* -> O)
			if board[r][c] == '*' {
				board[r][c] = 'O'
			}
		}
	}
}

func dfsSolve(board [][]byte, rows, cols, r, c int) {
	if r < 0 || r >= rows ||
		c < 0 || c >= cols ||
		board[r][c] == 'X' || board[r][c] == '*' {
		return
	}
	board[r][c] = '*'
	dfsSolve(board, rows, cols, r+1, c)
	dfsSolve(board, rows, cols, r-1, c)
	dfsSolve(board, rows, cols, r, c+1)
	dfsSolve(board, rows, cols, r, c-1)
}
