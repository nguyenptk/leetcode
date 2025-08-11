// https://leetcode.com/problems/maximum-swap/
package medium

import (
	"strconv"
)

func MaximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	n := len(s)
	maxDigitIdx := -1
	swapIdx1 := -1
	swapIdx2 := -1

	for i := 0; i < n; i++ {
		if maxDigitIdx == -1 || s[i] > s[maxDigitIdx] {
			maxDigitIdx = i
		} else if s[i] < s[maxDigitIdx] {
			swapIdx1 = i
			swapIdx2 = maxDigitIdx
		}
	}

	if swapIdx1 != -1 && swapIdx2 != -2 {
		s[swapIdx1], s[swapIdx2] = s[swapIdx2], s[swapIdx1]
	}

	result, _ := strconv.Atoi(string(s))
	return result
}
