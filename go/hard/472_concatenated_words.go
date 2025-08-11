// https://leetcode.com/problems/concatenated-words/
package hard

func findAllConcatenatedWordsInADict(words []string) []string {
	dict := make(map[string]bool)
	for _, w := range words {
		dict[w] = true
	}

	result := make([]string, 0)
	var visited []bool

	var dfs func(w string, idx int) bool
	dfs = func(w string, idx int) bool {
		if idx == len(w) {
			return true
		}
		if visited[idx] {
			return false
		}
		visited[idx] = true
		var x int
		if idx == 0 {
			x = 1
		}
		for i := len(w) - x; i > idx; i-- {
			if dict[w[idx:i]] && dfs(w, i) {
				return true
			}
		}
		return false
	}

	for _, w := range words {
		wLen := len(w)
		visited = make([]bool, wLen)
		if dfs(w, 0) {
			result = append(result, w)
		}
	}

	return result
}

// Time complexity: O(M^3⋅N)
// For each word, the constructed graph has M nodes and O(M^2) edges,
// and the DFS algorithm for reachability is O(M^2)
// without considering the time complexities of substring and HashSet. If we consider everything, the time complexity to check one word is O(M^3)
// and the total time complexity to check all words is O(M^3⋅N).

// Time complexity: O(M⋅N)
// This is the space to save all words in the dictionary, if we don't take M as a constant,
// there is also O(M) for the call stack to execute DFS, which wouldn't affect the space complexity anyways.
