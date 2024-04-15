// https://leetcode.com/problems/find-players-with-zero-or-one-losses/
package medium

func FindWinners(matches [][]int) [][]int {
	losses := [100001]int{}
	for i := range losses {
		losses[i] = -1
	}

	for _, m := range matches {
		winner := m[0]
		loser := m[1]
		if losses[winner] == -1 {
			losses[winner] = 0
		}
		if losses[loser] == -1 {
			losses[loser] = 1
		} else {
			losses[loser]++
		}
	}

	result := make([][]int, 2)
	for i := 1; i < 100001; i++ {
		if losses[i] == 0 {
			result[0] = append(result[0], i)
		} else if losses[i] == 1 {
			result[1] = append(result[1], i)
		}
	}

	return result
}
