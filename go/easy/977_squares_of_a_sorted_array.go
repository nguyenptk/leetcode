// https://leetcode.com/problems/squares-of-a-sorted-array/
package easy

import (
	"math"
	"sort"
)

func SortedSquares(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		result = append(result, int(math.Pow(float64(nums[i]), 2)))
	}
	sort.Ints(result) // quick sort
	return result
}
