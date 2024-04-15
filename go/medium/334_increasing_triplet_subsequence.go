// https://leetcode.com/problems/increasing-triplet-subsequence/
package medium

import "math"

func IncreasingTriplet(nums []int) bool {
	first := math.MaxInt
	second := math.MaxInt

	for _, v := range nums {
		if v <= first {
			first = v // first < num <= second
		} else if v <= second {
			second = v // first < second < num (third)
		} else {
			return true
		}
	}

	return false
}
