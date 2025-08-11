// https://leetcode.com/problems/delete-and-earn/
package medium

func DeleteAndEarn(nums []int) int {
	maxNum := 0
	points := make(map[int]int)

	for _, num := range nums {
		points[num] += num
		maxNum = max(maxNum, num)
	}

	maxPoints := make([]int, maxNum+1)
	maxPoints[1] = points[1]
	for num := 2; num < len(maxPoints); num++ {
		gain := points[num]
		maxPoints[num] = max(maxPoints[num-1], maxPoints[num-2]+gain)
	}
	return maxPoints[maxNum]
}
