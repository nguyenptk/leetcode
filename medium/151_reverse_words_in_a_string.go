// https://leetcode.com/problems/reverse-words-in-a-string/
package medium

func ReverseWords(s string) string {
	words := []string{}
	w := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if w != "" {
				words = append(words, w)
				w = ""
			}
		} else {
			w += string(s[i])
		}
	}
	// insert the last w when i is over len(s)
	if w != "" {
		words = append(words, w)
	}

	result := words[len(words)-1]
	for i := len(words) - 2; i >= 0; i-- {
		result += " " + words[i]
	}

	return string(result)
}
