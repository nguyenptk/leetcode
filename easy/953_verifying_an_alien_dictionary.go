// https://leetcode.com/problems/verifying-an-alien-dictionary
package easy

func IsAlienSorted(words []string, order string) bool {
	orderMap := make([]int, 26)
	for i := 0; i < len(order); i++ {
		orderMap[order[i]-'a'] = i
	}

	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words[i+1]) {
				return false
			}
			if words[i][j] != words[i+1][j] {
				curr := words[i][j] - 'a'
				next := words[i+1][j] - 'a'
				if orderMap[curr] > orderMap[next] {
					return false
				}
				break
			}
		}
	}

	return true
}
