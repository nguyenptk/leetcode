// https://leetcode.com/problems/decode-ways/
package medium

import "strconv"

func NumDecodings(s string) int {
	return recNumDecodings(s, 0, make(map[int]int))
}

func recNumDecodings(s string, idx int, memo map[int]int) int {
	if idx == len(s) {
		return 1
	}
	if s[idx] == '0' {
		return 0
	}
	if idx == len(s)-1 {
		return 1
	}

	if _, ok := memo[idx]; ok {
		return memo[idx]
	}

	count := recNumDecodings(s, idx+1, memo)
	c2, _ := strconv.Atoi(s[idx : idx+2])
	if c2 <= 26 {
		count += recNumDecodings(s, idx+2, memo)
	}
	memo[idx] = count

	return count
}
