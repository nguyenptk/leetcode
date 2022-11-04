// https://leetcode.com/problems/reverse-vowels-of-a-string/
package easy

func ReverseVowels(s string) string {
	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	queue := []string{}
	for i := 0; i < len(s); i++ {
		if contains(vowels, string(s[i])) {
			queue = append(queue, string(s[i]))
		}
	}
	result := ""
	for i := 0; i < len(s); i++ {
		if contains(vowels, string(s[i])) {
			end := queue[len(queue)-1]   // TOP
			queue = queue[:len(queue)-1] // POP
			result += end
		} else {
			result += string(s[i])
		}
	}
	return string(result)
}

func contains(vowels []string, s string) bool {
	for _, v := range vowels {
		if v == s {
			return true
		}
	}
	return false
}
