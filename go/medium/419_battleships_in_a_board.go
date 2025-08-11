// https://leetcode.com/problems/battleships-in-a-board/
package medium

func CountBattleships(board [][]byte) int {
	result := 0

	rows := len(board)
	cols := len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] != 'X' {
				continue
			}
			if i > 0 && board[i-1][j] == 'X' {
				continue
			}
			if j > 0 && board[i][j-1] == 'X' {
				continue
			}
			result++
		}
	}

	return result
}
