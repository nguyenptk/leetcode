// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/
package medium

import (
	"strings"
)

func RemoveDuplicates(s string, k int) string {
	// Init array from s string
	array := strings.Split(s, "")

	// Init stack and push 0
	stack := []int{}
	stack = append(stack, 0)

	// Init counter i, j
	i := 1 // Keeps track of the end of the in-place "stack"
	j := 1 // Iterates through SC normally

	// Iterate the s string by j
	for j < len(s) {
		// Override array[i] by array[j] to write i to stack
		array[i] = array[j]
		if i == 0 || array[i] != array[i-1] {
			stack = append(stack, i)

			// Remove k from stack
		} else if i-stack[len(stack)-1]+1 == k {
			val := stack[len(stack)-1]   // Top the stack
			stack = stack[:len(stack)-1] // Pop the stack
			i = val - 1                  // Override i by stack top - 1
		}
		i++
		j++
	}
	// Get the first parts of aray by i
	array = array[0:i]

	return strings.Join(array, "")
}
