// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

package medium

func LetterCombinations(digits string) []string {
	// Convert digits to string array
	numbers := []string{}
	for _, v := range digits {
		numbers = append(numbers, string(v))
	}
	return combine(numbers)
}

// Init map of letters
var mapLetter = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

// Combine
func combine(numbers []string) []string {
	return combineWords(numbers, len(numbers), 0, "")
}

// Recursive combine
func combineWords(numbers []string, size int, index int, s string) []string {
	if size == index {
		if s == "" {
			return []string{}
		}
		return []string{s}
	}
	// Init result array
	result := []string{}
	// Get the value of numbers by index
	digit := numbers[index]
	for _, v := range mapLetter[digit] {
		// Combine the backstring with current mapLetter's value
		sCopy := s + v
		// Recursive to append the result
		result = append(result, combineWords(numbers, size, index+1, sCopy)...)
	}
	return result
}
