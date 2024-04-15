// https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game
package easy

func Tictactoe(moves [][]int) string {
	n := 3
	rows := make([]int, n)
	cols := make([]int, n)
	diag := 0
	antiDiag := 0
	player := 1

	for _, v := range moves {
		row := v[0]
		col := v[1]

		rows[row] += player
		cols[col] += player

		if row == col {
			diag += player
		}

		if row+col == n-1 {
			antiDiag += player
		}

		if checkTheMove(rows[row], n) || checkTheMove(cols[col], n) ||
			checkTheMove(diag, n) || checkTheMove(antiDiag, n) {
			if player == 1 {
				return "A"
			}
			return "B"
		}

		player *= -1

	}

	if len(moves) == n*n {
		return "Draw"
	}
	return "Pending"
}

func checkTheMove(num, n int) bool {
	if num == n {
		return true
	} else if num == -n {
		return true
	}

	return false
}
