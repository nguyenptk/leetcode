// https://leetcode.com/problems/reverse-vowels-of-a-string/
package easy

func ReverseVowels(s string) string {
	m := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	l := 0
	r := len(s) - 1

	arr := []byte(s)
	for l <= r {
		if m[arr[l]] && m[arr[r]] {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		} else if !m[arr[l]] {
			l++
		} else if !m[arr[r]] {
			r--
		}
	}

	return string(arr)
}
