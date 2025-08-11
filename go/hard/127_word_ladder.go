// https://leetcode.com/problems/word-ladder/

package hard

func LadderLength(beginWord string, endWord string, wordList []string) int {
	n := len(beginWord)
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	// Build graph with intermediate patterns
	graph := make(map[string][]string)
	for _, w := range wordList {
		for i := 0; i < n; i++ {
			key := w[:i] + "*" + w[i+1:]
			graph[key] = append(graph[key], w)
		}
	}

	q := []string{beginWord}
	seen := map[string]bool{beginWord: true}
	steps := 1

	for len(q) > 0 {
		nextLevel := []string{}
		for _, word := range q {
			if word == endWord {
				return steps
			}
			for i := 0; i < n; i++ {
				pattern := word[:i] + "*" + word[i+1:]
				for _, nei := range graph[pattern] {
					if !seen[nei] {
						seen[nei] = true
						nextLevel = append(nextLevel, nei)
					}
				}
			}
		}
		q = nextLevel
		steps++
	}

	return 0
}

// n: length of each word
// L: number of words in wordList
// Time Complexity: O(L × n)
// Space Complexity: O(L × n)
