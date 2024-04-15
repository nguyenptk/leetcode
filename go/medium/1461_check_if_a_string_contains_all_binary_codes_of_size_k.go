// https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/
package medium

func HasAllCodes(s string, k int) bool {
	mapCodes := map[string]string{}
	for i := 0; i+k <= len(s); i++ {
		if _, ok := mapCodes[s[i:i+k]]; !ok {
			mapCodes[s[i:i+k]] = s[i : i+k]
		}
	}
	if len(mapCodes) == 1<<k {
		return true
	}
	return false
}
