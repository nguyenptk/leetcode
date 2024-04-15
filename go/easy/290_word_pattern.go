// https://leetcode.com/problems/word-pattern/
package easy

import (
	"strings"
)

func WordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	mapP := map[byte]int{}
	mapW := map[string]int{}

	for i := 0; i < len(pattern); i++ {
		if mapP[pattern[i]] != mapW[words[i]] {
			return false
		} else {
			mapP[pattern[i]] = i + 1
			mapW[words[i]] = i + 1
		}
	}

	return true
}
