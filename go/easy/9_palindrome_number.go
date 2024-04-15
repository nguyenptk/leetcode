// https://leetcode.com/problems/palindrome-number/
package easy

import "strconv"

func IsPalindrome(x int) bool {
	str := strconv.Itoa(x)
	r := []rune(str)
	n := len(r)

	if n == 2 {
		return r[0] == r[1]
	}

	for i := 0; i < n/2; i++ {
		if r[i] != r[n-1-i] {
			return false
		}
	}
	return true
}
