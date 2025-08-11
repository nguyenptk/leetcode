// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
package medium

func maxVowels(s string, k int) int {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	result, count := 0, 0

	l := 0
	for r := 0; r < len(s); r++ {
		if vowels[s[r]] {
			count++
		}
		if r-l+1 > k {
			if vowels[s[l]] {
				count--
			}
			l++
		}
		if r-l+1 == k {
			result = max(result, count)
		}
	}

	return result
}

// Time complexity: O(n)
// Space complexity: O(1)
