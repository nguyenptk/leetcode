// https://leetcode.com/problems/reverse-words-in-a-string-iii/
package medium

import (
	"strings"
)

func ReverseWords(s string) string {
	result := ""
	arrS := strings.Split(s, " ")
	for k, v := range arrS {
		arrB := []byte(v)
		for i := len(arrB) - 1; i >= 0; i-- {
			result += string(arrB[i])
		}
		if k < len(arrS)-1 {
			result += " "
		}
	}

	return result
}
