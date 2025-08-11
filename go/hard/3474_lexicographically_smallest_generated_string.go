// https://leetcode.com/problems/lexicographically-smallest-generated-string/
package hard

func GenerateString(s1 string, s2 string) string {
	n, m := len(s1), len(s2)
	result := make([]byte, n+m-1)
	for i := range result {
		result[i] = 'a'
	}
	visited := make([]bool, 10500)
	for i := 0; i < n; i++ {
		if s1[i] == 'T' {
			for j := 0; j < m; j++ {
				result[i+j] = s2[j]
				visited[i+j] = true
			}
		}
	}
	for i := 0; i < n; i++ {
		if s1[i] == 'F' {
			match := true
			for j := 0; j < m; j++ {
				if result[i+j] != s2[j] {
					match = false
					break
				}
			}
			if match {
				for j := m - 1; j >= 0; j-- {
					if !visited[i+j] {
						result[i+j] = 'b'
						break
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		match := true
		for j := 0; j < m; j++ {
			if result[i+j] != s2[j] {
				match = false
				break
			}
		}
		if (s1[i] == 'T' && !match) || (s1[i] == 'F' && match) {
			return ""
		}
	}
	return string(result)
}
