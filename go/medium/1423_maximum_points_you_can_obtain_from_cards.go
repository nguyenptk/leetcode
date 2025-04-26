// https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/
package medium

func MaxScore(cardPoints []int, k int) int {
	n := len(cardPoints)

	sum := 0
	for i := 0; i < n; i++ {
		sum += cardPoints[i]
	}
	if k >= n {
		return sum
	}

	result := 0
	l := 0
	subSum := 0

	for r := 0; r < n; r++ {
		subSum += cardPoints[r]

		if r-l+1 == n-k {
			result = max(result, sum-subSum)
			subSum -= cardPoints[l]
			l++
		}
	}

	return result
}
