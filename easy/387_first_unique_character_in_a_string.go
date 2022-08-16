// https://leetcode.com/problems/first-unique-character-in-a-string/
package easy

func FirstUniqChar(s string) int {
	mapW := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := mapW[s[i]]; !ok {
			mapW[s[i]] = 1
		} else {
			mapW[s[i]] += 1
		}
	}

	for i := 0; i < len(s); i++ {
		if mapW[s[i]] == 1 {
			return i
		}
	}

	return -1
}
