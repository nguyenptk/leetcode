// https://leetcode.com/problems/game-of-life/
package medium

func GameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, 1}, {1, 0}, {1, -1},
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			live := 0
			for _, d := range dirs {
				x, y := i+d[0], j+d[1]
				if x >= 0 && x < m && y >= 0 && y < n {
					if board[x][y] == 1 || board[x][y] == 2 {
						live++
					}
				}
			}
			if board[i][j] == 1 {
				if live < 2 || live > 3 {
					board[i][j] = 2 // was live, now dead
				}
			} else {
				if live == 3 {
					board[i][j] = 3 // was dead, now live
				}
			}
		}
	}

	// Finalize states
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}
}

// Time complexity: O(m × n)
// Space complexity: O(1) (in-place with encoded states)
