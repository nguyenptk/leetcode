// https://leetcode.com/problems/word-ladder/

package hard

func LadderLength(beginWord string, endWord string, wordList []string) int {
	if !contains(wordList, endWord) {
		return 0
	}

	n := len(beginWord)
	graph := make(map[string][]string) // h*t: [hot, hit]
	for _, w := range wordList {
		for i := 0; i < n; i++ {
			newWord := w[:i] + "*" + w[i+1:]
			graph[newWord] = append(graph[newWord], w)
		}
	}

	result := 1
	q := []string{beginWord}
	seen := map[string]bool{beginWord: true}

	for len(q) != 0 {
		for tmp := len(q); tmp > 0; tmp-- {
			word := q[0] // Pop front
			q = q[1:]
			if word == endWord {
				return result
			}
			for i := 0; i < len(word); i++ {
				pattern := word[:i] + "*" + word[i+1:]
				for _, neighbor := range graph[pattern] {
					if !seen[neighbor] {
						seen[neighbor] = true
						q = append(q, neighbor)
					}
				}
			}
		}
		result++
	}

	return 0
}

func contains(wordList []string, endWord string) bool {
	for _, w := range wordList {
		if w == endWord {
			return true
		}
	}
	return false
}
