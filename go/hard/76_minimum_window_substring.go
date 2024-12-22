// https://leetcode.com/problems/minimum-window-substring/
package hard

func MinWindow(s string, t string) string {
	l, r := 0, 0
	target := make(map[uint8]int)
	curr := make(map[uint8]int)
	distinct := 0
	result := ""

	for i := range t {
		target[t[i]]++
	}

	for r < len(s) {
		curr[s[r]]++
		if target[s[r]] != 0 && target[s[r]] == curr[s[r]] {
			distinct++
		}

		for distinct == len(target) {
			if result == "" || r-l+1 < len(result) {
				result = s[l : r+1]
			}
			curr[s[l]]--
			if curr[s[l]] < target[s[l]] {
				distinct--
			}
			l++
		}
		r++
	}

	return result
}
