// https://leetcode.com/problems/implement-trie-prefix-tree/

package medium

type Trie struct {
	root     *Trie
	children [26]*Trie
	isEnd    bool
}

func ConstructorTrie() Trie {
	return Trie{
		root:     &Trie{},
		children: [26]*Trie{},
		isEnd:    false,
	}
}

func (t *Trie) Insert(word string) {
	curr := t.root
	for _, w := range word {
		idx := w - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &Trie{}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root
	for _, w := range word {
		idx := w - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return curr.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root
	for _, p := range prefix {
		idx := p - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
