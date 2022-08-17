// https://leetcode.com/problems/unique-morse-code-words/
package easy

func UniqueMorseRepresentations(words []string) int {
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	mapM := map[string]int{}
	for _, w := range words {
		morseW := ""
		for i := 0; i < len(w); i++ {
			morseW += morse[w[i]-'a']
		}
		mapM[morseW] += 1
	}
	return len(mapM)
}
