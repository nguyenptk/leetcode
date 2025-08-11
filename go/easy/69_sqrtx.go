// https://leetcode.com/problems/sqrtx/
package easy

func MySqrt(x int) int {
	if x < 2 {
		return x
	}
	var num int
	l, r := 2, x/2
	for l <= r {
		m := l + (r-l)/2
		num = m * m
		if num > x {
			r = m - 1
		} else if num < x {
			l = m + 1
		} else {
			return m
		}
	}
	return r
}
