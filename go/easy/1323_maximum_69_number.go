// https://leetcode.com/problems/maximum-69-number/
package easy

func Maximum69Number(num int) int {
	sequence := []int{}
	intToSlice(num, &sequence)
	for i := 0; i < len(sequence); i++ {
		if sequence[i] != 9 {
			sequence[i] = 9
			break
		}
	}

	return sliceToInt(sequence)
}

func intToSlice(n int, sequence *[]int) []int {
	if n != 0 {
		i := n % 10
		*sequence = append([]int{i}, *sequence...)
		return intToSlice(n/10, sequence)
	}

	return *sequence
}

func sliceToInt(sequence []int) int {
	result := 0
	n := 1
	for i := len(sequence) - 1; i >= 0; i-- {
		result += sequence[i] * n
		n *= 10
	}

	return result
}
