package medium

import "fmt"

func GroupAnagrams(strs []string) [][]string {
	// Create a map to store the anagrams
	anagrams := make(map[string][]string)

	// Loop through the strs
	for _, str := range strs {
		// Create a slice to store the anagrams
		count := make([]int, 26)
		for _, c := range str {
			// Increment the count of each character
			count[c-'a']++
		}
		// For example of count: 10001000000000000001000000 -> eat, tea, ate
		key := countToString(count)

		// Append the s to the slice of anagrams
		anagrams[key] = append(anagrams[key], str)
	}

	// Convert the map to a slice of slices
	results := make([][]string, 0, len(anagrams))
	for _, v := range anagrams {
		results = append(results, v)
	}

	return results
}

// Convert count to a string
func countToString(count []int) string {
	var str string
	for _, c := range count {
		str += fmt.Sprintf("%02d", c)
	}
	return str
}
