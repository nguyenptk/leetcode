// https://leetcode.com/problems/word-search-ii/
package hard

type TrieNodeW struct {
	Children []*TrieNodeW
	Word     string
}

func NewTrieNodeW() *TrieNodeW {
	return &TrieNodeW{
		Children: make([]*TrieNodeW, 26),
		Word:     "",
	}
}

var root *TrieNodeW = NewTrieNodeW()

func FindWords(board [][]byte, words []string) []string {
	for _, w := range words {
		insertTrie(w)
	}
	result := []string{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfsFindWords(board, i, j, root, &result)
		}
	}

	return result
}

func insertTrie(word string) {
	node := root
	for i := 0; i < len(word); i++ {
		w := word[i] - 'a'
		if node.Children[w] == nil {
			node.Children[w] = NewTrieNodeW()
		}
		node = node.Children[w]
	}
	node.Word = word
}

func dfsFindWords(board [][]byte, i, j int, node *TrieNodeW, result *[]string) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}
	if board[i][j] == '*' {
		return
	}

	c := board[i][j]
	child := node.Children[c-'a']
	if child == nil {
		return
	}
	if child.Word != "" {
		*result = append(*result, child.Word)
		child.Word = ""
	}

	board[i][j] = '*'
	dfsFindWords(board, i+1, j, child, result)
	dfsFindWords(board, i-1, j, child, result)
	dfsFindWords(board, i, j+1, child, result)
	dfsFindWords(board, i, j-1, child, result)
	board[i][j] = c
}
