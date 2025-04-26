// https://leetcode.com/problems/find-the-punishment-number-of-an-integer/
package medium

import "strconv"

func PunishmentNumber(n int) int {
	result := 0
	var backtrack func(idx, num int, target string) bool
	backtrack = func(idx, num int, target string) bool {
		if idx == len(target) {
			return num == 0
		}

		for i := idx; i < len(target); i++ {
			val, _ := strconv.Atoi(target[idx : i+1])
			if num >= val && backtrack(i+1, num-val, target) {
				return true
			}
		}

		return false
	}

	for i := 1; i <= n; i++ {
		target := strconv.Itoa(i * i)
		if backtrack(0, i, target) {
			result += i * i
		}
	}

	return result
}
