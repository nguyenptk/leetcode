// https://leetcode.com/problems/unique-number-of-occurrences/
package easy

func UniqueOccurrences(arr []int) bool {
	freq := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		freq[arr[i]]++
	}

	set := make(map[int]int)
	for _, v := range freq {
		set[v]++
		if count, ok := set[v]; ok {
			if count > 1 {
				return false
			}
		}
	}

	return true
}
