// https://leetcode.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/
package medium

func GetHappyString(n int, k int) string {
	stack := []string{""}
	idx := 0

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(curr) == n {
			idx++
			if idx == k {
				return curr
			}
			continue
		}

		for _, c := range []string{"c", "b", "a"} {
			if len(curr) == 0 || curr[len(curr)-1:] != c {
				stack = append(stack, curr+c)
			}
		}
	}

	return ""
}
