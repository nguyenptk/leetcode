// https://leetcode.com/problems/stickers-to-spell-word/
package hard

import "math"

func MinStickers(stickers []string, target string) int {
	m := len(stickers)
	freq := make([][26]int, m)
	for i, s := range stickers {
		for _, c := range s {
			freq[i][c-'a']++
		}
	}

	memo := make(map[string]int)
	memo[""] = 0

	var dfs func(string) int
	dfs = func(remain string) int {
		if val, ok := memo[remain]; ok {
			return val
		}

		target := [26]int{}
		for _, c := range remain {
			target[c-'a']++
		}

		minSticker := math.MaxInt32
		for _, c := range freq {
			// Skip stickers that don't contain the first char of target
			if c[remain[0]-'a'] == 0 {
				continue
			}

			// Apply sticker
			newTarget := ""
			for i := 0; i < 26; i++ {
				if target[i] > 0 {
					remainCount := target[i] - c[i]
					for j := 0; j < max(0, remainCount); j++ {
						newTarget += string('a' + i)
					}
				}
			}

			count := dfs(newTarget)
			if count != -1 {
				minSticker = min(minSticker, count+1)
			}
		}

		if minSticker == math.MaxInt32 {
			memo[remain] = -1
		} else {
			memo[remain] = minSticker
		}

		return memo[remain]
	}

	return dfs(target)
}
