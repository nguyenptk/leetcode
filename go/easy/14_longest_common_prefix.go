// https://leetcode.com/problems/longest-common-prefix/
package easy

import (
	"sort"
	"strings"
)

// type TrieNode struct {
//     children map[byte]*TrieNode
//     isEnd bool
// }

// func longestCommonPrefix(strs []string) string {
//     root := &TrieNode{
//         children: make(map[byte]*TrieNode),
//     }
//     for _, str := range strs {
//         insert(root, str)
//     }
//     return findLongest(root)
// }

// func insert(root *TrieNode, str string) {
//     for i := 0; i < len(str); i++ {
//         c := str[i]
//         if root.children[c] == nil {
//             root.children[c] = &TrieNode{children: make(map[byte]*TrieNode)}
//         }
//         root = root.children[c]
//     }
//     root.isEnd = true
// }

// func findLongest(root *TrieNode) string {
//     if root == nil {
//         return ""
//     }
//     var result strings.Builder
//     node := root
//     for len(node.children) == 1 && !node.isEnd {
//         for ch, child := range node.children {
//             result.WriteByte(ch)
//             node = child
//         }
//     }
//     return result.String()
// }

func LongestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	l, r := strs[0], strs[len(strs)-1]
	var result strings.Builder

	for i := 0; i < len(l); i++ {
		if i < len(l) && l[i] == r[i] {
			result.WriteByte(l[i])
		} else {
			break
		}
	}

	return result.String()
}
