// https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/
package easy

func CountCharacters(words []string, chars string) int {
	count := 0

	mapChars := make(map[byte]int)
	for i := 0; i < len(chars); i++ {
		mapChars[chars[i]]++
	}

	for _, word := range words {
		mapWords := make(map[byte]int)
		flag := true
		for i := 0; i < len(word); i++ {
			mapWords[word[i]]++
			if mapChars[word[i]] < mapWords[word[i]] {
				flag = false
				break
			}
		}
		if flag {
			count += len(word)
		}
	}

	return count
}
