// https://leetcode.com/problems/find-minimum-time-to-reach-last-room-ii/
package medium

import (
	"math"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

type State struct {
	x, y, dist int
}

func MinTimeToReach(moveTime [][]int) int {
	n, m := len(moveTime), len(moveTime[0])
	data := make([][]int, n)
	visited := make([][]bool, n)
	for i := range data {
		data[i] = make([]int, m)
		visited[i] = make([]bool, m)
		for j := range data[i] {
			data[i][j] = math.MaxInt32
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	data[0][0] = 0
	q := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(*State).dist - b.(*State).dist
	})
	q.Enqueue(&State{0, 0, 0})

	for q.Size() > 0 {
		node, _ := q.Dequeue()
		curr := node.(*State)
		if visited[curr.x][curr.y] {
			continue
		}
		if curr.x == n-1 && curr.y == m-1 {
			break
		}
		visited[curr.x][curr.y] = true
		for _, dir := range dirs {
			nx, ny := curr.x+dir[0], curr.y+dir[1]
			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}
			dist := max(data[curr.x][curr.y], moveTime[nx][ny]) + (curr.x+curr.y)%2 + 1
			if data[nx][ny] > dist {
				data[nx][ny] = dist
				q.Enqueue(&State{nx, ny, dist})
			}
		}
	}

	return data[n-1][m-1]
}

// Time complexity: O(nmlog(nm)).
// Space complexity: O(nm).
