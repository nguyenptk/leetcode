// https://leetcode.com/problems/matchsticks-to-square/
package medium

import "sort"

func Makesquare(matchsticks []int) bool {
	total := 0
	for _, v := range matchsticks {
		total += v
	}
	if total%4 != 0 {
		return false
	}
	square := make([]int, 4)
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})
	return solve(matchsticks, total/4, 0, square)
}

func solve(matchsticks []int, target, pos int, square []int) bool {
	if pos >= len(matchsticks) {
		return square[0] == target && square[1] == target && square[2] == target
	}
	for i := 0; i < 4; i++ {
		if square[i]+matchsticks[pos] > target {
			continue
		}
		square[i] += matchsticks[pos]
		if solve(matchsticks, target, pos+1, square) {
			return true
		}
		square[i] -= matchsticks[pos]
	}
	return false
}
