package hard

func alienOrder(words []string) string {
	m := make(map[byte]map[byte]struct{})

	for i := range words {
		for j := range words[i] {
			if _, ok := m[words[i][j]]; !ok {
				m[words[i][j]] = make(map[byte]struct{})
			}
		}
	}

	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		minLen := min(len(w1), len(w2))
		if len(w1) > len(w2) && w1[:minLen] == w2[:minLen] {
			return ""
		}
		for j := 0; j < minLen; j++ {
			if w1[j] != w2[j] {
				m[w1[j]][w2[j]] = struct{}{}
				break
			}
		}
	}

	visited := make(map[byte]int)
	var data []byte

	var dfs func(c byte) bool
	dfs = func(c byte) bool {
		if visited[c] == 1 {
			return true
		}
		if visited[c] == -1 {
			return false
		}
		visited[c] = 1
		for i := range m[c] {
			if dfs(i) {
				return true
			}
		}
		visited[c] = -1
		data = append(data, c)
		return false
	}

	for i := range m {
		if dfs(i) {
			return ""
		}
	}

	var result []byte
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}

	return string(result)
}
