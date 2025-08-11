// https://leetcode.com/problems/snakes-and-ladders/
package medium

type PairSaL struct {
	row int
	col int
}

func SnakesAndLadders(board [][]int) int {
	n := len(board)

	// 1. Flatten board into a 1D array of positions
	cells := make([]*PairSaL, n*n+1)
	label := 1
	columns := make([]int, n)
	for i := range columns {
		columns[i] = i
	}
	for row := n - 1; row >= 0; row-- {
		for _, col := range columns {
			cells[label] = &PairSaL{row, col}
			label++
		}
		// Reverse columns to simulate zigzag
		reverse(columns)
	}

	// 2. BFS
	dist := make([]int, n*n+1)
	for i := range dist {
		dist[i] = -1
	}
	dist[1] = 0
	queue := []int{1}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for move := 1; move <= 6 && curr+move <= n*n; move++ {
			next := curr + move
			row := cells[next].row
			col := cells[next].col
			dest := next
			if board[row][col] != -1 {
				dest = board[row][col]
			}
			if dist[dest] == -1 {
				dist[dest] = dist[curr] + 1
				queue = append(queue, dest)
			}
		}
	}

	return dist[n*n]
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
