// https://leetcode.com/problems/koko-eating-bananas/
package medium

import (
	"math"
)

func MinEatingSpeed(piles []int, h int) int {
	l := 1
	r := 1
	for _, v := range piles {
		r = max(r, v)
	}

	result := r

	for l <= r {
		k := (l + r) / 2
		hours := float64(0)
		for _, v := range piles {
			hours += math.Ceil(float64(v) / float64(k))
		}

		if hours <= float64(h) {
			result = min(result, k)
			r = k - 1
		} else {
			l = k + 1
		}
	}

	return result
}
