package easy

func ValidWordAbbreviation(word string, abbr string) bool {
	m := len(word)
	n := len(abbr)
	i := 0
	j := 0

	for i < m && j < n {
		if word[i] == abbr[j] {
			i++
			j++
		} else if '0' <= abbr[j] && abbr[j] <= '9' {
			if abbr[j] == '0' {
				return false
			}
			number := 0
			for j < n && '0' <= abbr[j] && abbr[j] <= '9' {
				digit := int(abbr[j] - '0')
				number = number*10 + digit
				j++
			}
			i += number
		} else {
			return false
		}
	}

	return i == m && j == n
}
