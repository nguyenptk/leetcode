// https://leetcode.com/problems/finding-3-digit-even-numbers/
package easy

func FindEvenNumbers(digits []int) []int {
	result := make([]int, 0)
	freq := make(map[int]int)
	for _, d := range digits {
		freq[d]++
	}
	for i := 100; i < 1000; i += 2 {
		currFreq := make(map[int]int)
		num := i
		for num > 0 {
			d := num % 10
			currFreq[d]++
			num /= 10
		}
		isValid := true
		for d, count := range currFreq {
			if freq[d] < count {
				isValid = false
				break
			}
		}
		if isValid {
			result = append(result, i)
		}
	}
	return result
}
