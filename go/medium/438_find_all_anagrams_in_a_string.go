package medium

func FindAnagrams(s string, p string) []int {
	if len(p) > len(p) {
		return []int{}
	}

	mapP := map[byte]int{}
	for i := 0; i < len(p); i++ {
		if _, ok := mapP[p[i]]; ok {
			mapP[p[i]] += 1
		} else {
			mapP[p[i]] = 1
		}
	}

	result := []int{}
	start := 0
	end := 0

	for end < len(s) {
		if mapP[s[end]] > 0 {
			mapP[s[end]] -= 1
			end++
			if end-start == len(p) {
				result = append(result, start)
			}
		} else if start == end {
			start++
			end++
		} else {
			mapP[s[start]] += 1
			start++
		}
	}

	return result
}
