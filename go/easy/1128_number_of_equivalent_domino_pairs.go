// https://leetcode.com/problems/number-of-equivalent-domino-pairs/
package easy

func NumEquivDominoPairs(dominoes [][]int) int {
	result := 0
	freq := map[int]int{}

	for _, d := range dominoes {
		a, b := d[0], d[1]
		if a > b {
			a, b = b, a
		}
		key := a*10 + b
		result += freq[key]
		freq[key]++
	}

	return result
}
