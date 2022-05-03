// https://leetcode.com/problems/plus-one/
package easy

func PlusOne(digits []int) []int {
	n := len(digits) - 1
	last := digits[n] + 1

	// if last is not 10
	if last != 10 {
		digits = append(digits[:n])
		digits = append(digits, last)
		return digits

		// if there is 2 elements
	} else if n == 0 && digits[n] == 9 {
		return []int{1, 0}
	}

	// more than 2 elements
	result := []int{}
	for i := n; i >= 0; i-- {
		if i > 0 && digits[i] == 9 {
			if digits[i-1]+1 == 10 {
				result = append([]int{0}, result...)
			} else {
				result = append([]int{0}, result...)
				result = append([]int{digits[i-1] + 1}, result...)
				break
			}
		} else if i == 0 {
			if digits[0] == 9 {
				result = append([]int{0}, result...)
			}
		}
	}

	// remove digits's element by result length
	digits = append(digits[:len(digits)-len(result)])

	// if digits length is 0, all of digits's elements are plussed
	if len(digits) == 0 {
		if result[0] == 0 {
			result = append([]int{1}, result...)
		}
		return result
	}

	// digits still has the remaining elements so we need to append result
	digits = append(digits, result...)
	return digits
}
