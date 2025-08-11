// https://leetcode.com/problems/reorganize-string/
package medium

func ReorganizeString(s string) string {
	chars := make([]int, 26)
	for i := range s {
		chars[s[i]-'a']++
	}

	maxCount, letter := 0, 0
	for i := 0; i < 26; i++ {
		if chars[i] > maxCount {
			maxCount = chars[i]
			letter = i
		}
	}

	if maxCount > (len(s)+1)/2 {
		return ""
	}

	result := make([]byte, len(s))
	idx := 0

	// Place the most frequent character
	for chars[letter] > 0 {
		result[idx] = byte(letter + 'a')
		idx += 2
		chars[letter]--
	}

	// Place the rest of the characters
	for i := 0; i < 26; i++ {
		for chars[i] > 0 {
			if idx >= len(s) {
				idx = 1
			}
			result[idx] = byte(i + 'a')
			idx += 2
			chars[i]--
		}
	}

	return string(result)
}

// Time complexity: O(N)
// Space complexity: O(k)
