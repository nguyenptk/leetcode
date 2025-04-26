// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/
package medium

func NearestExit(maze [][]byte, entrance []int) int {
	rows := len(maze)
	cols := len(maze[0])
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	startR, startC := entrance[0], entrance[1]
	maze[startR][startC] = '+'

	queue := make([][]int, 0)
	queue = append(queue, []int{startR, startC, 0})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		currR, currC, currDis := node[0], node[1], node[2]

		for _, d := range dirs {
			nextR := currR + d[0]
			nextC := currC + d[1]

			if (nextR >= 0 && nextR < rows) && (nextC >= 0 && nextC < cols) && maze[nextR][nextC] == '.' {
				if nextR == 0 || nextR == rows-1 || nextC == 0 || nextC == cols-1 {
					return currDis + 1
				}
				maze[nextR][nextC] = '+'
				queue = append(queue, []int{nextR, nextC, currDis + 1})
			}
		}
	}

	return -1
}
