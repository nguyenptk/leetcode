// https://leetcode.com/problems/design-add-and-search-words-data-structure/
package medium

type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool
}

type WordDictionary struct {
	root *TrieNode
}

func ConstructorWordDictionary() WordDictionary {
	return WordDictionary{
		root: &TrieNode{
			children: make(map[byte]*TrieNode),
		},
	}
}

func (t *WordDictionary) AddWord(word string) {
	curr := t.root
	for c := 0; c < len(word); c++ {
		if _, ok := curr.children[word[c]]; !ok {
			curr.children[word[c]] = &TrieNode{children: make(map[byte]*TrieNode)}
		}
		curr = curr.children[word[c]]
	}
	curr.isEnd = true
}

func (t *WordDictionary) Search(word string) bool {
	var dfs func(int, *TrieNode) bool
	dfs = func(j int, root *TrieNode) bool {
		curr := root

		for i := j; i < len(word); i++ {
			c := word[i]
			if c == '.' {
				for _, child := range curr.children {
					if dfs(i+1, child) {
						return true
					}
				}
				return false
			} else {
				if _, ok := curr.children[c]; !ok {
					return false
				}
				curr = curr.children[c]
			}
		}
		return curr.isEnd
	}

	return dfs(0, t.root)
}

// Time complexity: O(M) for the "well-defined" words without dots, where M is the key length, and N is a number of keys, and O(N⋅26^M)
// for the "undefined" words. That corresponds to the worst-case situation of searching an undefined word M times .........
// Space complexity: O(1) for the search of "well-defined" words without dots, and up to O(M) for the "undefined" words, to keep the recursion stack.

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
