// https://leetcode.com/problems/word-break-ii/
package hard

type TrieNodeA struct {
	children []*TrieNodeA
	isEnd    bool
}

type WordDictionary struct {
	root *TrieNodeA
}

func ConstructorWordBreak() WordDictionary {
	return WordDictionary{
		root: &TrieNodeA{
			children: make([]*TrieNodeA, 26),
			isEnd:    false,
		},
	}
}

func (f *WordDictionary) insert(s string) {
	curr := f.root
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &TrieNodeA{children: make([]*TrieNodeA, 26)}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

func WordBreak(s string, wordDict []string) []string {
	trie := ConstructorWordBreak()
	for _, w := range wordDict {
		trie.insert(w)
	}
	subM := make(map[int][]string)
	for i := len(s); i >= 0; i-- {
		validSentences := make([]string, 0)
		curr := trie.root
		for j := i; j < len(s); j++ {
			c := s[j]
			idx := c - 'a'
			if curr.children[idx] == nil {
				break
			}
			curr = curr.children[idx]
			if curr.isEnd {
				currW := s[i : j+1]
				if j == len(s)-1 {
					validSentences = append(validSentences, currW)
				} else {
					nextvalidSentences := subM[j+1]
					for _, sen := range nextvalidSentences {
						validSentences = append(validSentences, currW+" "+sen)
					}
				}
			}
		}
		subM[i] = validSentences
	}
	return subM[0]
}

// Time complexity: O(n⋅2^n)
// Space complexity: O(n⋅2^n)
