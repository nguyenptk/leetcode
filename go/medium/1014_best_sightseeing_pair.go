// https://leetcode.com/problems/best-sightseeing-pair
package medium

func MaxScoreSightseeingPair(values []int) int {
	n := len(values)
	maxLeftScore := values[0]
	result := 0

	for i := 1; i < n; i++ {
		currRightScore := values[i] - i
		result = max(result, maxLeftScore+currRightScore)

		currLeftScore := values[i] + i
		maxLeftScore = max(maxLeftScore, currLeftScore)
	}

	return result
}
