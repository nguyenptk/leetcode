// https://leetcode.com/problems/synonymous-sentences/
package medium

import (
	"sort"
	"strings"
)

func GenerateSentences(synonyms [][]string, text string) []string {
	graph := make(map[string][]string)
	for _, s := range synonyms {
		k, v := s[0], s[1]
		graph[k] = append(graph[k], v)
		graph[v] = append(graph[v], k)
	}

	queue := []string{text}
	visited := map[string]bool{text: true}
	result := make([]string, 0)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		result = append(result, curr)
		words := strings.Split(curr, " ")
		for i, w := range words {
			if neis, ok := graph[w]; ok {
				for _, nei := range neis {
					newW := words
					newW[i] = nei
					newSen := strings.Join(newW, " ")
					if !visited[newSen] {
						visited[newSen] = true
						queue = append(queue, newSen)
					}
				}
			}
		}
	}
	sort.Strings(result)

	return result
}
