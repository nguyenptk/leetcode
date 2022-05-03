// https://leetcode.com/problems/roman-to-integer/
package easy

func RomanToInt(s string) int {
	if s == "" {
		return 0
	}

	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	total := 0
	for i := len(s) - 1; i >= 0; {
		n := string(s[i])
		var m string
		if i != 0 {
			m = string(s[i-1])
		}
		if i >= 1 && romanMap[n] > romanMap[m] {
			total = total + romanMap[n] - romanMap[m]
			i -= 2
		} else {
			total += romanMap[n]
			i -= 1
		}
	}
	return total
}
