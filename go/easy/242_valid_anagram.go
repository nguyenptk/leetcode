// https://leetcode.com/problems/valid-anagram/
package easy

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	table := make([]int, 26)
	for i := 0; i < len(s); i++ {
		table[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		table[t[i]-'a']--
		if table[t[i]-'a'] < 0 {
			return false
		}
	}

	return true
}
