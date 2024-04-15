// https://leetcode.com/problems/delete-columns-to-make-sorted/
package easy

func MinDeletionSize(strs []string) int {
	result := 0
	n := len(strs[0])
	for i := 0; i < n; i++ {
		for j := 0; j < len(strs)-1; j++ {
			if strs[j][i] > strs[j+1][i] {
				result++
				break
			}
		}
	}
	return result
}
