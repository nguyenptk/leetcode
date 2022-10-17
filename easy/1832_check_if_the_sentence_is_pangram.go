// https://leetcode.com/problems/check-if-the-sentence-is-pangram/
package easy

func CheckIfPangram(sentence string) bool {
	mapLetters := map[byte]bool{}

	for _, v := range sentence {
		key := byte(v)
		if _, ok := mapLetters[key]; !ok {
			mapLetters[key] = true
		}
	}

	if len(mapLetters) == 26 {
		return true
	}

	return false
}
