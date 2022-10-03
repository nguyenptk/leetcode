// https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
package medium

func MinCost(colors string, neededTime []int) int {
	sum := 0
	n := len(neededTime)
	for _, v := range neededTime {
		sum += v
	}

	if n == 1 {
		return neededTime[0]
	}

	curr := colors[0]
	currMax := neededTime[0]
	sumMax := 0

	for i := 1; i < n; i++ {
		c := colors[i]
		if c == curr {
			if currMax < neededTime[i] {
				currMax = neededTime[i]
			}
		} else {
			sumMax += currMax
			curr = c
			currMax = neededTime[i]
		}
	}

	return sum - sumMax - currMax
}
