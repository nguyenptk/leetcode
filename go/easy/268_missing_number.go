// https://leetcode.com/problems/missing-number/
package easy

func MissingNumber(nums []int) int {
	// Init sum of all numbers
	sum := 0
	// Init total by the sum of all numbers
	total := len(nums) * (len(nums) + 1) / 2
	// Iterate through the array
	for _, num := range nums {
		// Add the number to the sum
		sum += num
	}
	// Return the difference between the sum and the total that is the missing number
	return total - sum
}
