// https://leetcode.com/problems/backspace-string-compare/
package easy

import (
	"fmt"
)

func BackspaceCompare(s string, t string) bool {
	afterS := removeBackspace(s)
	afterT := removeBackspace(t)

	if afterS != afterT {
		return false
	}
	return true
}

func removeBackspace(s string) string {
	result := []string{}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "#" {
			if i != 0 {
				result = append(result[:len(result)-1])
			}
		} else {
			result = append(result, string(s[i]))
		}
	}
	return fmt.Sprint(result)
}
