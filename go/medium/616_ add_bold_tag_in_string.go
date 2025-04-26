// https://leetcode.com/problems/add-bold-tag-in-string/
package medium

import "strings"

func AddBoldTag(s string, words []string) string {
	n := len(s)
	bold := make([]bool, n)

	for _, w := range words {
		wl := len(w)
		for i := 0; i+wl <= len(s); i++ {
			if s[i:i+wl] == w {
				for j := i; j < i+wl; j++ {
					bold[j] = true
				}
			}
		}
	}

	var sb strings.Builder

	for i := 0; i < n; i++ {
		if bold[i] && (i == 0 || !bold[i-1]) {
			sb.WriteString("<b>")
		}
		sb.WriteByte(s[i])
		if bold[i] && (i == n-1 || !bold[i+1]) {
			sb.WriteString("</b>")
		}
	}

	return sb.String()
}
