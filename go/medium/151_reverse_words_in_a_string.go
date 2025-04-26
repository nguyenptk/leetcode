// https://leetcode.com/problems/reverse-words-in-a-string/
package medium

import "strings"

func ReverseWords(s string) string {
	s = strings.TrimSpace(s)
	arr := strings.Fields(s)
	l := 0
	r := len(arr) - 1

	for l <= r {
		if arr[l] != "" && arr[r] != "" {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	return strings.Join(arr, " ")
}
