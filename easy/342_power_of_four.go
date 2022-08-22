// https://leetcode.com/problems/power-of-four/
package easy

func IsPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}

	for n != 1 {
		if n%4 != 0 {
			return false
		}
		n = n / 4
	}
	return true
}
