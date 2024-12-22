// https://leetcode.com/problems/number-of-1-bits/
package easy

func HammingWeight(num uint32) int {
	result := 0
	for num != 0 {
		result++
		num = num & (num - 1)
	}

	return result
}
