// https://leetcode.com/problems/maximum-product-of-word-lengths/
package medium

import "math"

func MaxProduct(words []string) int {
	best := 0
	bitsets := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		wlen := len(words[i])
		bitset := 0
		for j := 0; j < wlen; j++ {
			bitset |= 1 << (words[i][j] - 'a')
		}
		for j := 0; j < i; j++ {
			if bitsets[j]&bitset == 0 {
				best = int(math.Max(float64(wlen*len(words[j])), float64(best)))
			}
		}
		bitsets[i] = bitset
	}
	return best
}
