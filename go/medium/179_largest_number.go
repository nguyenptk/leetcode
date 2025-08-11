// https://leetcode.com/problems/largest-number/
package medium

import (
	"sort"
	"strconv"
	"strings"
)

func LargestNumber(nums []int) string {
	s := make([]string, len(nums))
	for i, v := range nums {
		s[i] = strconv.Itoa(v)
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i]+s[j] > s[j]+s[i]
	})
	if len(s) >= 0 && s[0] == "0" {
		return "0"
	}
	return strings.Join(s, "")
}
