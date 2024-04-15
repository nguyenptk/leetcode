// https://leetcode.com/problems/find-and-replace-pattern/
package medium

var ans []string
var codex map[byte]byte
var cipher []byte

func FindAndReplacePattern(words []string, pattern string) []string {
	ans = []string{}
	codex = map[byte]byte{}
	cipher = []byte(pattern)
	for i := 0; i < len(pattern); i++ {
		cipher[i] = translate(cipher[i])
	}
	for _, word := range words {
		compare(word)
	}
	return ans
}

func translate(c byte) byte {
	if _, ok := codex[c]; !ok {
		codex[c] = byte(97 + len(codex))
	}
	return codex[c]
}

func compare(word string) {
	codex = map[byte]byte{}
	for i := 0; i < len(word); i++ {
		if translate(word[i]) != cipher[i] {
			return
		}
	}
	ans = append(ans, word)
}
