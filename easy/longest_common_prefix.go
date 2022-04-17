// https://leetcode.com/problems/longest-common-prefix/
package easy

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	s0 := strs[0]
	for _, s := range strs {
		s0 = compareAndPick(s0, s)
	}
	return s0
}

func compareAndPick(first, second string) string {
	var result []byte
	for i := 0; i < len(first) && i < len(second); i++ {
		if first[i] != second[i] {
			break
		}
		result = append(result, first[i])
	}
	return string(result)
}
