// https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/
package medium

func MinimumRounds(tasks []int) int {
	result := 0
	count := map[int]int{}

	for _, v := range tasks {
		count[v]++
	}

	// Freq = 1 -> impossible
	// Freq = 2 -> needs 1 round
	// Freq = 3 -> needs 1 round
	// Freq = 3k                           -> needs k rounds
	// Freq = 3k + 1 = 3 * (k - 1) + 2 * 2 -> needs k + 1 rounds
	// Freq = 3k + 2 = 3 * k       + 2 * 1 -> needs k + 1 rounds
	for _, freq := range count {
		if freq == 1 {
			return -1
		} else {
			result += (freq + 2) / 3
		}
	}

	return result
}
