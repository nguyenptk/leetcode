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

	mapP := map[byte]string{}
	mapW := map[string]byte{}

	for i := 0; i < len(pattern); i++ {
		if w, ok := mapP[pattern[i]]; ok {
			if w != words[i] {
				return false
			}
		}
		if p, ok := mapW[words[i]]; ok {
			if p != pattern[i] {
				return false
			}
		}
		mapW[words[i]] = pattern[i]
		mapP[pattern[i]] = words[i]
	}

	return true
}
