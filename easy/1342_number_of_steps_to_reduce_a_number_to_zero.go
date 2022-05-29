// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
package easy

var count int

func NumberOfSteps(num int) int {
	count = 0
	return countSteps(num)
}

// Recursive count the steps
func countSteps(num int) int {
	// We are at the final step when num = 0 then return the count
	if num == 0 {
		return count
	}

	// If num is even
	if num%2 == 0 {
		num = num / 2

	} else { // If num is odd
		num = num - 1
	}

	count++

	// Continue the recursion
	return countSteps(num)
}
