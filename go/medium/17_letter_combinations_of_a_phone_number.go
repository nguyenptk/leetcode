// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

package medium

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	mapData := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	result := make([]string, 0)
	curr := ""

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if len(curr) == len(digits) {
			result = append(result, curr)
			return
		}
		digit := digits[idx]
		strs := mapData[digit]
		for i := 0; i < len(strs); i++ {
			curr += strs[i]
			backtrack(idx + 1)
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(0)

	return result
}
