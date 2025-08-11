// https://leetcode.com/problems/alien-dictionary/
package hard

func AlienOrder(words []string) string {
	graph := make(map[byte][]byte)
	visited := make(map[byte]int)
	allChars := make(map[byte]bool)

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			allChars[word[i]] = true
		}
	}

	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		minLen := min(len(w1), len(w2))
		if len(w1) > len(w2) && w1[:minLen] == w2[:minLen] {
			return ""
		}
		for j := 0; j < minLen; j++ {
			if w1[j] != w2[j] {
				graph[w1[j]] = append(graph[w1[j]], w2[j])
				break
			}
		}
	}

	var result []byte
	var dfs func(byte) bool
	dfs = func(ch byte) bool {
		if visited[ch] == 1 {
			return false
		}
		if visited[ch] == 2 {
			return true
		}
		visited[ch] = 1
		for _, nei := range graph[ch] {
			if !dfs(nei) {
				return false
			}
		}
		visited[ch] = 2
		result = append(result, ch)
		return true
	}

	for ch := range allChars {
		if visited[ch] == 0 {
			if !dfs(ch) {
				return ""
			}
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
