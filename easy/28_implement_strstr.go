// https://leetcode.com/problems/implement-strstr/
package easy

func StrStr(haystack string, needle string) int {
	if haystack == "" && needle == "" {
		return 0
	}
	if haystack == needle {
		return 0
	}
	lengthSub := len(haystack) - len(needle)
	for i := 0; i <= lengthSub; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
