// https://leetcode.com/problems/text-justification
package medium

import "strings"

func FullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0)
	i := 0
	for i < len(words) {
		currLine := getWords(i, words, maxWidth)
		i += len(currLine)
		result = append(result, createLine(currLine, i, words, maxWidth))
	}

	return result
}

func getWords(i int, words []string, maxWidth int) []string {
	currLine := make([]string, 0)
	currLength := 0
	for i < len(words) && currLength+len(words[i]) <= maxWidth {
		currLine = append(currLine, words[i])
		currLength += len(words[i]) + 1
		i++
	}
	return currLine
}

func createLine(line []string, i int, words []string, maxWidth int) string {
	baseLength := -1
	for _, word := range line {
		baseLength += len(word) + 1
	}
	extraSpaces := maxWidth - baseLength
	if len(line) == 1 || i == len(words) {
		return strings.Join(line, " ") + strings.Repeat(" ", extraSpaces)
	}
	wordCount := len(line) - 1
	needsExtraSpace := extraSpaces % wordCount
	spacesPerWord := extraSpaces / wordCount
	for j := 0; j < needsExtraSpace; j++ {
		line[j] += " "
	}
	for j := 0; j < wordCount; j++ {
		line[j] += strings.Repeat(" ", spacesPerWord)
	}
	return strings.Join(line, " ")
}
