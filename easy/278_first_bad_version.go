// https://leetcode.com/problems/first-bad-version/
package easy

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func FirstBadVersion(n int) int {
	l := 1
	r := n
	for l < r {
		m := l + (r-l)/2
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

// mock isBadVersion API
func isBadVersion(n int) bool {
	if n == 4 {
		return true
	}
	return false
}
