// https://leetcode.com/problems/restore-ip-addresses/
package medium

import (
	"strconv"
	"strings"
)

func RestoreIpAddresses(s string) []string {
	var result []string
	var parts []string

	isValid := func(seg string) bool {
		if len(seg) > 1 && seg[0] == '0' {
			return false
		}
		// parse to int and check 0–255
		val, err := strconv.Atoi(seg)
		return err == nil && val >= 0 && val <= 255
	}

	var dfs func(start int)
	dfs = func(start int) {
		if len(parts) == 4 {
			if start == len(s) {
				result = append(result, strings.Join(parts, "."))
			}
			return
		}

		for l := 1; l <= 3 && start+l <= len(s); l++ {
			seg := s[start : start+l]
			if !isValid(seg) {
				continue
			}
			parts = append(parts, seg)
			dfs(start + l)
			parts = parts[:len(parts)-1]
		}
	}

	dfs(0)
	return result
}
