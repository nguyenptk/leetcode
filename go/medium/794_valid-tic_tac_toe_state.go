// https://leetcode.com/problems/valid-tic-tac-toe-state/
package medium

// 1. X always goes first.
// 2. Players alternate turns.
// 3. The number of Xs is equal to or one more than the number of Os.
// 4. There can be at most one winner.
// 5. If O has won, X count must equal O count.
// 6. If X has won, X count must be one more than O.

func ValidTicTacToe(board []string) bool {
	xCount, oCount := 0, 0

	for _, row := range board {
		for _, c := range row {
			if c == 'X' {
				xCount++
			} else if c == 'O' {
				oCount++
			}
		}
	}

	// Rule 1-3
	if oCount > xCount || xCount-oCount > 1 {
		return false
	}

	xWin := win(board, 'X')
	oWin := win(board, 'O')

	// Rule 4-6
	if xWin && oWin {
		return false
	}
	if xWin && xCount != oCount+1 {
		return false
	}
	if oWin && xCount != oCount {
		return false
	}

	return true
}

func win(board []string, player byte) bool {
	// Rows and Columns
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}

	// Diagonal
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}

	// Anti-diagonal
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}

	return false
}
