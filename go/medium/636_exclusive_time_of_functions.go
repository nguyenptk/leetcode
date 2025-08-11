// https://leetcode.com/problems/exclusive-time-of-functions/
package medium

import (
	"strconv"
	"strings"
)

func ExclusiveTime(n int, logs []string) []int {
	result := make([]int, n)

	stack := make([]int, 0)
	s := strings.Split(logs[0], ":")
	start, _ := strconv.Atoi(s[0])
	stack = append(stack, start)

	i := 1
	time, _ := strconv.Atoi(s[2])
	for i < len(logs) {
		s = strings.Split(logs[i], ":")
		currTime, _ := strconv.Atoi(s[2])
		for time < currTime {
			result[stack[len(stack)-1]]++
			time++
		}
		if s[1] == "start" {
			currStart, _ := strconv.Atoi(s[0])
			stack = append(stack, currStart)
		} else {
			result[stack[len(stack)-1]]++
			time++
			stack = stack[:len(stack)-1]
		}
		i++
	}

	return result
}
