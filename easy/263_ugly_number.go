// https://leetcode.com/problems/ugly-number/
package easy

func IsUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for _, v := range []int{2, 3, 5} {
		for n%v == 0 {
			n /= v
		}
	}

	return n == 1
}
