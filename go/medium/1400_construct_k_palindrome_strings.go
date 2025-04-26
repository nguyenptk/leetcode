// https://leetcode.com/problems/construct-k-palindrome-strings/
package medium

func CanConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	count := 0
	for _, amount := range m {
		if amount%2 == 1 {
			count++
		}
	}

	return count <= k
}
