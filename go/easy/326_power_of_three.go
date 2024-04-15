// https://leetcode.com/problems/power-of-three/
package easy

func IsPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}

	for n != 1 {
		if n%3 != 0 {
			return false
		}
		n = n / 3
	}
	return true
}
