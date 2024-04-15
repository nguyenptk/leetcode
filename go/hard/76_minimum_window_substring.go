// https://leetcode.com/problems/minimum-window-substring/
package hard

func MinWindow(s string, t string) string {
	start, end := 0, 0
	target := make(map[uint8]int)
	current := make(map[uint8]int)
	distinct := 0
	minSubString := ""

	for i := range t {
		target[t[i]]++
	}

	for end < len(s) {
		current[s[end]]++
		if target[s[end]] != 0 && target[s[end]] == current[s[end]] {
			distinct++
		}

		for distinct == len(target) {
			if minSubString == "" {
				minSubString = s[start : end+1]
			}
			if end-start+1 < len(minSubString) {
				minSubString = s[start : end+1]
			}

			current[s[start]]--
			if current[s[start]] < target[s[start]] {
				distinct--
			}

			start++
		}
		end++
	}

	return minSubString
}
