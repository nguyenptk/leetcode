// https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks/
package easy

import "math"

func MinimumRecolors(blocks string, k int) int {
	result := math.MaxInt
	l := 0
	whites := 0
	for r := 0; r < len(blocks); r++ {
		if blocks[r] == 'W' {
			whites++
		}
		if r-l+1 == k {
			result = min(result, whites)
			if blocks[l] == 'W' {
				whites--
			}
			l++
		}
	}

	return result
}
