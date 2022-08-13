// https://leetcode.com/problems/substring-with-concatenation-of-all-words/
package hard

func FindSubstring(s string, words []string) []int {
	result := []int{}

	// Init a map to store frequency of words
	mapW := map[string]int{}
	for _, w := range words {
		if _, ok := mapW[w]; ok {
			mapW[w] = mapW[w] + 1
		} else {
			mapW[w] = 1
		}
	}

	n := len(words[0])

	for i := 0; i < n; i++ {
		currMap := map[string]int{}
		start := i
		count := 0

		for j := i; j <= len(s)-n; j += n {
			sub := s[j : j+n]
			if _, ok := mapW[sub]; ok {
				// Set frequency in current map
				if _, ok := currMap[sub]; ok {
					currMap[sub] = currMap[sub] + 1
				} else {
					currMap[sub] = 1
				}

				count++

				for currMap[sub] > mapW[sub] {
					// Shift left
					left := s[start : start+n]
					currMap[left] = currMap[left] - 1
					count--
					start += n
				}

				if count == len(words) {
					// Add to result
					result = append(result, start)

					// Shift left
					left := s[start : start+n]
					currMap[left] = currMap[left] - 1
					count--
					start = start + n
				}
			} else {
				// Clear the current map
				currMap = map[string]int{}
				start = j + n
				count = 0
			}

		}
	}

	return result
}
