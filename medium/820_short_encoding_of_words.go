// https://leetcode.com/problems/short-encoding-of-words/
package medium

func MinimumLengthEncoding(words []string) int {
	// Initialize a map to store the words array by key
	mapW := make(map[string]bool)
	for _, w := range words {
		mapW[w] = true
	}

	// Iterate the mapW to get the key
	for k := range mapW {
		// Iterate the key to get character
		for i := 1; i < len(k); i++ {
			// Delete the key of mapW by character
			s := k[i:]
			delete(mapW, s)
		}
	}

	// Initialize the result with length of mapW because of #
	result := len(mapW)
	// Iterate the mapW to count the length of words
	for k := range mapW {
		result += len(k)
	}
	return result
}
