// https://leetcode.com/problems/zigzag-conversion/
package medium

import "strings"

func Convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	rows := make([]string, numRows)
	backward := true
	idx := 0

	for i := 0; i < len(s); i++ {
		rows[idx] += string(s[i])
		if idx == 0 || idx == numRows-1 {
			backward = !backward
		}
		if backward {
			idx--
		} else {
			idx++
		}
	}

	var result strings.Builder
	for _, row := range rows {
		result.WriteString(row)
	}
	return result.String()
}

// Time complexity: O(n)
// Space complexity: O(n)
