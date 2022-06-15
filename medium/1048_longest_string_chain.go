// https://leetcode.com/problems/longest-string-chain/
package medium

import "sort"

func LongestStrChain(words []string) int {
	// Init DP cache
	cache := map[string]int{}

	// Sort the words array
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	// Init the result
	best := 0

	for _, word := range words {
		longest := 0
		for i := 0; i < len(word); i++ {
			subWord := word[:i] + word[i+1:]        // get every subword of i->i+1 of word
			if _, exist := cache[subWord]; !exist { // check subWord exist?
				longest = max(longest, 1)
			} else {
				longest = max(longest, cache[subWord]+1)
			}
		}
		cache[word] = longest     // add word to cache
		best = max(best, longest) // compare best & longest
	}

	return best
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
