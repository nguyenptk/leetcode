// https://leetcode.com/problems/word-ladder-ii/
package hard

import "math"

func FindLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordList = append([]string{beginWord}, wordList...)
	paths := make([][]int, len(wordList))
	mapW := make([]int, len(wordList))
	endNode := -1

	for i := range wordList {
		mapW[i] = math.MaxInt32

		if wordList[i] == endWord {
			endNode = i
		}

		for j := i + 1; j < len(wordList); j++ {
			if isLadder(wordList[i], wordList[j]) {
				paths[i] = append(paths[i], j)
				paths[j] = append(paths[j], i)
			}
		}
	}

	if endNode == -1 {
		return [][]string{}
	}

	inBound := make([][]int, len(wordList))
	queue := []int{0}
	mapW[0] = 1

	for len(queue) > 0 && queue[0] != endNode {
		now := queue[0]
		queue = queue[1:]

		for _, next := range paths[now] {
			if mapW[now] >= mapW[next] {
				continue
			}

			if mapW[next] == math.MaxInt32 {
				mapW[next] = mapW[now] + 1
				queue = append(queue, next)
			}

			inBound[next] = append(inBound[next], now)
		}
	}

	if mapW[endNode] == math.MaxInt32 {
		return [][]string{}
	}

	result := [][]string{}
	buffer := make([]string, mapW[endNode])

	backtrack(&result, wordList, inBound, buffer, endNode, mapW[endNode]-1)

	return result
}

func backtrack(result *[][]string, wordList []string, inBound [][]int, buffer []string, node, level int) {
	buffer[level] = wordList[node]

	if node == 0 {
		*result = append(*result, append([]string{}, buffer...))
	}

	for _, v := range inBound[node] {
		backtrack(result, wordList, inBound, buffer, v, level-1)
	}
}

func isLadder(a, b string) bool {
	diff := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++

			if diff > 1 {
				return false
			}
		}
	}

	return diff == 1
}
