// https://leetcode.com/problems/integer-to-english-words/
package hard

import (
	"strings"
)

func NumberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	// Equivalent to the ordered dict in Python
	numToWords := []struct {
		val  int
		word string
	}{
		{1000000000, "Billion"}, {1000000, "Million"}, {1000, "Thousand"},
		{100, "Hundred"}, {90, "Ninety"}, {80, "Eighty"}, {70, "Seventy"},
		{60, "Sixty"}, {50, "Fifty"}, {40, "Forty"}, {30, "Thirty"},
		{20, "Twenty"}, {19, "Nineteen"}, {18, "Eighteen"}, {17, "Seventeen"},
		{16, "Sixteen"}, {15, "Fifteen"}, {14, "Fourteen"}, {13, "Thirteen"},
		{12, "Twelve"}, {11, "Eleven"}, {10, "Ten"}, {9, "Nine"}, {8, "Eight"},
		{7, "Seven"}, {6, "Six"}, {5, "Five"}, {4, "Four"}, {3, "Three"},
		{2, "Two"}, {1, "One"},
	}

	var rec func(int) string
	rec = func(n int) string {
		for _, entry := range numToWords {
			if n >= entry.val {
				prefix := ""
				if n >= 100 {
					prefix = rec(n/entry.val) + " "
				}
				suffix := ""
				if n%entry.val != 0 {
					suffix = " " + rec(n%entry.val)
				}
				return prefix + entry.word + suffix
			}
		}
		return ""
	}

	return strings.TrimSpace(rec(num))
}
