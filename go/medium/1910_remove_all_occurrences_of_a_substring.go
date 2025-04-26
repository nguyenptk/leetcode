// https://leetcode.com/problems/remove-all-occurrences-of-a-substring/
package medium

func RemoveOccurrences(s string, part string) string {
	stack := []rune{} // Stack to simulate a sliding window
	partLen := len(part)

	for _, char := range s {
		stack = append(stack, char) // Add current char to stack

		// Check if the last characters form `part`
		if len(stack) >= partLen && string(stack[len(stack)-partLen:]) == part {
			stack = stack[:len(stack)-partLen] // Remove the occurrence
		}
	}

	return string(stack) // Convert back to string
}
