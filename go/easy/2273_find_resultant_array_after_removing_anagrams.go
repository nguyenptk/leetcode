// https://leetcode.com/problems/find-resultant-array-after-removing-anagrams/

package easy

func RemoveAnagrams(words []string) []string {
	for i := 1; i < len(words); i++ {
		if isAnagrams(words[i-1], words[i]) {
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return words
}

func isAnagrams(s string, t string) bool {
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
