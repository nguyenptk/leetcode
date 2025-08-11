// https://leetcode.com/problems/minimum-swaps-to-group-all-1s-together/
package medium

func MinSwapsToGroupAll1sTogether(data []int) int {
	var ones int
	for _, d := range data {
		ones += d
	}
	cntOne, maxOne := 0, 0
	l, r := 0, 0
	for r < len(data) {
		cntOne += data[r]
		r++
		if r-l > ones {
			cntOne -= data[l]
			l++
		}
		maxOne = max(maxOne, cntOne)
	}

	return ones - maxOne
}
