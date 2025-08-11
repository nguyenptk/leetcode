// https://leetcode.com/problems/longest-duplicate-substring/
package hard

func LongestDupSubstring(s string) string {
	l, r := 0, len(s)
	result := ""
	for l < r {
		m := (l + r) / 2
		d := findDuplicated(s, m)
		if d == "" {
			if r == m {
				break
			}
			r = m
		} else {
			if l == m {
				break
			}
			l = m
			result = d
		}
	}

	return result
}

func findDuplicated(s string, length int) string {
	dict := make(map[string]bool)
	for i := 0; i <= len(s)-length; i++ {
		str := s[i : i+length]
		if dict[str] {
			return str
		}
		dict[str] = true
	}
	return ""
}
