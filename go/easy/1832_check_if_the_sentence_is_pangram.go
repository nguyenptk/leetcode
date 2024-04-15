// https://leetcode.com/problems/check-if-the-sentence-is-pangram/
package easy

func CheckIfPangram(sentence string) bool {
	mapLetters := map[byte]bool{}

	for _, v := range sentence {
		if _, ok := mapLetters[byte(v)]; !ok {
			mapLetters[byte(v)] = true
		}
	}

	if len(mapLetters) == 26 {
		return true
	}

	return false
}
