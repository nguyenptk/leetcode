// https://leetcode.com/problems/russian-doll-envelopes/
package hard

import "sort"

type sortEnvelopes [][]int

func (s sortEnvelopes) Len() int {
	return len(s)
}
func (s sortEnvelopes) Less(i, j int) bool {
	if s[i][0] == s[j][0] {
		return s[i][1] > s[j][1]
	}
	return s[i][0] < s[j][0]
}
func (s sortEnvelopes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func MaxEnvelopes(envelopes [][]int) int {
	sort.Sort(sortEnvelopes(envelopes))
	dp := []int{}
	for _, env := range envelopes {
		low, high := 0, len(dp)
		for low < high {
			mid := low + (high-low)/2
			if dp[mid] >= env[1] {
				high = mid
			} else {
				low = mid + 1
			}
		}
		if low == len(dp) {
			dp = append(dp, env[1])
		} else {
			dp[low] = env[1]
		}
	}
	return len(dp)
}
