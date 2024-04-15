// https://leetcode.com/problems/ransom-note/
package easy

func CanConstruct(ransomNote string, magazine string) bool {
	count := make([]int, 128)
	for _, v := range magazine {
		count[v]++
	}
	for _, v := range ransomNote {
		count[v]--
		if count[v] < 0 {
			return false
		}
	}
	return true
}
