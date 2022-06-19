// https://leetcode.com/problems/prefix-and-suffix-search/
package hard

type TrieNode struct {
	Children []*TrieNode
	Vals     []int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make([]*TrieNode, 27),
		Vals:     []int{},
	}
}

type WordFilter struct {
	pTrie *TrieNode
	sTrie *TrieNode
}

func Constructor(words []string) WordFilter {
	pTrie := NewTrieNode()
	sTrie := NewTrieNode()
	for i := 0; i < len(words); i++ {
		word := words[i]
		wlen := len(word)
		insert(word, i, pTrie, 0, wlen, 1)
		insert(word, i, sTrie, wlen-1, -1, -1)
	}
	return WordFilter{pTrie, sTrie}
}

func insert(word string, index int, trie *TrieNode, start, end, step int) {
	for i := start; i != end; i += step {
		c := word[i] - 'a'
		if trie.Children[c] == nil {
			trie.Children[c] = NewTrieNode()
		}
		trie = trie.Children[c]
		trie.Vals = append(trie.Vals, index)
	}
}

func retrieve(word string, trie *TrieNode, start, end, step int) []int {
	for i := start; i != end; i += step {
		trie = trie.Children[word[i]-'a']
		if trie == nil {
			return []int{}
		}
	}
	return trie.Vals
}

func (this *WordFilter) F(prefix string, suffix string) int {
	pVals := retrieve(prefix, this.pTrie, 0, len(prefix), 1)
	sVals := retrieve(suffix, this.sTrie, len(suffix)-1, -1, -1)
	sVix := len(sVals) - 1
	pVix := len(pVals) - 1
	for sVix >= 0 && pVix >= 0 {
		pVal := pVals[pVix]
		sVal := sVals[sVix]
		if pVal == sVal {
			return sVal
		}
		if sVal > pVal {
			sVix--
		} else {
			pVix--
		}
	}
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
