// https://leetcode.com/problems/greatest-common-divisor-of-strings/
package easy

func GcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		return GcdOfStrings(str2, str1)
	}

	if str1[:len(str2)] != str2 {
		return ""
	}

	str1 = str1[len(str2):]

	if len(str1) == 0 {
		return str2
	}

	return GcdOfStrings(str1, str2)
}
