// https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/
package medium

func MinOperationsMoveBalls(boxes string) []int {
	n := len(boxes)
	ballLeft := 0
	moveLeft := 0
	ballRight := 0
	moveRight := 0

	result := make([]int, n)
	for i := range boxes {
		result[i] += moveLeft
		ballLeft += int(boxes[i] - '0')
		moveLeft += ballLeft

		j := n - 1 - i
		result[j] += moveRight
		ballRight += int(boxes[j] - '0')
		moveRight += ballRight
	}

	return result
}
