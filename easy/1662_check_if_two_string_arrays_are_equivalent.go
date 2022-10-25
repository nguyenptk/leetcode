// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/
package easy

func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
	w1 := ""
	for _, v := range word1 {
		w1 += v
	}
	w2 := ""
	for _, v := range word2 {
		w2 += v
	}
	if w1 == w2 {
		return true
	}
	return false
}
