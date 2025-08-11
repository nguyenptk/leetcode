// https://leetcode.com/problems/design-search-autocomplete-system/
package hard

import (
	"strings"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

type TrieNodeAS struct {
	children  map[byte]*TrieNodeAS
	sentences map[string]int
}

type Pair struct {
	sentence string
	count    int
}

type AutocompleteSystem struct {
	root         *TrieNodeAS
	currNode     *TrieNodeAS
	dead         *TrieNodeAS
	currSentence strings.Builder
}

func ConstructorAutocompleteSystem(sentences []string, times []int) AutocompleteSystem {
	trieNodeAS := &TrieNodeAS{
		children:  make(map[byte]*TrieNodeAS),
		sentences: make(map[string]int),
	}
	sys := AutocompleteSystem{
		root:     trieNodeAS,
		currNode: trieNodeAS,
		dead:     &TrieNodeAS{}, // dead node represents a terminal "not found"
	}
	for i, sentence := range sentences {
		sys.addToTrie(sentence, times[i])
	}
	return sys
}

func (f *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		f.addToTrie(f.currSentence.String(), 1)
		f.currSentence.Reset()
		f.currNode = f.root
		return []string{}
	}

	f.currSentence.WriteByte(c)
	if f.currNode == f.dead || f.currNode.children[c] == nil {
		f.currNode = f.dead
		return []string{}
	}
	f.currNode = f.currNode.children[c]

	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		x := a.(Pair)
		y := b.(Pair)
		if x.count == y.count {
			return strings.Compare(y.sentence, x.sentence)
		}
		return x.count - y.count // min-heap by count
	})

	for s, c := range f.currNode.sentences {
		pq.Enqueue(Pair{sentence: s, count: c})
		if pq.Size() > 3 {
			pq.Dequeue()
		}
	}

	res := make([]string, pq.Size())
	for i := pq.Size() - 1; i >= 0; i-- {
		p, _ := pq.Dequeue()
		res[i] = p.(Pair).sentence
	}

	return res
}

func (f *AutocompleteSystem) addToTrie(sentence string, count int) {
	node := f.root
	for i := 0; i < len(sentence); i++ {
		ch := sentence[i]
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &TrieNodeAS{
				children:  make(map[byte]*TrieNodeAS),
				sentences: make(map[string]int),
			}
		}
		node = node.children[ch]
		node.sentences[sentence] += count
	}
}

// Time complexity: O(n⋅k+m⋅(n+ k/m))
// 1. constructor: We initialize the trie, which costs O(n⋅k) as we iterate over each character in each sentence.
// 2. input: We heapify these sentences and find the best 3 in linear time, which costs (n+ k/m),
// is called m times, which gives us a total of O(m⋅(n+ k/m))

// Space complexity: O(k⋅(n⋅k+m))
