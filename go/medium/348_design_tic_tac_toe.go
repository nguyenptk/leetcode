// https://leetcode.com/problems/design-tic-tac-toe/
package medium

type TicTacToe struct {
	rows        []int
	cols        []int
	diagnol     int
	antiDiagnol int
}

func ConstructorTicTacToe(n int) TicTacToe {
	return TicTacToe{
		rows: make([]int, n),
		cols: make([]int, n),
	}
}

func (f *TicTacToe) Move(row int, col int, player int) int {
	mark := 1
	n := len(f.rows)
	win := n
	if player == 2 {
		mark = -1
		win *= -1
	}

	f.rows[row] += mark
	f.cols[col] += mark
	if row == col {
		f.diagnol += mark
	}

	if row+col == n-1 {
		f.antiDiagnol += mark
	}
	if f.rows[row] == win ||
		f.cols[col] == win ||
		f.diagnol == win ||
		f.antiDiagnol == win {
		return player
	}

	return 0
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
