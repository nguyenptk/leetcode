// https://leetcode.com/problems/stamping-the-sequence/
package hard

func MovesToStamp(stamp string, target string) []int {
	if stamp == target {
		return []int{0}
	}

	sLen := len(stamp)
	tLen := len(target) - sLen + 1
	result := []int{}
	tDiff := true
	sDiff := true
	var i int
	var j int
	for tDiff {
		tDiff = false
		for i = 0; i < tLen; i++ {
			sDiff = false
			for j = 0; j < sLen; j++ {
				if target[i+j] == '*' {
					continue
				} else if target[i+j] != stamp[j] {
					break
				} else {
					sDiff = true
				}
			}
			if j == sLen && sDiff {
				tDiff = true
				for j = i; j < sLen+i; j++ {
					target = target[:j] + "*" + target[j+1:]
				}
				result = append([]int{i}, result...)
			}
		}
	}
	for i = 0; i < len(target); i++ {
		if target[i] != '*' {
			return []int{}
		}
	}

	return result
}
